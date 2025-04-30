package main

import (
	"io/fs"
	"path/filepath"
	"text/template"
	"time"

	"github.com/jagfiend/snippetbox/internal/models"
	"github.com/jagfiend/snippetbox/ui"
)

type templateData struct {
	CurrentYear     int
	Snippet         models.Snippet
	Snippets        []models.Snippet
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
}

func humanReadableDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.UTC().Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanReadableDate,
}

// use embedded fs
func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl.html")

	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		patterns := []string{
			"html/base.tmpl.html",
			"html/partials/*.tmpl.html",
			page,
		}

		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)

		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}

// without embedded fs
func oldTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html")

	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		// parse base template
		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.tmpl.html")

		if err != nil {
			return nil, err
		}

		// add partials
		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl.html")

		if err != nil {
			return nil, err
		}

		// add page
		ts, err = ts.ParseFiles(page)

		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
