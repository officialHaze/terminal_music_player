package terminal

import "os"

func GetMusicDirArg() string {
	args := os.Args;
	musicDir := args[len(args) - 1];

	return musicDir;
}