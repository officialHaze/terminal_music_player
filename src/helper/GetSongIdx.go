package helper


func GetNextSongIdx(initialIdx *int, trackListSize *int) int {
	idx := *initialIdx + 1;

	if idx >= *trackListSize {
		idx = 0;
	}

	return idx;
}


func GetPrevSongIdx(initialIdx *int, trackListSize *int) int {
	idx := *initialIdx - 1;

	if idx < 0 {
		idx = *trackListSize - 1;
	}

	return idx;
}