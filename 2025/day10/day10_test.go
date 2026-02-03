package day10

import (
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	exampleInput := `
[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}
`
	lines := strings.Split(strings.TrimSpace(exampleInput), "\n")
	expected := 7
	actualStr := SolvePart1(lines) // SolvePart1 returns string
	actualInt, err := strconv.Atoi(actualStr)
	if err != nil {
		t.Fatalf("Failed to convert Part 1 result to int: %v", err)
	}

	if actualInt != expected {
		t.Errorf("Part 1: Expected %d, got %s (converted: %d)", expected, actualStr, actualInt)
	}
}

// TestDay10InputFilePart1 reads input from "input.txt" and solves Part 1
func TestDay10InputFilePart1(t *testing.T) {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		t.Skip("input.txt not found, skipping TestDay10InputFilePart1")
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	result := SolvePart1(lines)

	t.Logf("Part 1 Result from input.txt = %s", result)
}

// TestPart2 reads the example input and solves Part 2
func TestPart2(t *testing.T) {
	exampleInput := `
[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}
`
	lines := strings.Split(strings.TrimSpace(exampleInput), "\n")
	expected := 33
	actualStr := SolvePart2(lines) // SolvePart2 returns string
	actualInt, err := strconv.Atoi(actualStr)
	if err != nil {
		t.Fatalf("Failed to convert Part 2 result to int: %v", err)
	}

	if actualInt != expected {
		t.Errorf("Part 2: Expected %d, got %s (converted: %d)", expected, actualStr, actualInt)
	}
}

// TestDay10InputFilePart2 reads input from "input.txt" and solves Part 2
func TestDay10InputFilePart2(t *testing.T) {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		t.Skip("input.txt not found, skipping TestDay10InputFilePart2")
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	result := SolvePart2(lines)

	t.Logf("Part 2 Result from input.txt = %s", result)
}

// TestParseInput has been removed as parsing is integrated into SolvePart1/SolvePart2
