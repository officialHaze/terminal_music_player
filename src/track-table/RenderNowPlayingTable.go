package tracktable

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/officialHaze/terminal_music_player/src/helper"
)

func RenderNowPlayingTable(songTitle *string, artist *string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	colWidth := GetColWidth();
	
	// Customization
	CustomizeTrackTable(&t, &colWidth)

	t.AppendHeader(table.Row{
		"",
		ApplyColor("(N)ext / (P)rev / (Q)uit", text.FgWhite, 0, true),
		ApplyColor(helper.Truncate("[   "+*songTitle+"   ]", &colWidth), text.FgHiCyan, 0, true),
		ApplyColor("Volume 100%", text.FgWhite, 0, true),
	})

	t.AppendHeader(table.Row{
		"",
		ApplyColor("[Playing]", text.FgWhite, 0, true),
		ApplyColor(*artist, text.FgHiBlue, 0, true),
		ApplyColor("[-----]", text.FgWhite, 0, true),
	})

	t.Render()
}
