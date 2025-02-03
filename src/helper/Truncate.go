package helper

func Truncate(s string, width *int) string {
	if len(s) >= *width {
		return s[:*width] // Trim if too long
	}

	return s
}
