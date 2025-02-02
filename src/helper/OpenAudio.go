package helper

import (
	"os"
	"github.com/officialHaze/terminal_music_player/src/handler"
)

func OpenAudio(audioPath *string, doneChan chan bool) *os.File {
	// fmt.Println(*audioPath);
	file, err := os.Open(*audioPath)
	handler.HandleError(err, doneChan);

	return file;
}