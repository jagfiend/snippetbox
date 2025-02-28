package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jagfiend/snippetbox/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// playing with redirects
	http.Redirect(w, r, "/snippets", http.StatusPermanentRedirect)
}

func (app *application) snippetsIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Wondini")

	snippets, err := app.snippets.Latest()

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := app.newTemplateData(r)
	data.Snippets = snippets

	app.render(w, r, http.StatusOK, "home.tmpl.html", data)
}

func (app *application) snippetsShow(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.snippets.Get(id)

	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}

		return
	}

	data := app.newTemplateData(r)
	data.Snippet = snippet

	app.render(w, r, http.StatusOK, "view.tmpl.html", data)
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
