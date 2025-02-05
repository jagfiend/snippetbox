package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	// read initialisation flags
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// spin up application
	app := &application{
		logger: slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		})),
	}

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

	app.logger.Info("attempting to start server", "addr", *addr)

	err := http.ListenAndServe(*addr, mux)

	app.logger.Error(err.Error())
	// explicit call to os.Exit() as slog does not have a fatal method like log
	os.Exit(1)
}
