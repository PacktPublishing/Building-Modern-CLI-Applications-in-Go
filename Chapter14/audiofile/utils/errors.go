package utils

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/fatih/color"
)

var errorColor = color.New(color.BgRed, color.FgWhite).SprintFunc()

func cleanup(errString string, err error) string {
	errString = strings.Map(func(r rune) rune {
		if unicode.IsGraphic(r) {
			return r
		}
		return -1
	}, errString)
	return fmt.Sprintf(strings.Replace(errString, "\n  ", "", 1), err.Error())

}

func Error(errString string, err error, verbose bool) error {
	errString = cleanup(errString, err)
	if err != nil {
		if verbose {
			// prints to stdout also
			Verbose.Error(errString)
		} else {
			Logger.Error(errString)
		}
		return fmt.Errorf(errString)
	}
	return nil
}
