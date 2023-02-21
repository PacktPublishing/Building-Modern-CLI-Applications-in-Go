package examples

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func Panic() {
	cmd := exec.Command(filepath.Join(os.Getenv("GOPATH"), "bin", "panic"))
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
}
