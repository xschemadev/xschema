package ui

import (
	"fmt"

	"github.com/charmbracelet/huh/spinner"
)

// SpinnerAction runs an action with a spinner, returning any error from the action
type SpinnerAction func() error

// RunWithSpinner runs an action with a spinner display
// If not TTY, just prints the title and runs the action
func RunWithSpinner(title string, action SpinnerAction) error {
	if !IsTTY() {
		// Non-TTY: just print and run
		fmt.Println(title)
		return action()
	}

	var actionErr error
	spinErr := spinner.New().
		Title(title).
		Action(func() {
			actionErr = action()
		}).
		Run()

	if spinErr != nil {
		return spinErr
	}
	return actionErr
}

// RunWithSpinnerSimple runs a simple action (no error return) with a spinner
func RunWithSpinnerSimple(title string, action func()) error {
	return RunWithSpinner(title, func() error {
		action()
		return nil
	})
}
