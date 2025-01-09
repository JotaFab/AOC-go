package main

import (
	"bufio"
	"log"
	"os"
	//"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	var soup [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var lis []rune
		for _, l := range line {
			lis = append(lis, l)
		}

		soup = append(soup, lis)
	}
	word := "MAS"

	// Solve the letter soup
	foundWords := countXMASOccurrences(soup, word)

	// Print the results
	log.Println("Numeber of X-MAS found: ", foundWords)

}

// Verifica si la palabra existe desde una posición específica
func isValidDiagonal(grid [][]rune, x, y, dx1, dy1, dx2, dy2 int) bool {
	rows := len(grid)
	cols := len(grid[0])

	if x+dx1 >= 0 && x+dx1 < rows && y+dy1 >= 0 && y+dy1 < cols && x+dx2 >= 0 && x+dx2 < rows && y+dy2 >= 0 && y+dy2 < cols {
		if grid[x+dx1][y+dy1] == rune('M') && grid[x+dx2][y+dy2] == rune('S') {
			return true
		}
		if grid[x+dx1][y+dy1] == rune('S') && grid[x+dx2][y+dy2] == rune('M') {
			return true
		}
	}
	return false
}

func countXMASOccurrences(grid [][]rune, word string) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			if grid[x][y] == rune(word[1]) {
				if isValidDiagonal(grid, x, y, -1, -1, 1, 1) && isValidDiagonal(grid, x, y, 1, -1, -1, 1) {
					count++
				}
			}
		}
	}

	return count
}
