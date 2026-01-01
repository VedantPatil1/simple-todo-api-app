package todos

import (
	"context"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/a-h/templ"
)

func TestTodoItemComponent(t *testing.T) {
	t.Run("test title and check box in component", func(t *testing.T) {
		todo := Todo{
			Id:          1,
			Title:       "test todo",
			IsCompleted: false,
		}
		doc := renderToDoc(t, TodoItem(todo))

		titleText := doc.Find(".todo-title").Text()
		if titleText != todo.Title {
			t.Errorf("error: got %v expected %s", titleText, todo.Title)
		}

		input := doc.Find("input[type='checkbox']")
		if input.Length() == 0 {
			t.Fatal("expected checkbox to be rendered")
		}
	})
	t.Run("test checkbox for incomplete", func(t *testing.T) {
		todo := Todo{
			Id:          1,
			Title:       "test todo",
			IsCompleted: false,
		}
		doc := renderToDoc(t, TodoItem(todo))

		input := doc.Find("input[type='checkbox']")
		if _, isChecked := input.Attr("checked"); isChecked {
			t.Fatal("expected todo to be unchecked for incomplete todo")
		}
	})
	t.Run("test checkbox for completeed", func(t *testing.T) {
		todo := Todo{
			Id:          1,
			Title:       "test todo",
			IsCompleted: true,
		}
		doc := renderToDoc(t, TodoItem(todo))

		input := doc.Find("input[type='checkbox']")
		if _, isChecked := input.Attr("checked"); !isChecked {
			t.Fatal("expected todo to be checked for incomplete todo")
		}
	})
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
