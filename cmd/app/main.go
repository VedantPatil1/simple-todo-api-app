package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/VedantPatil1/simple-todo-api-app.git/internal/middleware"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	slog.Info("Starting App")

	mux := http.NewServeMux()
	mux.HandleFunc("/", greet)

	logMiddleware := middleware.LogRequest(logger)
	server := logMiddleware(mux)

	http.ListenAndServe(":8080", server)
}
