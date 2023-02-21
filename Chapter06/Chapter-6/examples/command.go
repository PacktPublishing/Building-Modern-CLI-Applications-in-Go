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

func CreateCommandUsingStruct() {
	cmd := exec.Cmd{}
	cmd.Path = filepath.Join(os.Getenv("GOPATH"), "bin", "uppercase") // only required field, path of the command to run
	cmd.Args = []string{"uppercase", "hack the planet"}               // represents command and arguments
	cmd.Stdin = os.Stdin                                              // io.Reader
	cmd.Stdout = os.Stdout                                            // io.Writer
	cmd.Stderr = os.Stderr                                            // io.Writer
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	cmd.ExtraFiles = []*os.File{writer} // file descriptor to be inherited by the command!
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	err = cmd.Wait()
	if err != nil {
		panic(err)
	}
	var data string
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&data); err != nil {
		panic(err)
	}
	fmt.Println(data)

}

func CreateCommandUsingCommandFunction() {
	cmd := exec.Command(filepath.Join(os.Getenv("GOPATH"), "bin", "uppercase"), "hello world")
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	cmd.ExtraFiles = []*os.File{writer}
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	var data string
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&data); err != nil {
		panic(err)
	}
	fmt.Println(data)
}
