package musiccontroller

import (
	"time"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/speaker"
	"github.com/officialHaze/terminal_music_player/src/tracklist"
)

func Playsong(format *beep.Format, songIdx *int, strm *beep.StreamSeekCloser, metadataList *[]tracklist.SongMetadata) {
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	// re-render the table (showing the current song that is being played);
	tracklist.RenderTrackTable(songIdx, metadataList);

	speaker.PlayAndWait(*strm)
}