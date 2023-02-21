//go:build linux

package cmd

import (
	"fmt"
	"github.com/marianina8/audiofile/utils"
	"github.com/pterm/pterm"
	"os/exec"
)

func play(audiofilePath string, verbose bool) (int, error) {
    cmd := exec.Command("aplay", audiofilePath)
    if err := cmd.Start(); err != nil {
        return 0, utils.Error("\n  starting aplay command: %v", err, verbose)
    }
    spinnerInfo := &pterm.SpinnerPrinter{}
    if utils.IsaTTY() {
        spinnerInfo, _ = pterm.DefaultSpinner.Start("Enjoy the music...")
    }
    err := cmd.Wait()
    if err != nil {
        return 0, utils.Error("\n  running aplay command: %v", err, verbose)
    }
    if utils.IsaTTY() {
        spinnerInfo.Stop()
    }
    return cmd.Process.Pid, nil
}
