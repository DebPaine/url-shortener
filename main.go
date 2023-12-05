package main

import (
	"fmt"
	"net/http"

	"github.com/DebPaine/url-shortener/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/shorten", handlers.ShortenHandler)
	port := 8000

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
	if err != nil {
		fmt.Println(err)
	}
}
