package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	// serve static files
	// see https://www.alexedwards.net/blog/disable-http-fileserver-directory-listings for options to disable directory listings
	file_server := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", file_server))

	// handle routes and handlers
	mux.HandleFunc("GET /{$}", app.home)

	mux.HandleFunc("GET /snippets/{$}", app.snippetsIndex)
	mux.HandleFunc("GET /snippets/{id}", app.snippetsShow)
	mux.HandleFunc("GET /snippets/create", app.snippetsCreate)
	mux.HandleFunc("POST /snippets", app.snippetsStore)

	return mux
}
