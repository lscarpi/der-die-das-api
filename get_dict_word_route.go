package main

import (
	"DerDieDasApi/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func GetDictWord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word, ok := vars["word"]
	if !ok {
		w.Write(MakeErrorResponse("Please inform a word to be evaluated."))
		return
	}

	dictHandler := handlers.DictHandler{}
	dictWord, err := dictHandler.GetDictWord(word)
	if err != nil {
		w.Write(MakeErrorResponse(err.Error()))
		return
	}

	w.Write(MakeSuccessResponse(dictWord))
}
