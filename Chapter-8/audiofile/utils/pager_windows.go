//go:build windows

package utils

import (
	"os"
	"os/exec"
	"strings"
)

func Pager(data string) error {
	moreCmd := exec.Command("cmd", "/C", "more")
	moreCmd.Stdin = strings.NewReader(data)
	moreCmd.Stdout = os.Stdout
	moreCmd.Stderr = os.Stderr
	err := moreCmd.Run()
	if err != nil {
		return err
	}
	return nil
}
