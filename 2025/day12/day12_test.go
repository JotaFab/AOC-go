package day12

import (
	"os"
	"strings"
	"testing"
)

// The example shapes for parsing tests (without a trailing newline)
const exampleShapesInput = `0:
###
##.
##.

1:
###
##.
.##

2:
.##
###
##.

3:
##.
###
##.

4:
###
#..
###

5:
###
.#.
###` // No trailing newline here

func TestPart1(t *testing.T) {
	exampleInput := exampleShapesInput + "\n\n" + `
4x4: 0 0 0 0 2 0
12x5: 1 0 1 0 2 2
12x5: 1 0 1 0 3 2
`
	expected := 2
	actual := SolvePart1(exampleInput)

	if actual != expected {
		t.Errorf("Part 1: Expected %d, got %d", expected, actual)
	}
}

// TestDay12InputFilePart1 reads input from "input.txt" and solves Part 1
func TestDay12InputFilePart1(t *testing.T) {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		t.Skip("input.txt not found, skipping TestDay12InputFilePart1")
	}

	input := strings.TrimSpace(string(data))
	result := SolvePart1(input)

	t.Logf("Part 1 Result from input.txt = %d", result)
}

func TestParseInput(t *testing.T) {
	input := exampleShapesInput 
	
	allOrientedShapes, regions, err := parseInput(input)
	if err != nil {
		t.Fatalf("parseInput failed: %v", err)
	}

	if len(allOrientedShapes) != 6 {
		t.Errorf("Expected 6 unique shapes, got %d", len(allOrientedShapes))
	}
	if len(regions) != 1 {
		t.Errorf("Expected 1 region, got %d", len(regions))
	}

	// Verify a specific shape's properties after parsing and normalization
	if shape4Orientations, ok := allOrientedShapes[4]; ok {
		if len(shape4Orientations) == 0 {
			t.Errorf("Shape 4 has no orientations generated")
		} else {
			firstOrientation := shape4Orientations[0]
			if firstOrientation.Width != 3 || firstOrientation.Height != 3 {
				t.Errorf("Shape 4 first orientation dimensions: Expected 3x3, got %dx%d", firstOrientation.Width, firstOrientation.Height)
			}
			expectedPointsCount := 7
			if len(firstOrientation.Points) != expectedPointsCount {
				t.Errorf("Shape 4 first orientation point count: Expected %d, got %d", expectedPointsCount, len(firstOrientation.Points))
			}
		}
	} else {
		t.Errorf("Shape with index 4 not found")
	}

	// Verify region properties
	region := regions[0]
	if region.Width != 4 || region.Height != 4 {
		t.Errorf("Region dimensions: Expected 4x4, got %dx%d", region.Width, region.Height)
	}
	if q, ok := region.PresentsToFit[4]; !ok || q != 2 {
		t.Errorf("Region presentsToFit[4]: Expected 2, got %d", q)
	}
}

func TestShapeTransformations(t *testing.T) {
	// Simple square shape:
	// ##
	// ##
	squareGrid := [][]bool{
		{true, true},
		{true, true},
	}
	squareShape := normalizeShape(squareGrid)

	// Test rotation
	rotatedOnce := rotateShape(squareShape)
	if rotatedOnce.Width != 2 || rotatedOnce.Height != 2 {
		t.Errorf("Rotated square dimensions: Expected 2x2, got %dx%d", rotatedOnce.Width, rotatedOnce.Height)
	}
	// Correctly convert Shape to OrientedShape for canonicalString
	if canonicalString(OrientedShape{Points: squareShape.Points, Width: squareShape.Width, Height: squareShape.Height}) !=
	   canonicalString(OrientedShape{Points: rotatedOnce.Points, Width: rotatedOnce.Width, Height: rotatedOnce.Height}) {
		t.Errorf("Rotated square changed canonical form unexpectedly")
	}

	// Test horizontal flip
	flippedH := flipShapeHorizontal(squareShape)
	if flippedH.Width != 2 || flippedH.Height != 2 {
		t.Errorf("Flipped square dimensions: Expected 2x2, got %dx%d", flippedH.Width, flippedH.Height)
	}
	// Correctly convert Shape to OrientedShape for canonicalString
	if canonicalString(OrientedShape{Points: squareShape.Points, Width: squareShape.Width, Height: squareShape.Height}) !=
	   canonicalString(OrientedShape{Points: flippedH.Points, Width: flippedH.Width, Height: flippedH.Height}) {
		t.Errorf("Flipped square changed canonical form unexpectedly")
	}

	// L-shape:
	// #.
	// ##
	// .#
	lShapeGrid := [][]bool{
		{true, false},
		{true, true},
		{false, true},
	}
	lShape := normalizeShape(lShapeGrid)

	// Test rotation of L-shape
	rotatedLShape := rotateShape(lShape)
	if canonicalString(OrientedShape{Points: lShape.Points, Width: lShape.Width, Height: lShape.Height}) ==
	   canonicalString(OrientedShape{Points: rotatedLShape.Points, Width: rotatedLShape.Width, Height: rotatedLShape.Height}) {
		t.Errorf("Rotated L-shape did not change canonical form")
	}

	// Test generateAllOrientations logic
	allOrientationsLShape := generateAllOrientations(lShape)
	if len(allOrientationsLShape) < 1 { 
		t.Errorf("Expected at least one orientation for L-shape, got %d", len(allOrientationsLShape))
	}
}
