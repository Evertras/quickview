GO_FILES=$(shell find . -iname '*.go')
bin/quickview: $(GO_FILES) .git/hooks/pre-commit
	go build -o bin/quickview ./cmd/quickview/main.go

.PHONY: fmt
fmt:
	go fmt ./...
	npx prettier . --write

node_modules: package.json package-lock.json
	npm install

.git/hooks/pre-commit:
	cp .evertras/pre-commit.sh .git/hooks/pre-commit
	chmod +x .git/hooks/pre-commit
