package todos

import (
	"log/slog"
	"net/http"
	"strconv"
)

type TodoStore interface {
	GetTodos() []Todo
	AddTodo(string) Todo
	GetById(int) (Todo, error)
	ToggleCompletion(int) (Todo, error)
}

type TodoService struct {
	logger *slog.Logger
	store  TodoStore
}

func NewTodoService(s TodoStore, l *slog.Logger) *TodoService {
	return &TodoService{store: s, logger: l}
}

func (s *TodoService) TodoToggleHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid todo id", http.StatusBadRequest)
		return
	}

	todo, err := s.store.ToggleCompletion(id)
	if err != nil {
		http.Error(w, "todo not found", http.StatusNotFound)
		return
	}

	s.logger.Info("updated store", "todos", s.store.GetTodos())
	s.logger.Info("toggled todo", "id", id, "isCompleted", todo.IsCompleted)

	TodoItem(todo).Render(r.Context(), w)
}

func (s *TodoService) IndexPage(w http.ResponseWriter, r *http.Request) {
	todos := s.store.GetTodos()
	s.logger.Info("fetchecd todos", "todos", todos)

	IndexView(todos).Render(r.Context(), w)
}
func (s *TodoService) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("GET /todos/", http.HandlerFunc(s.IndexPage))
	mux.Handle("PATCH /todos/{id}/toggle", http.HandlerFunc(s.TodoToggleHandler))
}
