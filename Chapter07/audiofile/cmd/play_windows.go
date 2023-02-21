//go:build windows

package cmd

import (
	"fmt"
	"os/exec"
)

func play(audiofilePath string) error {
	cmd := exec.Command("cmd", "/C", "start", audiofilePath)
	if err := cmd.Start(); err != nil {
		return err
	}
	fmt.Println("enjoy the music!")
	err := cmd.Wait()
	if err != nil {
		return err
	}
}
