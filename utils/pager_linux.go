//go:build linux

package utils

import (
	"os"
	"os/exec"
	"strings"
)

func Pager(data string) error {
	lessCmd := exec.Command("less", "-r")
	lessCmd.Stdin = strings.NewReader(data)
	lessCmd.Stdout = os.Stdout
	lessCmd.Stderr = os.Stderr
	err := lessCmd.Run()
	if err != nil {
		return err
	}
	return nil
}
