
.PHONY: all test clean build

all: clean test build

build:
	go build -o bin/audiofile main.go
	chmod +x bin/audiofile

clean:
	go clean -cache -testcache -modcache
	rm -rf bin/