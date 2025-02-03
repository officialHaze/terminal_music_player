package tracktable

import (
	"github.com/jedib0t/go-pretty/v6/text"
)

func ApplyColor(s string, fg text.Color, bg text.Color, noBg bool) string {
	if noBg {
		return text.Colors{text.Bold, fg}.Sprint(s)
	}
	return text.Colors{text.FgCyan, bg}.Sprint(s);
}