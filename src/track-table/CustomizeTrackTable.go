package tracktable

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	// "github.com/officialHaze/terminal_music_player/src/terminal"
)

func CustomizeTrackTable(tableWriter *table.Writer, colWidth *int) {
	t := *tableWriter

	// t.SetStyle(table.StyleLight)
	// t.Style().Options.SeparateRows = true
	t.SetStyle(table.Style{
		Name: "Custom",
		Box: table.BoxStyle{
			BottomLeft: "╰",
			BottomRight: "╯",
			BottomSeparator: "─",
			Left: "│",
			Right: "│",
			MiddleHorizontal: "─",
			MiddleSeparator: "─",
			TopLeft: "╭",
			TopRight: "╮",
			TopSeparator: "─",
		},
		Color: table.ColorOptions{
			Header: text.Colors{text.FgHiMagenta, text.Bold},
			Row:    text.Colors{text.FgHiBlue},
		},
		Format: table.FormatOptions{
			Header: text.FormatDefault,
			Row:    text.FormatDefault,
			RowAlign: text.AlignCenter,
		},
		Options: table.Options{
			DrawBorder:      false,
			SeparateColumns: true,
			SeparateHeader:  true,
			SeparateRows:    false,
		},
	})
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, WidthMax: 2, WidthMin: 2, AlignHeader: text.AlignCenter, Align: text.AlignCenter}, // Left arrow, denoting the song being played
		{Number: 2, WidthMax: *colWidth, WidthMin: *colWidth, AlignHeader: text.AlignCenter, Align: text.AlignCenter},
		{Number: 3, WidthMax: *colWidth, WidthMin: *colWidth, AlignHeader: text.AlignCenter, Align: text.AlignCenter},
		{Number: 4, WidthMax: *colWidth, WidthMin: *colWidth, AlignHeader: text.AlignCenter, Align: text.AlignCenter},
	})
}
