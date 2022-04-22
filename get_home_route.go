package main

import (
	"net/http"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	w.Write(MakeErrorResponse("Please use the correct url structure. E.g. /bier"))
	return
}
