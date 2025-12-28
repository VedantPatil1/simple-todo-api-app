package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/VedantPatil1/simple-todo-api-app.git/internal"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	slog.Info("Starting App")

	server := internal.NewServer(logger)

	http.ListenAndServe(":8080", server)
}
