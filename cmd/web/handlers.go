package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jagfiend/snippetbox/internal/models"
	"github.com/jagfiend/snippetbox/internal/validator"
)

type snippetCreateForm struct {
	Title               string `form:"title`
	Content             string `form:"content"`
	Expires             int    `form:"expires"`
	validator.Validator `form:"-"`
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
	data := app.newTemplateData(r)
	data.Form = snippetCreateForm{Expires: 365}

	app.render(w, r, http.StatusOK, "create.tmpl.html", data)
}

func (app *application) snippetsStore(w http.ResponseWriter, r *http.Request) {
	var form snippetCreateForm

	err := app.decodePostForm(r, &form)

	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// error handling
	form.CheckField(validator.NotBlank(form.Title), "title", "The title field cannot be blank")
	form.CheckField(validator.MaxChars(form.Title, 100), "title", "The title field cannot be more than 100 characters")
	form.CheckField(validator.NotBlank(form.Content), "content", "The content field cannot be blank")
	form.CheckField(validator.PermittedValue(form.Expires, 1, 7, 365), "expires", "Invalid value give for field expires")

	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, r, http.StatusUnprocessableEntity, "create.tmpl.html", data)
		return
	}

	id, err := app.snippets.Create(form.Title, form.Content, form.Expires)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippets/%d", id), http.StatusSeeOther)
}
