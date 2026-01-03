package day6

import (
	"os"
	"testing"
)

func TestExample(t *testing.T) {
	input := `
123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +
`

	want := 4277556
	got := Solve(input)

	if got != want {
		t.Fatalf("expected %d, got %d", want, got)
	}
}


func TestRealInput(t *testing.T) {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		t.Fatalf("failed to read input.txt: %v", err)
	}

	result := Solve(string(data))
	t.Logf("Day 6 result: %d", result)
}

func TestPartTwoExample(t *testing.T) {
	input := `
123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +
`

	want := 3263827
	got := SolvePartTwo(input)

	if got != want {
		t.Fatalf("expected %d, got %d", want, got)
	}
}

func TestPartTwoRealInput(t *testing.T) {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		t.Fatalf("failed to read input.txt: %v", err)
	}

	result := SolvePartTwo(string(data))
	t.Logf("Day 6 Part Two result: %d", result)
}