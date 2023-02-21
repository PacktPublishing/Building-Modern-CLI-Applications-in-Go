//go:build windows

package examples

import (
	"fmt"
	"os"
	"os/exec"
)

func RunMethod() {
	cmd := exec.Command("cmd", "/c", "ping google.com")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func StartMethod() {
	cmd := exec.Command("cmd", "/c", "ping google.com")
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
	err = cmd.Wait()
	if err != nil {
		panic(err)
	}
}

func OutputMethod() {
	cmd := exec.Command("cmd", "/c", "ping google.com")
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
}

func CombinedOutputMethod() {
	output, err := exec.Command("cmd", "/c", "ping google.com").CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
}
