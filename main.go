package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hellope world"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte(fmt.Sprintf("Display snippet id: %d", id)))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a snippet"))
}

func snippetStore(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Store a snippet"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippets/{id}", snippetView)
	mux.HandleFunc("GET /snippets/create", snippetCreate)
	mux.HandleFunc("POST /snippets", snippetStore)

	log.Print("starting server on port :4000")

	err := http.ListenAndServe(":4000", mux)

	log.Fatal(err)
}
