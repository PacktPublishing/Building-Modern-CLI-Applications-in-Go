//go:build windows

package cmd

import (
	"github.com/marianina8/audiofile/utils"
	"os/exec"
)

func play(audiofilePath string) error {
	cmd := exec.Command("cmd", "/C", "start", audiofilePath)
	if err := cmd.Start(); err != nil {
		return err
	}
	fmt.Println("Enjoy the music...")
	err := cmd.Wait()
	if err != nil {
		return err
	}
	return nil
}
