vet:
	go vet ./...

fmt: vet
	go vet ./...

test: fmt
	go test ./...

build:
	go build -o ./bin/todo-app cmd/app/main.go

run: build
	./bin/todo-app
