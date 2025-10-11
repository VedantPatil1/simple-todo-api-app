vet:
	go vet ./...

fmt: vet
	go vet ./...

test: fmt
	go test ./...

build:
	go build -o ./bin/todo-api-app cmd/api/main.go

run: build
	./bin/todo-api-app
