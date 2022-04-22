package fetchers

import (
	"DerDieDasApi/types"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	url2 "net/url"
	"regexp"
	"strings"
)

type DictCC struct {
}

// Fetch returns a string for the article found, and a bool if it's plural or not.
func (f DictCC) Fetch(word string) (types.DictWord, bool) {

	url := fmt.Sprintf("https://dict.cc/?s=%s", url2.QueryEscape(strings.ToLower(word)))
	// Get the html page
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln(fmt.Sprintf("Error while fetching word infromation from Dict.cc, word: %s", word))
	}

	// Parse all response to byte array
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	// Fetch only the main table row needed
	tableRow := findTableRow(&body)

	// If nothing
	if tableRow == nil {
		// Ouch
		return types.DictWord{}, false
	}
	// Then find the article and plural
	article, plural := findArticle(&tableRow)

	return types.DictWord{
		Word:     word,
		Article:  string(article),
		Type:     "noun",
		IsPlural: plural,
	}, true
}

func findTableRow(body *[]byte) []byte {
	// <td .*?>(.*?)<\/td>
	r, err := regexp.Compile("<tr title=\"article sg \\| article pl\">(.*?)<\\/tr>")
	if err != nil {
		log.Fatalln(err)
	}

	// Find all rows matching
	result := r.FindAll(*body, -1)

	// Loop through each result
	for _, row := range result {
		// If it contains <u> means, it's underlined and selected
		if bytes.Contains(row, []byte("<u>")) {
			// That's the one we want
			return row
		}
	}

	// Reaching here, nothing found
	return nil
}

func findArticle(tableRow *[]byte) ([]byte, bool) {

	singularRegex, err := regexp.Compile("<a.*?>(der|die|das).*?<u>.*?<\\/a>\\s?<span.*?>\\|")
	pluralRegex, err := regexp.Compile("<span.*?>\\|<\\/span>\\s?<a.*?>(der|die|das).*?.*?<\\/a>")

	if err != nil {
		log.Fatalln(err)
	}

	singularMatches := singularRegex.FindSubmatch(*tableRow)

	if singularMatches != nil {
		return singularMatches[1], false
	}

	pluralMatches := pluralRegex.FindSubmatch(*tableRow)

	if pluralMatches != nil {
		return pluralMatches[1], true
	}

	log.Fatalln(fmt.Sprintf("Could not define the article of the word in the string below: \n %b", tableRow))
	return []byte(""), false
}

//https://dept.dict.cc/?s=Arbeit