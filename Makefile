
.PHONY: all test clean

all: clean test

build-darwin:
	go build -tags darwin -o bin/audiofile main.go
	chmod +x bin/audiofile

clean:
	go clean -cache -testcache -modcache
	rm -rf bin/