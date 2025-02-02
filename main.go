package main

import (
	"fmt"

	"github.com/officialHaze/terminal_music_player/src/helper"
	"github.com/officialHaze/terminal_music_player/src/keystroke"
	musiccontroller "github.com/officialHaze/terminal_music_player/src/music-controller"
)

func main() {
	fmt.Println("Staring program...")

	controlChan := make(chan rune)
	doneChan := make(chan bool)

	musicDir := "C:/Users/moina/Downloads/Music/mp3"; // assigned music directory
	trackList := helper.GetTrackList(&musicDir); // Get the list of tracks from the music directory

	go keystroke.CaptureKeystroke(controlChan); // Run it as a sepearate go routine for non blocking IO
	go musiccontroller.Playmp3(trackList, 0, controlChan, doneChan) // Run it as a separate go routine to not block keystroke capturing

	<-doneChan
	fmt.Println("Stopping program!")
}
