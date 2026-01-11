package response

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatusCodeRecorder(t *testing.T) {
	t.Run("default to 200", func(t *testing.T) {
		rr := httptest.NewRecorder()
		rec := NewStatusRecorder(rr)

		assertStatusCode(t, rec.StatusCode, http.StatusOK)
	})
	t.Run("specific code 201", func(t *testing.T) {
		rr := httptest.NewRecorder()
		rec := NewStatusRecorder(rr)
		rec.WriteHeader(http.StatusCreated)

		assertStatusCode(t, rec.StatusCode, http.StatusCreated)
		assertStatusCode(t, rr.Code, http.StatusCreated)
	})
}

func assertStatusCode(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("wrong status code: got %d want %d", got, want)
	}
}
