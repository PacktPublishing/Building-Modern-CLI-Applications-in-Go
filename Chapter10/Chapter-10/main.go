package main

import (
	"github.com/marianina8/Chapter-10/dashboard"
)

func main() {
	// Guiding users with prompts
	//survey.UserExperience()

	// Building a useful dashboard - Learn about termdash
	err := dashboard.Basic()
	if err != nil {
		panic(err)
	}

	err = dashboard.BinaryTreeWithStyle()
	if err != nil {
		panic(err)
	}
}
