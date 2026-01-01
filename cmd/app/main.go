package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/VedantPatil1/simple-todo-api-app.git/internal"
	"github.com/VedantPatil1/simple-todo-api-app.git/internal/todos"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	slog.Info("Starting App")

	todoStore := todos.NewInMemoryDataStore()
	deps := internal.NewContainer(logger, todoStore)
	server := internal.NewServer(deps)

	http.ListenAndServe(":8080", server)
}
