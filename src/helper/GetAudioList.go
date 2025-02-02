package helper

import (
	"fmt"
	"os"
	"terminalmusicplayer/src/handler"
)

func GetAudioList(audioDir *string) []string {

	var audioPaths []string;

	entries, err := os.ReadDir(*audioDir);
	handler.HandleError(err);

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