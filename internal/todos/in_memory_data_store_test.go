package todos

import (
	"reflect"
	"testing"
)

func TestGetTodos(t *testing.T) {
	todo := Todo{
		Id:          1,
		Title:       "test todo",
		IsCompleted: false,
	}
	store := InMemoryDataStore{
		Todos: map[int]Todo{1: todo},
	}

	retrieved_todo := store.GetTodos()
	if !reflect.DeepEqual(retrieved_todo, []Todo{todo}) {
		t.Errorf("error: got %v want %v", retrieved_todo, []Todo{todo})
	}
}

func TestAddTodos(t *testing.T) {
	store := InMemoryDataStore{
		Todos: make(map[int]Todo),
	}

	id := store.AddTodo("test todo")

	if id != 0 {
		t.Errorf("generated wrong id: got %d want %d", id, 0)
	}
	got := store.Todos[0]
	want := Todo{
		Id:          0,
		Title:       "test todo",
		IsCompleted: false,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("error: got %v want %v", got, want)
	}
}

func TestGetById(t *testing.T) {
	store := InMemoryDataStore{
		Todos: make(map[int]Todo),
	}

	t.Run("get todo by id", func(t *testing.T) {
		store.AddTodo("test todo")

		got, _ := store.GetById(0)
		want := Todo{
			Id:          0,
			Title:       "test todo",
			IsCompleted: false,
		}

		if got != want {
			t.Errorf("error: got %v want %v", got, want)
		}
	})
	t.Run("error for id not found", func(t *testing.T) {
		_, got := store.GetById(2)
		want := ErrTodoNotFound

		if got != want {
			t.Errorf("error not returned: got %v want %v", got, want)
		}
	})
}
