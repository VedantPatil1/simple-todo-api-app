package middleware

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogRequest(t *testing.T) {
	testCases := []struct {
		desc   string
		method string
		path   string
		status int
	}{
		{
			desc:   "simple get request",
			method: http.MethodGet,
			path:   "/test-get",
			status: http.StatusOK,
		},
		{
			desc:   "simple post request",
			method: http.MethodPost,
			path:   "/test-post",
			status: http.StatusCreated,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// Set Listener for logs
			var buf bytes.Buffer
			testLogger := slog.New(slog.NewJSONHandler(&buf, nil))
			oldLogger := slog.Default()
			slog.SetDefault(testLogger)
			defer slog.SetDefault(oldLogger)

			nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tC.status)
				w.Write([]byte("created"))
			})

			mw := LogRequest(testLogger)
			handlerToTest := mw(nextHandler)
			req := httptest.NewRequest(tC.method, tC.path, nil)
			rr := httptest.NewRecorder()

			handlerToTest.ServeHTTP(rr, req)
			if rr.Code != tC.status {
				t.Errorf("handler returned wrong status code. got %v want %v", rr.Code, tC.status)
			}

			logs := parseLogs(t, &buf)

			assertLogField(t, logs[0], "msg", "request started")
			assertLogField(t, logs[0], "method", tC.method)
			assertLogField(t, logs[0], "path", tC.path)

			assertLogField(t, logs[1], "msg", "request completed")
			assertLogField(t, logs[1], "status", tC.status)
		})
	}
}

func parseLogs(t *testing.T, buf *bytes.Buffer) []map[string]any {
	t.Helper()
	var logs []map[string]any
	decoder := json.NewDecoder(buf)

	for decoder.More() {
		var entry map[string]any
		if err := decoder.Decode(&entry); err != nil {
			t.Fatalf("failed to decode log: %v", err)
		}
		logs = append(logs, entry)
	}
	return logs
}

func assertLogField(t *testing.T, log map[string]any, key string, expected any) {
	t.Helper()
	actual, ok := log[key]
	if !ok {
		t.Errorf("missing log field: %s", key)
	}

	if expInt, ok := expected.(int); ok {
		if int(actual.(float64)) != expInt {
			t.Errorf("field %s: got %v want %v", key, actual, expInt)
		}
		return
	}

	if actual != expected {
		t.Errorf("field %s: got %v want %v", key, actual, expected)
	}
}
