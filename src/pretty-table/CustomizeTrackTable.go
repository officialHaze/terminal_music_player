package prettytable

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/officialHaze/terminal_music_player/src/terminal"
)

func CustomizeTrackTable(tableWriter *table.Writer) {
	t := *tableWriter;

	t.SetStyle(table.StyleDefault)
	t.Style().Options.SeparateRows = true

	// Get terminal width
	terminalWidth := terminal.GetTerminalWidth();

	// Set column widths dynamically
	colWidth := terminalWidth / 3
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, WidthMax: colWidth, WidthMin: colWidth},
		{Number: 2, WidthMax: colWidth, WidthMin: colWidth},
		{Number: 3, WidthMax: colWidth, WidthMin: colWidth},
	})
}