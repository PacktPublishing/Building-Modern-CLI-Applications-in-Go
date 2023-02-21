package utils

import (
	"fmt"
	"path/filepath"

	survey "github.com/AlecAivazis/survey/v2"
)

func AskForID() (string, error) {
	id := ""
	prompt := &survey.Input{
		Message: "What is the id of the audiofile?",
	}
	survey.AskOne(prompt, &id)
	if id == "" {
		return "", missingRequiredArumentError("id")
	}
	return id, nil
}

func AskForFilename() (string, error) {
	file := ""
	prompt := &survey.Input{
		Message: "What is the filename of the audio to upload for metadata extraction?",
		Suggest: func(toComplete string) []string {
			files, _ := filepath.Glob(toComplete + "*")
			return files
		},
	}
	survey.AskOne(prompt, &file)
	if file == "" {
		return "", missingRequiredArumentError("file")
	}
	return file, nil
}

func AskForValue() (string, error) {
	value := ""
	prompt := &survey.Input{
		Message: "\U0001F50DWhat value are you searching for?",
	}
	survey.AskOne(prompt, &value)
	if value == "" {
		return "", missingRequiredArumentError("value")
	}
	return value, nil
}

var (
	missingRequiredArumentError = func(missingArg string) error {
		return fmt.Errorf(errorColor(fmt.Sprintf("missing required argument (%s)", missingArg)))
	}
)
