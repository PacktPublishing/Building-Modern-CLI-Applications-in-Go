package main

import (
	"flag"
	"fmt"
)

func main() {
	var helloFlag bool
	flag.BoolVar(&helloFlag, "hello", false, "Print 'Hello, World!'")
	flag.Parse()
	if helloFlag {
		fmt.Println("Hello, World!")
	}
}
