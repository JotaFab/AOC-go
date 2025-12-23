package day5

import (
	"os"
	"testing"
)

const exampleInput = `
3-5
10-14
16-20
12-18

1
5
8
11
17
32
`

func TestExample(t *testing.T) {
	got := CountFreshIngredients(exampleInput)
	want := 3
	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestRealInput(t *testing.T) {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		t.Skip("input.txt not found")
	}

	result := CountFreshIngredients(string(data))
	t.Logf("Day 5 — Part 1 answer: %d", result)

	if result <= 0 {
		t.Fatalf("unexpected result: %d", result)
	}
}

func TestExamplePart2(t *testing.T) {
	input := `
3-5
10-14
16-20
12-18
`

	got := CountFreshIDsFromRanges(input)
	want := 14

	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestRealInputPart2(t *testing.T) {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		t.Skip("input.txt not found")
	}

	result := CountFreshIDsFromRanges(string(data))
	t.Logf("Day 5 — Part 2 answer: %d", result)

	if result <= 0 {
		t.Fatalf("invalid result: %d", result)
	}
}
