//go:build darwin
// +build darwin

package examples

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Pagination() {
	pagesCmd := exec.Command(filepath.Join(os.Getenv("GOPATH"), "bin", "pages"))
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	pagesCmd.Stdin = os.Stdin
	pagesCmd.Stdout = os.Stdout
	pagesCmd.Stderr = os.Stderr
	pagesCmd.ExtraFiles = []*os.File{writer}
	if err := pagesCmd.Run(); err != nil {
		panic(err)
	}
	var data string
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&data); err != nil {
		panic(err)
	}
	lessCmd := exec.Command("/usr/bin/less")
	lessCmd.Stdin = strings.NewReader(data)
	lessCmd.Stdout = os.Stdout
	err = lessCmd.Run()
	if err != nil {
		panic(err)
	}
}
