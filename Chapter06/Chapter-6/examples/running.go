//go:build darwin
// +build darwin

package examples

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func RunMethod() {
	cmd := exec.Command(filepath.Join(os.Getenv("GOPATH"), "bin", "lettercount"), "four")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	var count int
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	cmd.ExtraFiles = []*os.File{writer}
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	if err := json.NewDecoder(reader).Decode(&count); err != nil {
		panic(err)
	}
	fmt.Println("letter count: ", count)
}

func StartMethod() {
	cmd := exec.Command(filepath.Join(os.Getenv("GOPATH"), "bin", "lettercount"), "four")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	var count int
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	cmd.ExtraFiles = []*os.File{writer}
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	err = cmd.Wait()
	if err != nil {
		panic(err)
	}
	if err := json.NewDecoder(reader).Decode(&count); err != nil {
		panic(err)
	}
	fmt.Println("letter count: ", count)
}

func OutputMethod() {
	cmd := exec.Command(filepath.Join(os.Getenv("GOPATH"), "bin", "lettercount"), "four")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	var count int
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	cmd.ExtraFiles = []*os.File{writer}
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Printf("stdout --> %v\n", string(out))
	if err := json.NewDecoder(reader).Decode(&count); err != nil {
		panic(err)
	}
	fmt.Println("letter count: ", count)
}

func CombinedOutputMethod() {
	cmd := exec.Command(filepath.Join(os.Getenv("GOPATH"), "bin", "lettercount"), "four")
	cmd.Stdin = os.Stdin
	var count int
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	cmd.ExtraFiles = []*os.File{writer}
	CombinedOutput, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Printf("combined stderr + stdout --> ****\n%v***\n", string(CombinedOutput))
	if err := json.NewDecoder(reader).Decode(&count); err != nil {
		panic(err)
	}
	fmt.Println("letter count: ", count)
}
