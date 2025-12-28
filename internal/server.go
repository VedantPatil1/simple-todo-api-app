package internal

import (
	"log/slog"
	"net/http"

	"github.com/VedantPatil1/simple-todo-api-app.git/internal/middleware"
)

func NewServer(logger *slog.Logger) http.Handler {
	mux := http.NewServeMux()
	loggingMiddleware := middleware.LogRequest(logger)

	server := loggingMiddleware(mux)
	return server
}
