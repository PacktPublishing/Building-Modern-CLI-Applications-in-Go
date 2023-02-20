package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

func checkRuntime() {
	fmt.Println("Operating System:", runtime.GOOS)
	fmt.Println("Architecture:", runtime.GOARCH)
	fmt.Println("Go Root:", runtime.GOROOT())
	fmt.Println("Compiler:", runtime.Compiler)
	fmt.Println("No. of CPU:", runtime.NumCPU())
	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())
	fmt.Println("Version:", runtime.Version())
	debug.PrintStack()
}
