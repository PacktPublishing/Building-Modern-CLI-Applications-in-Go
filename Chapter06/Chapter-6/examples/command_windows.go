//go:build windows

package examples

import (
	"fmt"
	"os/exec"
)

func CreateCommandUsingStruct() {
	cmd := exec.Command("cmd", "/C", "ping")
	cmd.Args = []string{"google"}
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
}

func CreateCommandUsingCommandFunction() {
	cmd := exec.Command("cmd", "/C", "ping", "google.com")
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
}
