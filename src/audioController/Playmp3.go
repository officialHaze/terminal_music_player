package audiocontroller

import (
	"fmt"
	"os"
	"strings"
	"terminalmusicplayer/src/handler"
	"time"

	"github.com/gopxl/beep/v2/mp3"
	"github.com/gopxl/beep/v2/speaker"
)

func Playmp3(songlist []string, count int, controlChan chan rune, doneChan chan bool) {
	file, err := os.Open(songlist[count]);
	if err != nil {
		doneChan <- true;
	}
	handler.HandleError(err);

	defer file.Close();

	songpath := songlist[count];

	fmt.Printf("Playing ----> %s\n", getSongname(&songpath));

	strm, format, err := mp3.Decode(file);
	if err != nil {
		doneChan <- true;
	}
	handler.HandleError(err)

	defer strm.Close()

	go func() {
		speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
		speaker.PlayAndWait(strm)
	}() // Play the audio in a separate go routine

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
				Playmp3(songlist, getNewCount(&command, &count, len(songlist)), controlChan, doneChan)

			case 'p': // previous
				strm.Close() // close the current stream
				Playmp3(songlist, getNewCount(&command, &count, len(songlist)), controlChan, doneChan)
			}

		default:
			time.Sleep(10 * time.Millisecond) // Prevents CPU overuse
		}
	}
}

func getSongname(songpath *string) string {
	splits := strings.Split(*songpath, "/");
	return splits[len(splits) - 1];
}

func getNewCount(control *rune, initalCount *int, songListSize int) int {
	newCount := 0
	if *control == 'n' {
		newCount = *initalCount + 1;
	} else if *control == 'p' {
		newCount = *initalCount - 1;
	}

	if newCount < 0 {
		newCount = songListSize - 1; // last song
	} else if newCount >= songListSize {
		newCount = 0; // first song
	}

	return newCount;
}
