package tracktable

import (
	// "fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/officialHaze/terminal_music_player/src/helper"
)

func RenderTrackTable(currentSongIdx *int, metadataList *[]SongMetadata) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	colWidth := GetColWidth();

	// Customization
	CustomizeTrackTable(&t, &colWidth)

	t.AppendHeader(table.Row{"", "Artist", "Title", "Album"})

	// Reset
	t.ResetRows()
	// Populate
	for i, song := range *metadataList {
		title, artist, album := song.Title, song.Artist, song.Album

		if i == *currentSongIdx {
			// Highlight the entire row
			t.AppendRow(table.Row{
				Highlight(">", &colWidth, text.FgHiWhite, text.BgHiRed, false),
				Highlight(artist, &colWidth, text.FgHiWhite, text.BgHiCyan, false),
				Highlight(title, &colWidth, text.FgHiWhite, text.BgHiGreen, false),
				Highlight(album, &colWidth, text.FgHiBlue, text.BgHiYellow, false),
			})
		} else {
			t.AppendRow(table.Row{
				"",
				Highlight(artist, &colWidth, text.FgHiCyan, 0, true),
				Highlight(title, &colWidth, text.FgHiGreen, 0, true),
				Highlight(album, &colWidth, text.FgHiYellow, 0, true),
			})
		}

	}

	t.Render()
}

func padString(s string, width *int) string {
	return helper.Truncate(s, width) + strings.Repeat(" ", *width-len(s)) // Pad spaces
}
