package handlers

import (
	"DerDieDasApi/fetchers"
	"DerDieDasApi/repositories"
	"DerDieDasApi/types"
	"errors"
	"fmt"
)

type DictHandler struct {
}

func (r DictHandler) GetDictWord(word string) (types.DictWord, error) {

	// Try locally
	dictWord, found := getDictWordLocally(word)
	if found {
		return dictWord, nil
	}

	// Try remotely
	dictWord, found = getDictWordRemotely(word)

	if found {
		dictRepo := repositories.DictRepository{}
		dictRepo.Store(dictWord)
		return dictWord, nil
	}

	// Not found anywhere
	return types.DictWord{}, errors.New(fmt.Sprintf("Could not find word %s. This could also be a non-valid word for this application.", word))
}

func getDictWordLocally(word string) (types.DictWord, bool) {
	dictRepo := repositories.DictRepository{}
	return dictRepo.Find(word)
}

func getDictWordRemotely(word string) (types.DictWord, bool) {
	fetcher := fetchers.DictCC{}
	return fetcher.Fetch(word)
}
