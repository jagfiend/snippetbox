package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// serve static files
	file_server := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", file_server))

	// handle routes
	mux.HandleFunc("GET /{$}", home)

	mux.HandleFunc("GET /snippets/{$}", snippetsIndex)
	mux.HandleFunc("GET /snippets/{id}", snippetsShow)
	mux.HandleFunc("GET /snippets/create", snippetsCreate)
	mux.HandleFunc("POST /snippets", snippetsStore)

	log.Print("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)

	log.Fatal(err)
}
