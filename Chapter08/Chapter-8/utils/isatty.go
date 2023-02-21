package utils

import (
	"fmt"
	"os"

	isatty "github.com/mattn/go-isatty"
)

func IsaTTY() {
	if isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd()) {
		fmt.Println("Is a TTY")
	} else {
		fmt.Println("Is not a TTY")
	}
}
