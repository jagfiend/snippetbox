package main

import (
	"net/http"

	"github.com/jagfiend/snippetbox/ui"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	// serve static files from disk
	// see https://www.alexedwards.net/blog/disable-http-fileserver-directory-listings for options to disable directory listings
	// file_server := http.FileServer(http.Dir("./ui/static/"))
	// mux.Handle("GET /static/", http.StripPrefix("/static", file_server))

	// serve static files from embedded file system
	mux.Handle("GET /static/", http.FileServerFS(ui.Files))

	dynamic := alice.New(app.sessionManager.LoadAndSave, app.noSurf, app.authenticate)

	mux.Handle("GET /{$}", dynamic.ThenFunc(app.snippetsIndex))
	mux.Handle("GET /snippet/view/{id}", dynamic.ThenFunc(app.snippetsShow))
	mux.Handle("GET /user/signup", dynamic.ThenFunc(app.userSignup))
	mux.Handle("POST /user/signup", dynamic.ThenFunc(app.userSignupPost))
	mux.Handle("GET /user/login", dynamic.ThenFunc(app.userLogin))
	mux.Handle("POST /user/login", dynamic.ThenFunc(app.userLoginPost))

	protected := dynamic.Append(app.requireAuthentication)

	mux.Handle("GET /snippet/create", protected.ThenFunc(app.snippetsCreate))
	mux.Handle("POST /snippet/create", protected.ThenFunc(app.snippetsStore))
	mux.Handle("POST /user/logout", protected.ThenFunc(app.userLogoutPost))

	// setup middleware chain, using alice package instead of nesting funcs
	// eg: app.recoverPanic(app.logRequest(commonHeaders(mux)))
	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}
