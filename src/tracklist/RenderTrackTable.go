package tracklist

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
)

func RenderTrackTable(currentSongIdx *int, metadataList *[]SongMetadata) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	// Customization
	CustomizeTrackTable(&t)

	// songs := []Song{
	// 	{Title: "Cradled in love", Artist: "Poets of the fall", Album: "unknown"},
	// 	{Title: "F!en", Artist: "Travis Scott", Album: "unknown"},
	// }

	t.AppendHeader(table.Row{"Title", "Artist", "Album"})

	// Reset
	t.ResetRows()
	// Populate
	for i, song := range *metadataList {
		// Clear terminal screen (overwrites previous render)
		fmt.Print("\033[H\033[2J") // Works on Linux/macOS; for Windows, use `cls` command

		title, artist, album := song.Title, song.Artist, song.Album

		if i == *currentSongIdx {
			title = text.Colors{text.FgHiGreen, text.Bold}.Sprint(song.Title)
		}
		t.AppendRow(table.Row{
			title,
			artist,
			album,
		})

	}

	t.Render()
}
