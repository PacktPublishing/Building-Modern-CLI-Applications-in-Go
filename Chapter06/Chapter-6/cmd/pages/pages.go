package main

import (
	_ "embed"
	"encoding/json"
	"os"
)

//go:embed data.txt
var contents []byte

func main() {
	pipe := os.NewFile(uintptr(3), "pipe")
	err := json.NewEncoder(pipe).Encode(string(contents))
	if err != nil {
		panic(err)
	}
}
