package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/jagfiend/snippetbox/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	logger   *slog.Logger
	snippets *models.SnippetModel
}

func main() {
	// read initialisation flags
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))

	db, err := openDb(*dsn)

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// this wouldnt run as the os.Exit() below will shut down the application without running
	// the deferred functions, leaving here in case we add graceful shutdown later
	defer db.Close()

	// spin up application
	app := &application{
		logger:   logger,
		snippets: &models.SnippetModel{DB: db},
	}

	logger.Info("attempting to start server", "addr", *addr)

	err = http.ListenAndServe(*addr, app.routes())

	logger.Error(err.Error())
	// explicit call to os.Exit() as slog does not have a fatal method like log
	os.Exit(1)
}

func openDb(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		db.Close()

		return nil, err
	}

	return db, nil
}
