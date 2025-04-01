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

	dynamic := alice.New(app.sessionManager.LoadAndSave)

	// homepage and CRUD routes
	mux.Handle("GET /{$}", dynamic.ThenFunc(app.snippetsIndex))
	mux.Handle("GET /snippet/view/{id}", dynamic.ThenFunc(app.snippetsShow))
	mux.Handle("GET /snippet/create", dynamic.ThenFunc(app.snippetsCreate))
	mux.Handle("POST /snippet/create", dynamic.ThenFunc(app.snippetsStore))

	// setup middleware chain, using alice package instead of nesting funcs eg: app.recoverPanic(app.logRequest(commonHeaders(mux)))
	middleware := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return middleware.Then(mux)
}
