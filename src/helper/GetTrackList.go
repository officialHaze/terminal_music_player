package helper

import (
	"fmt"
	"os"

	"github.com/officialHaze/terminal_music_player/src/handler"
)

func GetTrackList(audioDir *string) []string {

	var audioPaths []string;

	entries, err := os.ReadDir(*audioDir);
	handler.HandleError(err, nil);

	// Print the name of all the audios
	for _, audio := range entries {
		if audio.IsDir() {
			continue;
		}
		fmt.Println(audio.Name());
		audioPaths = append(audioPaths, *audioDir+"/"+audio.Name());
	}

	return audioPaths;
}