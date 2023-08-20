GO_FILES=$(shell find . -iname '*.go')
bin/quickview: $(GO_FILES)
	go build -o bin/quickview ./cmd/quickview/main.go
