package todos

import (
	"context"
	"fmt"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/a-h/templ"
)

func TestTodoItemComponent(t *testing.T) {
	t.Run("renders incomplete state with correct root attributes", func(t *testing.T) {
		todo := Todo{
			Id:          1,
			Title:       "Test Todo",
			IsCompleted: false,
		}
		doc := renderToDoc(t, TodoItem(todo))

		root := doc.Find("[data-testid='TodoItem']")
		if root.Length() == 0 {
			t.Fatal("expected root data-testid='TodoItem' to be present")
		}

		got, _ := root.Attr("data-todo-id")
		want := fmt.Sprintf("%d", todo.Id)
		if got != fmt.Sprintf("%d", todo.Id) {
			t.Errorf("wrong id attribute: got %v, want %s", got, want)
		}

		if title := root.Find(".todo-title").Text(); title != todo.Title {
			t.Errorf("expected title '%s', got '%s'", todo.Title, title)
		}

		input := root.Find("input[type='checkbox']")
		if _, isChecked := input.Attr("checked"); isChecked {
			t.Error("expected checkbox to be unchecked")
		}
	})

	t.Run("renders checked state for completed todo", func(t *testing.T) {
		todo := Todo{Id: 1, IsCompleted: true}
		doc := renderToDoc(t, TodoItem(todo))

		input := doc.Find("input[type='checkbox']")
		if _, isChecked := input.Attr("checked"); !isChecked {
			t.Error("expected checkbox to be checked")
		}
	})
}

func TestTodoListComponent(t *testing.T) {
	todos := []Todo{
		{Id: 1, Title: "test 1"},
		{Id: 2, Title: "test 2"},
	}
	doc := renderToDoc(t, TodoList(todos))

	comp := doc.Find("[data-testid=TodoList]")
	if comp.Length() == 0 {
		t.Error("component id 'TodoList' not found")
	}

	items := comp.Find("[data-testid=TodoItem]")
	if nItems := items.Length(); nItems != 2 {
		t.Errorf("error number of components of id 'TodoItem': got %d want 2", nItems)
	}

	for _, todo := range todos {
		got := items.Filter(fmt.Sprintf("[data-todo-id='%d']", todo.Id))
		if got.Length() != 1 {
			t.Errorf("error items with todo-id %d: got %d want 1", todo.Id, got.Length())
		}
	}
}

func renderToDoc(t *testing.T, component templ.Component) *goquery.Document {
	r, w := io.Pipe()
	go func() {
		_ = component.Render(context.Background(), w)
		_ = w.Close()
	}()

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Errorf("failed to read template: %v", err)
	}

	return doc
}
