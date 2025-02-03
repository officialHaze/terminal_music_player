package terminal

import (
	"os"

	"golang.org/x/term"
)

func GetTerminalWidth() int {
	// Get terminal width dynamically
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 80 // Default fallback width
	}
	return int(float32(width) / 1.09);
}