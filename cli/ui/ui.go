package ui

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
	"golang.org/x/term"
)

var (
	isTTY   bool
	verbose bool

	// Colors (bun-style)
	cyan   = lipgloss.Color("6")
	green  = lipgloss.Color("2")
	red    = lipgloss.Color("1")
	yellow = lipgloss.Color("3")
	dim    = lipgloss.Color("8")

	// Styles - exported for use in other packages
	Primary = lipgloss.NewStyle().Foreground(cyan)
	Success = lipgloss.NewStyle().Foreground(green)
	Error   = lipgloss.NewStyle().Foreground(red)
	Warning = lipgloss.NewStyle().Foreground(yellow)
	Dim     = lipgloss.NewStyle().Foreground(dim)
	Bold    = lipgloss.NewStyle().Bold(true)
)

func init() {
	isTTY = term.IsTerminal(int(os.Stdout.Fd()))
	if !isTTY {
		// Disable colors in non-TTY
		lipgloss.SetColorProfile(termenv.Ascii)
	}
}

// SetVerbose enables/disables verbose mode
func SetVerbose(v bool) {
	verbose = v
}

// IsVerbose returns whether verbose mode is enabled
func IsVerbose() bool {
	return verbose
}

// IsTTY returns whether stdout is a terminal
func IsTTY() bool {
	return isTTY
}

// Step prints a step indicator: [1/5] Parsing client file
func Step(num, total int, msg string) {
	prefix := Dim.Render(fmt.Sprintf("[%d/%d]", num, total))
	fmt.Printf("%s %s\n", prefix, msg)
}

// Detail prints indented secondary info with arrow
func Detail(msg string) {
	fmt.Printf("  %s %s\n", Dim.Render("→"), msg)
}

// Verbose prints a message only in verbose mode (indented, dim)
func Verbose(msg string) {
	if verbose {
		fmt.Printf("  %s %s\n", Dim.Render("→"), Dim.Render(msg))
	}
}

// Verbosef prints a formatted message only in verbose mode
func Verbosef(format string, a ...any) {
	if verbose {
		fmt.Printf("  %s %s\n", Dim.Render("→"), Dim.Render(fmt.Sprintf(format, a...)))
	}
}

// SuccessMsg prints a success message with checkmark
func SuccessMsg(msg string) {
	fmt.Printf("%s %s\n", Success.Render("✓"), msg)
}

// ErrorMsg prints an error with formatting and optional hints
func ErrorMsg(title string, err error, hints ...string) {
	fmt.Printf("%s %s\n", Error.Render("✗"), title)
	if err != nil {
		fmt.Printf("  %s\n", Dim.Render(err.Error()))
	}
	for _, hint := range hints {
		fmt.Printf("  %s %s\n", Dim.Render("Hint:"), hint)
	}
}

// WarnMsg prints a warning message
func WarnMsg(msg string) {
	fmt.Printf("%s %s\n", Warning.Render("!"), msg)
}

// FormatDuration formats duration nicely (e.g., "234ms" or "1.2s")
func FormatDuration(d time.Duration) string {
	if d < time.Second {
		return fmt.Sprintf("%dms", d.Milliseconds())
	}
	return fmt.Sprintf("%.1fs", d.Seconds())
}

// FormatBytes formats bytes nicely (e.g., "890B" or "1.2KB")
func FormatBytes(b int) string {
	if b < 1024 {
		return fmt.Sprintf("%dB", b)
	}
	return fmt.Sprintf("%.1fKB", float64(b)/1024)
}

// Println is a simple wrapper for fmt.Println
func Println(a ...any) {
	fmt.Println(a...)
}

// Printf is a simple wrapper for fmt.Printf
func Printf(format string, a ...any) {
	fmt.Printf(format, a...)
}
