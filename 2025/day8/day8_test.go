package day8

import (
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestSolveDay8(t *testing.T) {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		t.Fatalf("failed to read input.txt: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	points := make([]Point, 0, len(lines))

	for _, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			t.Fatalf("invalid line: %s", line)
		}

		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])

		points = append(points, Point{x, y, z})
	}

	result := SolveDay8(points)
	t.Logf("Day 8 result: %d", result)
}
