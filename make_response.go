package main

import (
	"DerDieDasApi/types"
	"encoding/json"
	"log"
	"strings"
)

func MakeSuccessResponse(dictWord types.DictWord) []byte {

	// Uppercase first letter just in case
	dictWord.Word = strings.Title(dictWord.Word)

	bytes, err := json.Marshal(dictWord)

	if err != nil {
		log.Fatalln("Error while making Success response")
	}

	return bytes
}

func MakeErrorResponse(message string) []byte {
	res := ErrorResponse{
		Message: message,
	}

	bytes, err := json.Marshal(res)

	if err != nil {
		log.Fatalln("Error while making Error response")
	}

	return bytes
}
