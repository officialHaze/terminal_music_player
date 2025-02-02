package main

import (
	// "bufio"
	"fmt"
	// "log"
	// "os"

	audiocontroller "terminalmusicplayer/src/audioController"
	"terminalmusicplayer/src/helper"
	"terminalmusicplayer/src/keystroke"
	// "terminalAudioPlayer/src/handler"
	// "terminalAudioPlayer/src/helper"
)

func main() {
	fmt.Println("Staring program...")

	controlChan := make(chan rune)
	doneChan := make(chan bool)

	songDir := "C:/Users/moina/Downloads/Music/mp3/"
	songs := helper.GetAudioList(&songDir)

	go keystroke.CaptureKeystroke(controlChan); // Run it as a sepearate go routine for non blocking IO
	go audiocontroller.Playmp3(songs, 0, controlChan, doneChan) // Run it as a separate go routine to not block keystroke capturing

	<-doneChan
	fmt.Println("Stopping program!")
}
