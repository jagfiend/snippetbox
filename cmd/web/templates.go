package main

import (
	"path/filepath"
	"text/template"
	"time"

	"github.com/jagfiend/snippetbox/internal/models"
)

type templateData struct {
	CurrentYear int
	Snippet     models.Snippet
	Snippets    []models.Snippet
}

func humanReadableDate(t time.Time) string {
	// this format is oddly specific, need to investigate
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanReadableDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
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
