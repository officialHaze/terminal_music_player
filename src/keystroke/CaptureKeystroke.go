package keystroke

import (
	// "bufio"
	// "strings"

	// "fmt"
	"terminalmusicplayer/src/handler"
	// "os"

	"github.com/mattn/go-tty"
)

// Capture keystroke and send it to the control channel
func CaptureKeystroke(controlChan chan rune) {
	tty, err := tty.Open()
	handler.HandleError(err);
	defer tty.Close()

	// fmt.Println("\nControls: [N]ext | [P]revious | [Q]uit")
	for {
		char, err := tty.ReadRune()
		handler.HandleError(err);
		controlChan <- char
		if char == 'q' {
			break
		}
	}
}