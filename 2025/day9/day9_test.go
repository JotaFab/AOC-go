package day9

import (
	"os"
	"strings"
	"testing"
)

const sampleInput = `
7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3
`

func TestDay9SamplePart1(t *testing.T) {
	lines := strings.Split(strings.TrimSpace(sampleInput), "\n")
	result := Part1(lines)
	expected := 50

	if result != expected {
		t.Errorf("Part1 sample = %d, want %d", result, expected)
	}
}

func TestDay9InputFilePart1(t *testing.T) {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		t.Skip("input.txt not found")
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	result := Part1(lines)

	t.Logf("Part1 Result = %d", result)
}

func TestDay9SamplePart2(t *testing.T) {
	lines := strings.Split(strings.TrimSpace(sampleInput), "\n")
	result := Part2(lines)
	var expected int64 = 24

	if result != expected {
		t.Errorf("Part2 sample = %d, want %d", result, expected)
	}
}

func TestDay9InputFilePart2(t *testing.T) {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		t.Skip("input.txt not found")
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	result := Part2(lines)
	t.Logf("Part2 Result = %d", result)

}

