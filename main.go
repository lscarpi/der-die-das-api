package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func BootServer() {
	router := mux.NewRouter()
	router.Use(commonMiddleware)
	router.HandleFunc("/{word}", GetDictWord)
	log.Fatal(http.ListenAndServe(":80", router))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {
	BootServer()
}
