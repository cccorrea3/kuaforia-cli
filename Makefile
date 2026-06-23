BINARY_NAME=kuaforia
VERSION=$(shell git describe --tags 2>/dev/null || echo "dev")
LDFLAGS=-ldflags "-X main.Version=$(VERSION)"

build:
	go build $(LDFLAGS) -o bin/$(BINARY_NAME) .

install:
	go install $(LDFLAGS) .

test:
	go test ./... -v

lint:
	golangci-lint run ./...

clean:
	rm -rf bin/
