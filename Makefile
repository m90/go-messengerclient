default: test

test:
	go test -cover -v ./...

.PHONY: test
