package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	// serve static files
	// see https://www.alexedwards.net/blog/disable-http-fileserver-directory-listings for options to disable directory listings
	file_server := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", file_server))

	// homepage and CRUD routes
	mux.HandleFunc("GET /{$}", app.snippetsIndex)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetsShow)
	mux.HandleFunc("GET /snippet/create", app.snippetsCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetsStore)

	// setup middleware chain, using alice package instead of nesting funcs eg: app.recoverPanic(app.logRequest(commonHeaders(mux)))
	middleware := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return middleware.Then(mux)
}
