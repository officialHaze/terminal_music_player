package musiccontroller

import (
	"time"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/speaker"
	prettytable "github.com/officialHaze/terminal_music_player/src/pretty-table"
)

func Playsong(format *beep.Format, songIdx *int, strm *beep.StreamSeekCloser) {
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	// re-render the table (showing the current song that is being played);
	prettytable.RenderTrackTable(songIdx)

	speaker.PlayAndWait(*strm)
}