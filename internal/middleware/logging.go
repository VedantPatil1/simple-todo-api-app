package middleware

import (
	"log/slog"
	"net/http"

	"github.com/VedantPatil1/simple-todo-api-app.git/internal/response"
)

func LogRequest(l *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			l.Info("request started", "method", r.Method, "path", r.URL.Path)
			rec := response.NewStatusRecorder(w)
			next.ServeHTTP(rec, r)
			l.Info("request completed", "status", rec.StatusCode)

		})
	}

}
