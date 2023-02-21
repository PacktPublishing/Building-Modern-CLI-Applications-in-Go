package main

import (
	"encoding/json"
	"os"
	"strings"
)

func main() {
	input := os.Args[1:]
	output := strings.ToUpper(strings.Join(input, ""))
	pipe := os.NewFile(uintptr(3), "pipe")
	err := json.NewEncoder(pipe).Encode(output)
	if err != nil {
		panic(err)
	}
}
