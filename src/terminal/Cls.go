package terminal

import "fmt"

func Cls() {
	// Clear terminal screen (overwrites previous render)
	fmt.Print("\033[H\033[2J") // Works on Linux/macOS; for Windows, use `cls` command
}