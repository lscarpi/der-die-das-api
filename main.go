package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func BootServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	router := mux.NewRouter()
	router.Use(commonMiddleware)

	router.HandleFunc("/{word}", GetDictWord)

	fmt.Println(fmt.Sprintf("Server ready and listening at port %s", port))
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(fmt.Sprintf("Incoming request: %s %s from %s", r.Method, r.RequestURI, getRequestIP(r)))
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func getRequestIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

func main() {
	BootServer()
}
