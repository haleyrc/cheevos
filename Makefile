all: vet static build test

.PHONY: vet
vet:
	echo Running go vet...
	@go vet ./...

.PHONY: static
static:
	echo Running staticcheck...
	@staticcheck ./...

.PHONY: build
build:
	echo Building...
	@go build ./...

.PHONY: test
test:
	echo Running tests...
	@go test ./...

.PHONY: build-server
build-server:
	go build -o bin/server ./cmd/server

.PHONY: server
server: build-server
	./bin/server
