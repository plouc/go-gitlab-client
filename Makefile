.PHONY: test run update format install

install:
	go get github.com/stretchr/testify/assert
	go list -f '{{range .Imports}}{{.}} {{end}}' ./... | xargs go get -v
	go list -f '{{range .TestImports}}{{.}} {{end}}' ./... | xargs go get -v

update:
	go get -u all

format:
	gofmt -l -w -s .
	go fix ./...

test:
	go test -v ./...
	go vet ./...
	exit `gofmt -l -s -e . | wc -l`
