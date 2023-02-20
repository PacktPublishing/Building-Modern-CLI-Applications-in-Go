
.PHONY: all test clean

all: clean test

build-all: build-darwin-free build-darwin-pro build-darwin-pro-profile build-linux-free build-linux-pro build-linux-pro-profile build-windows-free build-windows-pro build-windows-pro-profile

build-darwin-amd64-free:
	GOOS=darwin GOARCH=amd64 go build -tags "darwin free" -o builds/free/darwin/audiofile main.go
	chmod +x builds/free/darwin/audiofile

install-linux-amd64-free:
    GOOS=linux GOARCH=amd64 go install -tags "linux free" github.com/marianina8/audiofile 

build-darwin-free:
	GOOS=darwin GOARCH=amd64 go build -tags "darwin free" -o builds/free/darwin/audiofile main.go
	chmod +x builds/free/darwin/audiofile

build-darwin-pro:
	GOOS=darwin GOARCH=amd64 go build -tags "darwin pro" -o builds/pro/darwin/audiofile main.go
	chmod +x builds/pro/darwin/audiofile

build-darwin-pro-profile:
	GOOS=darwin GOARCH=amd64 go build -tags "darwin pro profile" -o builds/profile/darwin/audiofile main.go
	chmod +x builds/profile/darwin/audiofile

build-linux-free:
	GOOS=linux go build -tags "linux free" -o builds/free/linux/audiofile main.go
	chmod +x builds/free/linux/audiofile

build-linux-pro:
	GOOS=linux go build -tags "linux pro" -o builds/pro/linux/audiofile main.go
	chmod +x builds/pro/linux/audiofile

build-linux-pro-profile:
	GOOS=linux go build -tags "linux pro profile" -o builds/profile/linux/audiofile main.go
	chmod +x builds/profile/linux/audiofile

build-windows-free:
	GOOS=windows go build -tags "windows free" -o builds/free/windows/audiofile.exe main.go

build-windows-pro:
	GOOS=windows go build -tags "windows pro" -o builds/pro/windows/audiofile.exe main.go

build-windows-pro-profile:
	GOOS=windows go build -tags "windows pro profile" -o builds/profile/windows/audiofile.exe main.go

install-darwin-free:
	go install -tags "darwin free" github.com/marianina8/audiofile

install-darwin-pro:
	go install -tags "darwin pro" github.com/marianina8/audiofile

install-darwin-pro-profile:
	go install -tags "darwin pro profile" github.com/marianina8/audiofile

install-linux-free:
	GOOS=linux go install -tags "linux free" github.com/marianina8/audiofile

install-linux-pro:
	GOOS=linux go install -tags "linux pro" github.com/marianina8/audiofile

install-linux-pro-profile:
	GOOS=linux go install -tags "linux pro profile" github.com/marianina8/audiofile

install-windows-free:
	GOOS=windows go install -tags "windows free" github.com/marianina8/audiofile

install-windows-pro:
	GOOS=windows go install -tags "windows pro" github.com/marianina8/audiofile

install-windows-pro-profile:
	GOOS=windows go install -tags "windows pro profile" github.com/marianina8/audiofile

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