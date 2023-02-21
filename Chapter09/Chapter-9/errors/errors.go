package errors

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/pkg/errors"
)

const (
	checkMark    = "\U00002705"
	crossMark    = "\U0000274C"
	votingBallot = "\U0001F5F3"
)

type customError struct {
	Task string
	Err  error
}

func (e *customError) Error() string {
	var errorColor = color.New(color.BgRed, color.FgWhite).SprintFunc()
	return fmt.Sprintf("%s: %s %s", errorColor(e.Task), crossMark, e.Err)
}

func Examples() {
	// using fmt.Errorf
	birthYear := -1981
	err := fmt.Errorf("%d is negative\nYear can't be negative", birthYear)
	if birthYear < 0 {
		fmt.Println(err)
	} else {
		fmt.Printf("Birth year: %d\n", birthYear)
	}
	birthYear = 2010
	currentYear := 2022
	age := currentYear - birthYear
	// using errors.Wrap
	err = wrapping()
	fmt.Printf("%+v\n", err)
	// custom error(s)
	err = eligibleToVote(age)
	if err != nil {
		fmt.Println("error occurred: ", err)
	}
	fmt.Println()
	err = eligibleToVote(21)
	if err != nil {
		fmt.Println("error occurred: ", err)
	}
}

func eligibleToVote(age int) error {
	fmt.Printf("%s Attempting to vote at %d years old...\n", votingBallot, age)
	minimumAge := 18
	err := &customError{
		Task: "eligibleToVote",
	}
	if age < minimumAge && age > 0 {
		years := minimumAge - age
		err.Err = fmt.Errorf("too young to vote, at %d, wait %d more years", age, years)
		return err
	}
	if age < 0 {
		err.Err = fmt.Errorf("age cannot be negative: %d", age)
		return err
	}
	fmt.Println("Voted.", checkMark)
	return nil
}

func wrapping() error {
	err := errors.New("error")
	err1 := operation1()
	if err1 != nil {
		err1 = errors.Wrap(err, "operation1")
	}
	err2 := operation2()
	if err != nil {
		err2 = errors.Wrap(err1, "operation2")
	}
	err3 := operation3()
	if err != nil {
		err3 = errors.Wrap(err2, "operation3")
	}
	return err3
}

func operation1() error {
	return fmt.Errorf("random error message 1")
}

func operation2() error {
	return fmt.Errorf("random error message 2")
}

func operation3() error {
	return fmt.Errorf("random error message 3")
}
