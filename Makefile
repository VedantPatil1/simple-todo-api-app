vet:
	go vet ./...

fmt: vet
	go vet ./...

test: fmt
	go test ./...

templ:
	templ generate

build: templ
	go build -o ./bin/todo-app cmd/app/main.go

run: build
	./bin/todo-app
