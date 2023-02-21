package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		panic(fmt.Errorf("[Err]: expected 2 args received %d", len(os.Args)))
	}
	input := os.Args[1]
	runes := []rune(input)
	pipe := os.NewFile(uintptr(3), "pipe")
	err := json.NewEncoder(pipe).Encode(len(runes))
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(os.Stderr, "this is where the errors go")
	fmt.Printf("successfully counted the letters of \"%v\" as %d\n", input, len(runes))
}
