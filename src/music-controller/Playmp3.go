package musiccontroller

import (
	"fmt"
	"os"
	"time"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/mp3"
	"github.com/officialHaze/terminal_music_player/src/handler"
	"github.com/officialHaze/terminal_music_player/src/helper"
)

func decodeMp3(file *os.File, doneChan chan bool) (beep.StreamSeekCloser, *beep.Format) {
	fmt.Println("Decoding file...");
	strm, format, err := mp3.Decode(file);
	handler.HandleError(err, doneChan);

	fmt.Println("File decoded!");
	return strm, &format;
}

func Playmp3(tracklist []string, count int, controlChan chan rune, doneChan chan bool) {
	file := helper.OpenAudio(&tracklist[count], doneChan); // Open/Read the audio file

	metadataList := GetMetadataList(&tracklist, doneChan);

	strm, format := decodeMp3(file, doneChan);

	defer file.Close();
	defer strm.Close();

	go Playsong(format, &count, &strm, &metadataList); // Play the song in a separate go routine in non blocking format

	trackListSize := len(tracklist);
	
	for {
		select {
		case command := <-controlChan:
			switch command {
			case 'q': // Quit
				fmt.Println("\nQuitting audio playback.")
				strm.Close() // close the current stream
				doneChan <- true
				return

			case 'n': // next
				strm.Close() // close the current stream
				Playmp3(tracklist, helper.GetNextSongIdx(&count, &trackListSize), controlChan, doneChan)

			case 'p': // previous
				strm.Close() // close the current stream
				Playmp3(tracklist, helper.GetPrevSongIdx(&count, &trackListSize), controlChan, doneChan)
			}

		default:
			time.Sleep(10 * time.Millisecond) // Prevents CPU overuse
		}
	}
}
