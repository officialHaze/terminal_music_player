package tracktable

import "github.com/officialHaze/terminal_music_player/src/terminal"

func GetColWidth() int {
	terminalWidth := terminal.GetTerminalWidth();

	return terminalWidth / 4;
}