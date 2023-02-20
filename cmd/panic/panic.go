package main

import (
	"fmt"
	"os"
	"runtime/debug"
)

func main() {
	defer func() {
		if panicMessage := recover(); panicMessage != nil {
			fmt.Fprintf(os.Stderr, "(panic) : %v\n", panicMessage)
			debug.PrintStack()
			os.Exit(1)
		}
	}()
	panic("help!")
}
