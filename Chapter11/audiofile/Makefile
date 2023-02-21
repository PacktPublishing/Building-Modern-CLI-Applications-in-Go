
.PHONY: all test clean

all: clean test

build-darwin-free:
	go build -tags "darwin free" -o bin/audiofile main.go
	chmod +x bin/audiofile

build-darwin-dev:
	go build -tags "darwin dev pro" -o bin/audiofile main.go
	chmod +x bin/audiofile

build-darwin-pro:
	go build -tags "darwin pro" -o bin/audiofile main.go
	chmod +x bin/audiofile

build-darwin-pro-profile:
	go build -tags "darwin pro profile" -o bin/audiofile main.go
	chmod +x bin/audiofile

test:
	go test ./cmd -tags pro

test-verbose:
	go test -v ./cmd -tags pro

manpages:
	mkdir -p pages
	go run documentation/main.go

clean:
	go clean -cache -testcache -modcache
	rm -rf bin/