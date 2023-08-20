GO_FILES=$(shell find . -iname '*.go')
bin/quickview: $(GO_FILES)
	go build -o bin/quickview ./cmd/quickview/main.go

.PHONY: fmt
fmt:
	go fmt ./...
	npx prettier . --write

node_modules: package.json package-lock.json
	npm install
