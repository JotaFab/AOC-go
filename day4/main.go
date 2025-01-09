package main

import (
	"bufio"
	"log"
	"os"
	//"strings"
	"fmt"
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

		log.Println(lis)
		soup = append(soup, lis)
	}
	word := "XMAS"

	// Solve the letter soup
	foundWords := countWordOccurrences(soup, word)

	// Print the results
	log.Println("Found Word XMAS: ", foundWords)

}

var directions = [][2]int{
	{0, 1},   // Right
	{0, -1},  // Left
	{1, 0},   // Down
	{-1, 0},  // Up
	{1, 1},   // Diagonal down-right
	{1, -1},  // Diagonal down-left
	{-1, 1},  // Diagonal up-right
	{-1, -1}, // Diagonal up-left
}

// Verifica si la palabra existe desde una posición específica
func searchFromPosition(grid [][]rune, word string, x, y, dx, dy int) bool {
	rows := len(grid)
	cols := len(grid[0])

	for i := 0; i < 3; i++ {
		nx := x + i*dx
		ny := y + i*dy
		// Verifica si está fuera de los límites o si no coincide
		if nx < 0 || nx >= rows || ny < 0 || ny >= cols || grid[nx][ny] != rune(word[i]) {
			return false
		}
	}
	return true
}

func countWordOccurrences(grid [][]rune, word string) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0
	visited := map[string]bool{}

	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			// Recorre todas las direcciones desde la posición actual
			for _, dir := range directions {
				dx, dy := dir[0], dir[1]

				// Verifica si la palabra puede encontrarse desde aquí en la dirección actual
				if searchFromPosition(grid, word, x, y, dx, dy) {
					// Marca la palabra como encontrada con una clave única basada en posición y dirección
					key := fmt.Sprintf("%d,%d->%d,%d", x, y, x+(len(word)-1)*dx, y+(len(word)-1)*dy)
					if !visited[key] {
						visited[key] = true
						count++
						fmt.Printf("Encontrada '%s' desde (%d, %d) hasta (%d, %d)\n", word, x, y, x+(len(word)-1)*dx, y+(len(word)-1)*dy)
					}
				}
			}
		}
	}

	return count
}
