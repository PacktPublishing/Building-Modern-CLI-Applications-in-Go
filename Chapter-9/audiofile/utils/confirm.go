package utils

import (
	survey "github.com/AlecAivazis/survey/v2"
)

func Confirm(confirmationText string) bool {
	confirmed := false
	prompt := &survey.Confirm{
		Message: confirmationText,
	}
	survey.AskOne(prompt, &confirmed)
	return confirmed
}
