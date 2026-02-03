package day11

import (
	"os"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	exampleInput := `
aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out
`
	expected := 5
	actual := SolvePart1(exampleInput)

	if actual != expected {
		t.Errorf("Part 1: Expected %d, got %d", expected, actual)
	}
}

// TestDay11InputFilePart1 reads input from "input.txt" and solves Part 1
func TestDay11InputFilePart1(t *testing.T) {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		t.Skip("input.txt not found, skipping TestDay11InputFilePart1")
	}

	input := strings.TrimSpace(string(data))
	result := SolvePart1(input)

	t.Logf("Part 1 Result from input.txt = %d", result)
}

// TestPart2 reads the example input and solves Part 2
func TestPart2(t *testing.T) {
	exampleInput := `
svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out
`
	expected := 2
	actual := SolvePart2(exampleInput)

	if actual != expected {
		t.Errorf("Part 2: Expected %d, got %d", expected, actual)
	}
}

// TestDay11InputFilePart2 reads input from "input.txt" and solves Part 2
func TestDay11InputFilePart2(t *testing.T) {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		t.Skip("input.txt not found, skipping TestDay11InputFilePart2")
	}

	input := strings.TrimSpace(string(data))
	result := SolvePart2(input)

	t.Logf("Part 2 Result from input.txt = %d", result)
}


// TestParseInput to verify graph parsing
func TestParseInput(t *testing.T) {
	input := `
you: a b
a: out
b: out
`
	graph := parseInput(input)

	if len(graph) != 3 {
		t.Errorf("Expected 3 nodes, got %d", len(graph))
	}

	if dests, exists := graph["you"]; !exists || len(dests) != 2 || dests[0] != "a" || dests[1] != "b" {
		t.Errorf("Incorrect destinations for 'you': %v", dests)
	}
	if dests, exists := graph["a"]; !exists || len(dests) != 1 || dests[0] != "out" {
		t.Errorf("Incorrect destinations for 'a': %v", dests)
	}
	if dests, exists := graph["b"]; !exists || len(dests) != 1 || dests[0] != "out" {
		t.Errorf("Incorrect destinations for 'b': %v", dests)
	}
	if _, exists := graph["out"]; !exists { 
		t.Logf("'out' node not explicitly in graph, which is fine if it only acts as a sink.")
	}
}
