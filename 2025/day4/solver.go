package day4

type Pos struct {
	r, c int
}

var dirs = [8][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1},           {0, 1},
	{1, -1},  {1, 0},  {1, 1},
}

// Part 1
func CountAccessible(grid [][]rune) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] != '@' {
				continue
			}
			if adjacentCount(grid, r, c) < 4 {
				count++
			}
		}
	}
	return count
}

// Part 2
func RemoveAllAccessible(grid [][]rune) int {
	rows := len(grid)
	cols := len(grid[0])
	total := 0

	for {
		var toRemove []Pos

		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				if grid[r][c] != '@' {
					continue
				}
				if adjacentCount(grid, r, c) < 4 {
					toRemove = append(toRemove, Pos{r, c})
				}
			}
		}

		if len(toRemove) == 0 {
			break
		}

		for _, p := range toRemove {
			grid[p.r][p.c] = '.'
		}

		total += len(toRemove)
	}

	return total
}

func adjacentCount(grid [][]rune, r, c int) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	for _, d := range dirs {
		nr, nc := r+d[0], c+d[1]
		if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
			if grid[nr][nc] == '@' {
				count++
			}
		}
	}
	return count
}
