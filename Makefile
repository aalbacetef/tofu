

lint: fmt
	golangci-lint run .

test: fmt
	go test -v .

fmt:
	goimports -w .

.PHONY: test fmt lint 
