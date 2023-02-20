//go:build darwin

package cmd

import (
	"os/exec"

	"github.com/marianina8/audiofile/utils"
	"github.com/pterm/pterm"
)

func play(audiofilePath string, verbose bool) (int, error) {
	cmd := exec.Command("afplay", audiofilePath)
	if err := cmd.Start(); err != nil {
		return 0, utils.Error("\n  starting afplay command: %v", err, verbose)
	}
	spinnerInfo := &pterm.SpinnerPrinter{}
	if utils.IsaTTY() {
		spinnerInfo, _ = pterm.DefaultSpinner.Start("Enjoy the music...")
	}
	err := cmd.Wait()
	if err != nil {
		return 0, utils.Error("\n  running afplay command: %v", err, verbose)
	}
	if utils.IsaTTY() {
		spinnerInfo.Stop()
	}
	return cmd.Process.Pid, nil
}
