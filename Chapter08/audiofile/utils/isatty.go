package utils

import (
	"os"

	isatty "github.com/mattn/go-isatty"
)

func IsaTTY() bool {
	if isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd()) {
		return true
	}
	return false
}
