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
	logger.Info("Initialized in-memory todo data store")

	// Adding initial data for testing.
	todo := todoStore.AddTodo("Add test todo for ui testing")
	todoStore.AddTodo("remove the previous todo")
	todoStore.ToggleCompletion(todo.Id)
	logger.Info("Added initial test todos")
	logger.Info("Initial Todos", "todos", todoStore.GetTodos())

	deps := internal.NewContainer(logger, todoStore)
	server := internal.NewServer(deps)

	http.ListenAndServe(":8080", server)
}
