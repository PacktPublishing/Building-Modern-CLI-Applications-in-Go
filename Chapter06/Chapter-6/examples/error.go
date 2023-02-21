package examples

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func HandlingDoesNotExistErrors() {
	cmd := exec.Command("doesnotexist", "arg1")
	if errors.Is(cmd.Err, exec.ErrDot) {
		fmt.Println("path lookup resolved to a local directory")
	}
	if err := cmd.Run(); err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			fmt.Println("executable failed to resolve")
		}
	}
}

func HandlingOtherErrors() {
	cmd := exec.Command(filepath.Join(os.Getenv("GOPATH"), "bin", "error"))
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	}
}
