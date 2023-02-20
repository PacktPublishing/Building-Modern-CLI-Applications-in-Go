# Chapter-6
Building Modern CLIs with Go - Chapter 6

## Prerequisites
These applications were build to run on a UNIX operating system.  Make sure your GOPATH and GOROOT are properly set!

## Installing on MacOS
To install the executables used for the commands that exist within the `/cmd` folder then run:
`make install`

## Installing on Windows
To install the executables used for the commands that exist within the `/cmd` folder then run:
`.\build-windows.ps1`

## Running
To see all examples run, simply run `go run main.go` All of the following examples will be run.  However to follow along with the book, it's probably a better idea to comment out each individually to focus on one example at a time.

```
func main() {
	examples.CreateCommandUsingStruct()
	examples.CreateCommandUsingCommandFunction()
	examples.RunMethod()
	examples.StartMethod()
	examples.OutputMethod()
	examples.CombinedOutputMethod()
	examples.Pagination()
	examples.Limit()
	examples.Timeout()
	examples.HandlingDoesNotExistErrors()
	examples.HandlingOtherErrors()
	examples.Panic()
	examples.HTTPTimeout()
	examples.HTTPError()
}
```