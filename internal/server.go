package internal

import (
	"log/slog"
	"net/http"

	"github.com/VedantPatil1/simple-todo-api-app.git/internal/middleware"
	"github.com/VedantPatil1/simple-todo-api-app.git/internal/todos"
	"github.com/VedantPatil1/simple-todo-api-app.git/internal/ui/pages"
	"github.com/a-h/templ"
)

type Container struct {
	TodoStore todos.TodoStore
	Logger    *slog.Logger
}

func NewContainer(logger *slog.Logger, todoStore todos.TodoStore) *Container {
	return &Container{
		TodoStore: todoStore,
		Logger:    logger,
	}
}

func NewServer(deps *Container) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/", templ.Handler(pages.HomePage()))

	staticDir := "./internal/assets/static"
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

	// StaticFS, err := fs.Sub(assets.EmbeddedFiles, "static")
	// if err != nil {
	// 	panic(err)
	// }
	// fs := http.FileServer(http.FS(StaticFS))
	// mux.Handle("/static/", http.StripPrefix("/static/", fs))

	todoService := todos.NewTodoService(deps.TodoStore, deps.Logger)
	todoService.RegisterRoutes(mux)

	loggingMiddleware := middleware.LogRequest(deps.Logger)

	server := loggingMiddleware(mux)
	return server
}
