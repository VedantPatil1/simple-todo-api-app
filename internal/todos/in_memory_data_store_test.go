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

	got := store.AddTodo("test todo")

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

func TestToggleCompletion(t *testing.T) {

	t.Run("toggle complete task", func(t *testing.T) {
		store := InMemoryDataStore{
			Todos: make(map[int]Todo),
		}
		store.AddTodo("test todo")

		updated_todo, _ := store.ToggleCompletion(0)

		if !updated_todo.IsCompleted {
			t.Errorf("invalid completion status: got %v want %t", updated_todo.IsCompleted, true)
		}

		updated_todo, _ = store.ToggleCompletion(0)

		if updated_todo.IsCompleted {
			t.Errorf("invalid completion status: got %v want %t", updated_todo.IsCompleted, false)
		}
	})

	t.Run("check error if id not found", func(t *testing.T) {
		store := InMemoryDataStore{
			Todos: make(map[int]Todo),
		}
		store.AddTodo("test todo")

		_, got := store.ToggleCompletion(2)
		want := ErrTodoNotFound

		if got != want {
			t.Errorf("error not returned: got %v want %v", got, want)
		}

	})
}
