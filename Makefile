GO_FILES=$(shell find . -iname '*.go')
bin/quickview: $(GO_FILES) ./pkg/server/templates/* .git/hooks/pre-commit
	go build -o bin/quickview ./cmd/quickview/main.go

# Format everything
.PHONY: fmt
fmt:
	go fmt ./...
	npx prettier . --write

# Prettier installs node_modules
node_modules: package.json package-lock.json
	npm install

# Run prettier on pre-commit
.git/hooks/pre-commit:
	cp .evertras/pre-commit.sh .git/hooks/pre-commit
	chmod +x .git/hooks/pre-commit
