package ui

import (
	"fmt"
	"strings"
	"sync"
	"time"
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

	s := NewProgressSpinner()
	s.Start(title)

	err := action()

	if err != nil {
		s.Fail(title)
	} else {
		s.Complete(title)
	}

	return err
}

// RunWithSpinnerSimple runs a simple action (no error return) with a spinner
func RunWithSpinnerSimple(title string, action func()) error {
	return RunWithSpinner(title, func() error {
		action()
		return nil
	})
}

// ProgressSpinner provides a live-updating spinner with progress info
type ProgressSpinner struct {
	frames   []string
	index    int
	message  string
	mu       sync.Mutex
	stopCh   chan struct{}
	doneCh   chan struct{}
	interval time.Duration
}

// NewProgressSpinner creates a new progress spinner
func NewProgressSpinner() *ProgressSpinner {
	return &ProgressSpinner{
		frames:   []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
		interval: 80 * time.Millisecond,
		stopCh:   make(chan struct{}),
		doneCh:   make(chan struct{}),
	}
}

// Start begins the spinner animation
func (p *ProgressSpinner) Start(initialMsg string) {
	p.message = initialMsg

	if !IsTTY() {
		// Non-TTY: just print the message
		fmt.Println(initialMsg)
		return
	}

	go func() {
		defer close(p.doneCh)
		ticker := time.NewTicker(p.interval)
		defer ticker.Stop()

		for {
			select {
			case <-p.stopCh:
				return
			case <-ticker.C:
				p.mu.Lock()
				frame := p.frames[p.index]
				msg := p.message
				p.index = (p.index + 1) % len(p.frames)
				p.mu.Unlock()

				// Clear line and print spinner + message
				fmt.Printf("\r\033[K%s %s", Primary.Render(frame), msg)
			}
		}
	}()
}

// Update changes the spinner message
func (p *ProgressSpinner) Update(msg string) {
	p.mu.Lock()
	p.message = msg
	p.mu.Unlock()

	if !IsTTY() {
		// Non-TTY: print each update on its own line
		fmt.Println(msg)
	}
}

// Complete stops the spinner and prints a success message
func (p *ProgressSpinner) Complete(msg string) {
	if IsTTY() {
		close(p.stopCh)
		<-p.doneCh
		// Clear spinner line and print completion
		fmt.Printf("\r\033[K%s %s\n", Success.Render("✓"), msg)
	} else {
		fmt.Printf("%s %s\n", Success.Render("✓"), msg)
	}
}

// Fail stops the spinner and prints a failure message
func (p *ProgressSpinner) Fail(msg string) {
	if IsTTY() {
		close(p.stopCh)
		<-p.doneCh
		fmt.Printf("\r\033[K%s %s\n", Error.Render("✗"), msg)
	} else {
		fmt.Printf("%s %s\n", Error.Render("✗"), msg)
	}
}

// PrintAboveSpinner prints a line above the spinner without stopping it
// Used to show completion of one phase before moving to next
func (p *ProgressSpinner) PrintAboveSpinner(msg string) {
	if IsTTY() {
		// Clear current line, print message, then spinner will redraw on next tick
		fmt.Printf("\r\033[K%s\n", msg)
	} else {
		fmt.Println(msg)
	}
}

// Stop halts the spinner without printing anything
func (p *ProgressSpinner) Stop() {
	if IsTTY() {
		close(p.stopCh)
		<-p.doneCh
		fmt.Print("\r\033[K") // Clear line
	}
}

// FormatProgress returns a formatted progress string like "[12/47]"
func FormatProgress(current, total int) string {
	return Dim.Render(fmt.Sprintf("[%d/%d]", current, total))
}

// FormatProgressWithLabel returns "[12/47] label"
func FormatProgressWithLabel(current, total int, label string) string {
	// Truncate label if too long
	maxLen := 30
	displayLabel := label
	if len(label) > maxLen {
		displayLabel = label[:maxLen-3] + "..."
	}
	return fmt.Sprintf("%s %s", FormatProgress(current, total), displayLabel)
}

// FormatDraftProgress returns "draft [12/47] keyword"
func FormatDraftProgress(draft string, current, total int, keyword string) string {
	parts := []string{Bold.Render(draft)}
	if total > 0 {
		parts = append(parts, FormatProgress(current, total))
	}
	if keyword != "" {
		parts = append(parts, Dim.Render(keyword))
	}
	return strings.Join(parts, " ")
}
