package repositories

import (
	"DerDieDasApi/types"
	"strings"
)

type DictRepository struct {
}

func (r DictRepository) Find(word string) (types.DictWord, bool) {
	ret := types.DictWord{}
	fb := Firebase{}

	result := fb.FindByKey("words", word)

	if result == nil {
		return ret, false
	}

	ret.Word = word
	ret.Article = result["article"].(string)
	ret.IsPlural = result["is_plural"].(bool)
	ret.Type = result["type"].(string)

	return ret, true
}

func (r DictRepository) Store(dictWord types.DictWord) {
	fb := Firebase{}

	fb.Store("words", strings.Title(dictWord.Word), map[string]interface{}{
		"article":   dictWord.Article,
		"is_plural": dictWord.IsPlural,
		"type":      dictWord.Type,
	})
}
