package musiccontroller

import (
	"time"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/speaker"
	"github.com/officialHaze/terminal_music_player/src/terminal"
	tracktable "github.com/officialHaze/terminal_music_player/src/track-table"
)

func Playsong(format *beep.Format, songIdx *int, strm *beep.StreamSeekCloser, metadataList *[]tracktable.SongMetadata) {
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	metalist := *metadataList;
	currentSongTitle := metalist[*songIdx].Title;
	currentSongArtist := metalist[*songIdx].Artist;

	// Clear the screen before each render
	terminal.Cls();
	// re-render the now playing table
	tracktable.RenderNowPlayingTable(&currentSongTitle, &currentSongArtist);
	// re-render the table (highlighting the current song that is being played);
	tracktable.RenderTrackTable(songIdx, metadataList);

	speaker.PlayAndWait(*strm)
}