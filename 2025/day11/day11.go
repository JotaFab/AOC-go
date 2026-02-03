package day11

import (
	"strings"
)

// Graph represents the device connections as an adjacency list.
type Graph map[string][]string

// parseInput parses the raw input string into a Graph.
func parseInput(input string) Graph {
	graph := make(Graph)
	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue // Skip malformed lines
		}
		source := strings.TrimSpace(parts[0])
		destinationsStr := strings.TrimSpace(parts[1])
		
		var destinations []string
		if destinationsStr != "" { // Handle cases where a device has no outputs
			destinations = strings.Fields(destinationsStr)
		}
		graph[source] = destinations
	}
	return graph
}

// countPathsDFS recursively counts all paths from start to end (for Part 1).
// This uses memoization (dynamic programming) to avoid recomputing paths
// for the same subproblem (node), which is crucial for efficiency in DAGs.
func countPathsDFS(graph Graph, start, end string, memo map[string]int) int {
	if start == end {
		return 1
	}
	if val, found := memo[start]; found {
		return val
	}

	count := 0
	if destinations, exists := graph[start]; exists {
		for _, nextNode := range destinations {
			count += countPathsDFS(graph, nextNode, end, memo)
		}
	}
	memo[start] = count
	return count
}

// SolvePart1 solves Day 11 Part 1.
func SolvePart1(input string) int {
	graph := parseInput(input)
	memo := make(map[string]int) // Memoization map for DFS
	return countPathsDFS(graph, "you", "out", memo)
}

// --- Part 2 Logic ---

// Special node flags for memoization key
const (
	FLAG_DAC = 1 << 0
	FLAG_FFT = 1 << 1
	TARGET_FLAGS = FLAG_DAC | FLAG_FFT // Both DAC and FFT must be visited
)

// countPathsPart2 recursively counts paths from start to end, visiting required special nodes.
// The memoization key combines the current node and the visited special nodes flags.
func countPathsPart2(graph Graph, current, end string, visitedFlags uint, memo map[string]map[uint]int) int {
	// Update visited flags if current node is a special node
	if current == "dac" {
		visitedFlags |= FLAG_DAC
	} else if current == "fft" {
		visitedFlags |= FLAG_FFT
	}

	// Base case: If we reached the end node
	if current == end {
		if visitedFlags == TARGET_FLAGS {
			return 1 // Path found, and both special nodes visited
		}
		return 0 // Reached end, but not both special nodes visited
	}

	// Check memoization
	if nodeMemo, found := memo[current]; found {
		if val, foundFlags := nodeMemo[visitedFlags]; foundFlags {
			return val
		}
	} else {
		memo[current] = make(map[uint]int)
	}

	count := 0
	if destinations, exists := graph[current]; exists {
		for _, nextNode := range destinations {
			count += countPathsPart2(graph, nextNode, end, visitedFlags, memo)
		}
	}

	memo[current][visitedFlags] = count
	return count
}


// SolvePart2 solves Day 11 Part 2.
func SolvePart2(input string) int {
	graph := parseInput(input)
	// Memoization map for DFS. Key is current node, value is another map
	// where key is visitedFlags and value is path count.
	memo := make(map[string]map[uint]int) 
	return countPathsPart2(graph, "svr", "out", 0, memo)
}
