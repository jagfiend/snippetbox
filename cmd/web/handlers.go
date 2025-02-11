package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// playing with redirects
	http.Redirect(w, r, "/snippets", http.StatusPermanentRedirect)
}

func (app *application) snippetsIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Wondini")

	templates := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	ts, err := template.ParseFiles(templates...)

	if err != nil {
		app.serverError(w, r, err)

		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)

	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) snippetsShow(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Display the snippet with ID %d...", id)
}

func (app *application) snippetsCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display the create snippet form..."))
}

func (app *application) snippetsStore(w http.ResponseWriter, r *http.Request) {
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Create(title, content, expires)

	if err != nil {
		app.serverError(w, r, err)

		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippets/%d", id), http.StatusSeeOther)
}
