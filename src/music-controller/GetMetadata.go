package musiccontroller

import (
	"github.com/dhowden/tag"
	"github.com/officialHaze/terminal_music_player/src/handler"
	"github.com/officialHaze/terminal_music_player/src/helper"
	tracktable "github.com/officialHaze/terminal_music_player/src/track-table"
)

func GetMetadataList(songPaths *[]string, doneChan chan bool) []tracktable.SongMetadata {

	var metadatalist []tracktable.SongMetadata

	for _, songPath := range *songPaths {
		// Open the audio file
		file := helper.OpenAudio(&songPath, doneChan)

		// Read the tag from file
		m, err := tag.ReadFrom(file)
		handler.HandleError(err, nil)
		metadatalist = append(metadatalist, tracktable.SongMetadata{Title: m.Title(), Artist: m.Artist(), Album: m.Album()})
	}

	return metadatalist
}
