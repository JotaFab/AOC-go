package day4

import (
	"os"
	"strings"
	"testing"
)

func gridFromStrings(lines []string) [][]rune {
	grid := make([][]rune, len(lines))
	for i := range lines {
		grid[i] = []rune(lines[i])
	}
	return grid
}

func cloneGrid(grid [][]rune) [][]rune {
	cp := make([][]rune, len(grid))
	for i := range grid {
		cp[i] = make([]rune, len(grid[i]))
		copy(cp[i], grid[i])
	}
	return cp
}

var example = []string{
	"..@@.@@@@.",
	"@@@.@.@.@@",
	"@@@@@.@.@@",
	"@.@@@@..@.",
	"@@.@@@@.@@",
	".@@@@@@@.@",
	".@.@.@.@@@",
	"@.@@@.@@@@",
	".@@@@@@@@.",
	"@.@.@@@.@.",
}

func TestExamplePart1(t *testing.T) {
	grid := gridFromStrings(example)
	if got := CountAccessible(grid); got != 13 {
		t.Fatalf("Part1 = %d, want 13", got)
	}
}

func TestExamplePart2(t *testing.T) {
	grid := gridFromStrings(example)
	if got := RemoveAllAccessible(grid); got != 43 {
		t.Fatalf("Part2 = %d, want 43", got)
	}
}

func TestRealInput(t *testing.T) {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		t.Skip("input.txt not found")
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	grid := gridFromStrings(lines)

	part1 := CountAccessible(grid)
	t.Logf("Day 4 — Part 1 answer: %d", part1)

	if part1 != 1351 {
		t.Fatalf("Part1 = %d, want 1351", part1)
	}

	part2 := RemoveAllAccessible(cloneGrid(grid))
	t.Logf("Day 4 — Part 2 answer: %d", part2)

	if part2 <= part1 {
		t.Fatalf("Part2 (%d) must be > Part1 (%d)", part2, part1)
	}
}
