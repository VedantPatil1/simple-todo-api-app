package todos

import (
	"errors"
	"sync"
)

var ErrTodoNotFound = errors.New("error todo not found for given id")

type InMemoryDataStore struct {
	mu     sync.RWMutex
	Todos  map[int]Todo
	nextId int
}

func (s *InMemoryDataStore) GetTodos() []Todo {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var todos []Todo
	for _, todo := range s.Todos {
		todos = append(todos, todo)
	}
	return todos
}

func (s *InMemoryDataStore) AddTodo(title string) Todo {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := s.nextId
	s.nextId++

	todo := Todo{
		Id:          id,
		Title:       title,
		IsCompleted: false,
	}

	s.Todos[id] = todo

	return todo
}

func (s *InMemoryDataStore) GetById(id int) (Todo, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	todo, ok := s.Todos[id]
	if !ok {
		return Todo{}, ErrTodoNotFound
	}
	return todo, nil
}

func (s *InMemoryDataStore) ToggleCompletion(id int) (Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo, ok := s.Todos[id]
	if !ok {
		return Todo{}, ErrTodoNotFound
	}

	todo.IsCompleted = !todo.IsCompleted

	s.Todos[id] = todo

	return todo, nil
}

func NewInMemoryDataStore() *InMemoryDataStore {
	return &InMemoryDataStore{Todos: make(map[int]Todo)}
}
