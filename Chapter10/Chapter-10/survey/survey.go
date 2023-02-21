package survey

import (
	"fmt"

	survey "github.com/AlecAivazis/survey/v2"
)

// UserExperience is a survey that
func UserExperience() {
	questions := []*survey.Question{
		{
			Name:      "email",
			Prompt:    &survey.Input{Message: "What is your email addres?"},
			Validate:  survey.Required,
			Transform: survey.Title,
		},
		{
			Name: "rating",
			Prompt: &survey.Select{
				Message: "How would you rate your experience with the CLI?",
				Options: []string{"Hated it", "Disliked", "Decent", "Great", "Loved it"},
			},
		},
		{
			Name: "issues",
			Prompt: &survey.MultiSelect{
				Message: "Have you encountered any of these issues?",
				Options: []string{"audio player issues", "upload issues", "search issues", "other technical issues"},
			},
		},
		{
			Name: "suggestions",
			Prompt: &survey.Multiline{
				Message: "Please provide any other feedback or suggestions you may have.",
			},
		},
	}

	results := struct {
		Email       string
		Rating      string
		Issues      []string
		Suggestions string
	}{}

	err := survey.Ask(questions, &results)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
