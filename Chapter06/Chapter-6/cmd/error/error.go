package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 0 { // not passing in any arguments in this example throws an error
		fmt.Fprintf(os.Stderr, "missing arguments\n")
		os.Exit(1)
	}
	fmt.Println("executing command with no errors")
}
