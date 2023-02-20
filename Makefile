
.PHONY: all test clean build

all: clean test build

install:
	go install cmd/sleep/sleep.go

run:
	go run *.go -tags buildChecks

clean:
	go clean -cache -testcache -modcache
	rm -rf bin/