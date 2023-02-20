//go:build darwin

package cmd

import (
	"os/exec"

	"github.com/marianina8/audiofile/utils"
	"github.com/pterm/pterm"
)

func play(audiofilePath string) error {
	cmd := exec.Command("afplay", audiofilePath)
	if err := cmd.Start(); err != nil {
		return err
	}
	spinnerInfo := &pterm.SpinnerPrinter{}
	if utils.IsaTTY() {
		spinnerInfo, _ = pterm.DefaultSpinner.Start("Enjoy the music...")
	}
	err := cmd.Wait()
	if err != nil {
		return err
	}
	if utils.IsaTTY() {
		spinnerInfo.Stop()
	}
	return nil
}
