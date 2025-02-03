package tracktable

import "github.com/jedib0t/go-pretty/v6/text"

func Highlight(s string, width *int, fg text.Color, bg text.Color, noBg bool) string {
	if noBg {
		return text.Colors{text.Bold, fg}.Sprintf("%s", padString(s, width)) // Left-align & fill with spaces
	}
	return text.Colors{text.Bold, fg, bg}.Sprintf("%s", padString(s, width)) // Left-align & fill with spaces
}
