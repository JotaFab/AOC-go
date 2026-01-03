package day7

func CountSplits(grid []string) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])

	// find S
	startCol := -1
	for c := 0; c < cols; c++ {
		if grid[0][c] == 'S' {
			startCol = c
			break
		}
	}
	if startCol == -1 {
		return 0
	}

	active := map[int]struct{}{startCol: {}}
	splits := 0

	for r := 1; r < rows && len(active) > 0; r++ {
		next := make(map[int]struct{})

		for c := range active {
			if c < 0 || c >= cols {
				continue
			}
			if grid[r][c] == '^' {
				splits++
				next[c-1] = struct{}{}
				next[c+1] = struct{}{}
			} else {
				next[c] = struct{}{}
			}
		}
		active = next
	}

	return splits
}


func CountTimelines(grid []string) uint64 {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])

	// find S
	startCol := -1
	for c := 0; c < cols; c++ {
		if grid[0][c] == 'S' {
			startCol = c
			break
		}
	}
	if startCol == -1 {
		return 0
	}

	// column -> number of timelines
	active := map[int]uint64{
		startCol: 1,
	}

	for r := 1; r < rows && len(active) > 0; r++ {
		next := make(map[int]uint64)

		for c, count := range active {
			if c < 0 || c >= cols {
				continue
			}

			if grid[r][c] == '^' {
				// split timelines
				next[c-1] += count
				next[c+1] += count
			} else {
				// continue straight
				next[c] += count
			}
		}

		active = next
	}

	// sum all remaining timelines
	var total uint64
	for _, count := range active {
		total += count
	}

	return total
}
