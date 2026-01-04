package todos

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestTodoToggleHandler(t *testing.T) {
	request, _ := http.NewRequest(http.MethodPatch, "/todos/0/toggle", nil)
	response := httptest.NewRecorder()

	store := NewInMemoryDataStore()
	store.AddTodo("test todo 1")

	todoService := NewTodoService(store, slog.Default())

	server := http.NewServeMux()
	todoService.RegisterRoutes(server)

	server.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status %d but got %d", http.StatusOK, response.Code)
	}

	todo, _ := store.GetById(0)

	if !todo.IsCompleted {
		t.Fatalf("expected todo to be completed")
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		t.Fatalf("could not parse response body as html: %v", err)
	}

	checkbox := doc.Find("input[type='checkbox']")
	if _, isChecked := checkbox.Attr("checked"); !isChecked {
		t.Errorf("expected checkbox to be checked after toggle: got %v", isChecked)
	}
}
