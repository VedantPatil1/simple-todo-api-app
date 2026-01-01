package todos

import (
	"net/http"

	"github.com/a-h/templ"
)

type TodoStore interface {
	GetTodos() []Todo
	AddTodo(string) Todo
	GetById(int) (Todo, error)
	ToggleCompletion(int) (Todo, error)
}

type TodoService struct {
	store TodoStore
}

func NewTodoService(s TodoStore) *TodoService {
	return &TodoService{store: s}
}

func (s *TodoService) RegisterRoutes(mux *http.ServeMux) {
	s.store.AddTodo("test todo 1")
	todo := s.store.AddTodo("test todo 2")
	s.store.ToggleCompletion(todo.Id)
	mux.Handle("GET /todos/", templ.Handler(IndexView(s.store.GetTodos())))
}
