default: vet test

test:
	go test -cover -v ./...

vet:
	go vet -v ./...

.PHONY: test vet
