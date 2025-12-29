package internal

import (
	"io/fs"
	"log/slog"
	"net/http"

	"github.com/VedantPatil1/simple-todo-api-app.git/internal/assets"
	"github.com/VedantPatil1/simple-todo-api-app.git/internal/middleware"
	"github.com/VedantPatil1/simple-todo-api-app.git/internal/templates/pages"
	"github.com/a-h/templ"
)

func NewServer(logger *slog.Logger) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/", templ.Handler(pages.HomePage()))

	StaticFS, err := fs.Sub(assets.EmbeddedFiles, "static")
	if err != nil {
		panic(err)
	}

	fs := http.FileServer(http.FS(StaticFS))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	loggingMiddleware := middleware.LogRequest(logger)

	server := loggingMiddleware(mux)
	return server
}
