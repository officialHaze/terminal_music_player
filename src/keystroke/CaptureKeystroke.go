package keystroke

import (
	"github.com/mattn/go-tty"
	"github.com/officialHaze/terminal_music_player/src/handler"
)

// Capture keystroke and send it to the control channel
func CaptureKeystroke(controlChan chan rune) {
	tty, err := tty.Open()
	handler.HandleError(err, nil);
	defer tty.Close()

	// fmt.Println("\nControls: [N]ext | [P]revious | [Q]uit")
	for {
		char, err := tty.ReadRune()
		handler.HandleError(err, nil);
		controlChan <- char
		if char == 'q' {
			break
		}
	}
}