package day7

import (
	"os"
	"strings"
	"testing"
)

func TestCountSplits(t *testing.T) {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		t.Fatalf("failed to read input.txt: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	result := CountSplits(lines)

	t.Logf("Total splits: %d", result)
}

func TestCountTimelines(t *testing.T) {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		t.Fatalf("failed to read input.txt: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	result := CountTimelines(lines)

	t.Logf("Total timelines: %d", result)
}