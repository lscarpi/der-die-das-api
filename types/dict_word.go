package types

type DictWord struct {
	Word     string `json:"word"`
	Article  string `json:"article"`
	Type     string `json:"type"`
	IsPlural bool   `json:"is_plural"`
}
