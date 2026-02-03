package day12

import (
	// "bytes" // Removed unused import
	"fmt"
	"strconv"
	"strings"
)

// Point represents a coordinate (row, col)
type Point struct {
	R, C int
}

// Shape represents a present shape as a slice of points relative to its top-leftmost occupied cell.
type Shape struct {
	Points []Point
	Width  int
	Height int
}

// OrientedShape represents a shape in a specific orientation (rotation/flip).
// Points are relative to its top-left corner, and also stores its dimensions.
type OrientedShape struct {
	Points []Point
	Width  int
	Height int
}

// Region represents a tree region and the presents to fit.
type Region struct {
	Width        int
	Height       int
	PresentsToFit map[int]int // map[shapeIndex]quantity
}

// parseInput parses the raw puzzle input into shapes and regions.
func parseInput(input string) (map[int][]OrientedShape, []Region, error) {
	parts := strings.Split(strings.TrimSpace(input), "\n\n")
	if len(parts) != 2 {
		return nil, nil, fmt.Errorf("expected two sections (shapes and regions), got %d", len(parts))
	}

	// --- Parse Shapes ---
	shapeLines := strings.Split(parts[0], "\n")
	rawShapes := make(map[int]Shape)
	currentShapeIdx := -1
	var currentShapeGrid [][]bool
	var maxShapeWidth, maxShapeHeight int

	for _, line := range shapeLines {
		if strings.HasSuffix(line, ":") {
			if currentShapeIdx != -1 { // Save previous shape
				rawShapes[currentShapeIdx] = normalizeShape(currentShapeGrid)
			}
			idxStr := strings.TrimSuffix(line, ":")
			idx, err := strconv.Atoi(idxStr)
			if err != nil {
				return nil, nil, fmt.Errorf("invalid shape index: %s", idxStr)
			}
			currentShapeIdx = idx
			currentShapeGrid = nil // Reset for new shape
			maxShapeWidth, maxShapeHeight = 0, 0
		} else {
			row := make([]bool, len(line))
			for i, r := range line {
				if r == '#' {
					row[i] = true
				}
			}
			currentShapeGrid = append(currentShapeGrid, row)
			if len(row) > maxShapeWidth {
				maxShapeWidth = len(row)
			}
			maxShapeHeight++
		}
	}
	if currentShapeIdx != -1 { // Save the last shape
		rawShapes[currentShapeIdx] = normalizeShape(currentShapeGrid)
	}

	// Generate all unique orientations for each raw shape
	allOrientedShapes := make(map[int][]OrientedShape)
	for idx, shape := range rawShapes {
		allOrientedShapes[idx] = generateAllOrientations(shape)
	}

	// --- Parse Regions ---
	regionLines := strings.Split(parts[1], "\n")
	var regions []Region
	for _, line := range regionLines {
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue // Skip malformed lines
		}
		
		dimParts := strings.SplitN(fields[0], "x", 2)
		if len(dimParts) != 2 {
			return nil, nil, fmt.Errorf("invalid region dimensions: %s", fields[0])
		}
		width, err := strconv.Atoi(dimParts[0])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid region width: %s", dimParts[0])
		}
		height, err := strconv.Atoi(dimParts[1])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid region height: %s", dimParts[1])
		}

		presentsToFit := make(map[int]int)
		// Process present quantities (e.g., "0 0 0 0 2 0" -> {4:2})
		if len(fields) >= 2 {
			for i := 1; i < len(fields); i += 2 {
				if i+1 >= len(fields) {
					return nil, nil, fmt.Errorf("malformed presents list: %s", line)
				}
				shapeIdx, err := strconv.Atoi(fields[i])
				if err != nil {
					return nil, nil, fmt.Errorf("invalid present shape index: %s", fields[i])
				}
				quantity, err := strconv.Atoi(fields[i+1])
				if err != nil {
					return nil, nil, fmt.Errorf("invalid present quantity: %s", fields[i+1])
				}
				if quantity > 0 { // Only add if quantity is positive
					presentsToFit[shapeIdx] = quantity
				}
			}
		}
		regions = append(regions, Region{Width: width, Height: height, PresentsToFit: presentsToFit})
	}

	return allOrientedShapes, regions, nil
}

// normalizeShape converts a raw grid representation to a Shape struct
// and shifts points so the top-leftmost occupied cell is at (0,0).
func normalizeShape(grid [][]bool) Shape {
	var points []Point
	minR, minC := 1_000_000, 1_000_000 // Effectively infinity

	// Handle empty grid (e.g. from parsing error or empty shape)
	if len(grid) == 0 {
		return Shape{}
	}

	// Find occupied points and min R/C
	for r, row := range grid {
		for c, val := range row {
			if val {
				points = append(points, Point{R: r, C: c})
				if r < minR {
					minR = r
				}
				if c < minC {
					minC = c
				}
			}
		}
	}

	// If no points, it's an empty shape
	if len(points) == 0 {
		return Shape{}
	}

	// Shift points
	for i := range points {
		points[i].R -= minR
		points[i].C -= minC
	}

	// Calculate dimensions after normalization
	maxWidth, maxHeight := 0, 0
	for _, p := range points {
		if p.C + 1 > maxWidth {
			maxWidth = p.C + 1
		}
		if p.R + 1 > maxHeight {
			maxHeight = p.R + 1
		}
	}

	return Shape{Points: points, Width: maxWidth, Height: maxHeight}
}

// generateAllOrientations generates all unique rotations and flips of a given shape.
func generateAllOrientations(s Shape) []OrientedShape {
	uniqueOrientations := make(map[string]OrientedShape)

	currentShape := s
	for r := 0; r < 4; r++ { // 4 rotations (0, 90, 180, 270)
		// Add current rotation
		addUniqueOrientation(uniqueOrientations, currentShape)

		// Add horizontal flip of current rotation
		flippedH := flipShapeHorizontal(currentShape)
		addUniqueOrientation(uniqueOrientations, flippedH)
		
		// Rotate for next iteration
		currentShape = rotateShape(currentShape)
	}

	var orientations []OrientedShape
	for _, os := range uniqueOrientations {
		orientations = append(orientations, os)
	}
	return orientations
}

// addUniqueOrientation adds a shape to the unique set if its canonical string representation is new.
func addUniqueOrientation(unique map[string]OrientedShape, s Shape) {
	os := OrientedShape{Points: s.Points, Width: s.Width, Height: s.Height}
	key := canonicalString(os)
	unique[key] = os
}

// rotateShape rotates a shape 90 degrees clockwise.
func rotateShape(s Shape) Shape {
	if len(s.Points) == 0 {
		return Shape{}
	}
	var rotatedPoints []Point
	newHeight := s.Width
	newWidth := s.Height

	for _, p := range s.Points {
		rotatedPoints = append(rotatedPoints, Point{R: p.C, C: newWidth - 1 - p.R})
	}
	rotatedShape := Shape{Points: rotatedPoints, Width: newWidth, Height: newHeight}
	// Re-normalize after rotation to ensure (0,0) is top-left occupied and dimensions are correct
	return normalizeShape(convertPointsToGrid(rotatedShape.Points, rotatedShape.Width, rotatedShape.Height))
}

// flipShapeHorizontal flips a shape horizontally.
func flipShapeHorizontal(s Shape) Shape {
	if len(s.Points) == 0 {
		return Shape{}
	}
	var flippedPoints []Point
	for _, p := range s.Points {
		flippedPoints = append(flippedPoints, Point{R: p.R, C: s.Width - 1 - p.C})
	}
	flippedShape := Shape{Points: flippedPoints, Width: s.Width, Height: s.Height}
	// Re-normalize after flip
	return normalizeShape(convertPointsToGrid(flippedShape.Points, flippedShape.Width, flippedShape.Height))
}

// convertPointsToGrid converts a slice of Points back to a 2D boolean grid.
func convertPointsToGrid(points []Point, width, height int) [][]bool {
	grid := make([][]bool, height)
	for i := range grid {
		grid[i] = make([]bool, width)
	}
	for _, p := range points {
		if p.R >= 0 && p.R < height && p.C >= 0 && p.C < width {
			grid[p.R][p.C] = true
		}
	}
	return grid
}

// canonicalString creates a unique string representation of an OrientedShape for map keys.
func canonicalString(os OrientedShape) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%dx%d:", os.Width, os.Height))
	// Create a sorted copy of points to ensure canonical representation
	// as the order of points in the slice might vary.
	sortedPoints := make([]Point, len(os.Points))
	copy(sortedPoints, os.Points)
	
	// Manual sort for Point slice (Go's sort.Slice is more efficient than bubble sort)
	// Using a simple insertion sort like algorithm here.
	for i := 1; i < len(sortedPoints); i++ {
		j := i
		for j > 0 && (sortedPoints[j].R < sortedPoints[j-1].R || 
		             (sortedPoints[j].R == sortedPoints[j-1].R && sortedPoints[j].C < sortedPoints[j-1].C)) {
			sortedPoints[j], sortedPoints[j-1] = sortedPoints[j-1], sortedPoints[j]
			j--
		}
	}

	for _, p := range sortedPoints {
		sb.WriteString(fmt.Sprintf("(%d,%d)", p.R, p.C))
	}
	return sb.String()
}


// canFitPresents is the main backtracking function to place all presents.
// It takes the current grid state, remaining presents to place, and all available shape orientations.
func canFitPresents(
	grid [][]bool, 
	presentsToPlace map[int]int, // shapeIdx -> quantity
	allOrientedShapes map[int][]OrientedShape,
) bool {
	// Base case: If no presents left, all fit!
	allPresentsPlaced := true
	for _, quantity := range presentsToPlace {
		if quantity > 0 {
			allPresentsPlaced = false
			break
		}
	}
	if allPresentsPlaced {
		return true
	}

	// Find the next empty cell (top-leftmost for consistency)
	// This optimization helps prune the search space.
	nextR, nextC := -1, -1
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if !grid[r][c] {
				nextR, nextC = r, c
				break
			}
		}
		if nextR != -1 {
			break
		}
	}

	// If no empty cells but presents remain, cannot fit
	// This covers the case where all usable space is filled.
	if nextR == -1 { 
		return false
	}
	
	// Iterate through presents to place (try most constrained first?)
	// For now, simple iteration.
	for shapeIdx, quantity := range presentsToPlace {
		if quantity <= 0 {
			continue // No more of this shape to place
		}

		// Try every orientation of the current shape
		for _, orientedShape := range allOrientedShapes[shapeIdx] {
			// Optimized Iteration: Iterate over all points in the oriented shape.
			// If `p` is an occupied point in `orientedShape` (relative to its (0,0)),
			// and we align `p` with `(nextR, nextC)` in the grid,
			// then the anchor for `orientedShape` would be `(nextR - p.R, nextC - p.C)`. 
			// This ensures we try to cover the `(nextR, nextC)` cell.
			for _, p := range orientedShape.Points {
				anchorR := nextR - p.R
				anchorC := nextC - p.C

				if canPlace(grid, orientedShape, anchorR, anchorC) {
					// Place the present
					placePresent(grid, orientedShape, anchorR, anchorC, true)
					
					// Update presentsToPlace for the recursive call
					presentsToPlace[shapeIdx]--

					// Recurse
					if canFitPresents(grid, presentsToPlace, allOrientedShapes) {
						return true // Found a solution!
					}

					// Backtrack: Remove the present and revert presentsToPlace count
					placePresent(grid, orientedShape, anchorR, anchorC, false)
					presentsToPlace[shapeIdx]++
				}
			}
		}
	}

	return false // No solution found from this path
}

// canPlace checks if an oriented shape can be placed at (anchorR, anchorC)
func canPlace(grid [][]bool, os OrientedShape, anchorR, anchorC int) bool {
	// Check if the bounding box fits within the grid
	if anchorR < 0 || anchorR+os.Height > len(grid) || 
	   anchorC < 0 || anchorC+os.Width > len(grid[0]) {
		return false
	}

	for _, p := range os.Points {
		r, c := anchorR+p.R, anchorC+p.C
		// Check bounds (should be covered by bounding box check, but defensive)
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
			return false 
		}
		// Check for overlap with existing presents
		if grid[r][c] {
			return false
		}
	}
	return true
}

// placePresent places (or removes) an oriented shape from the grid.
func placePresent(grid [][]bool, os OrientedShape, anchorR, anchorC int, place bool) {
	for _, p := range os.Points {
		r, c := anchorR+p.R, anchorC+p.C
		grid[r][c] = place
	}
}

// SolvePart1 solves Day 12 Part 1
func SolvePart1(input string) int {
	allOrientedShapes, regions, err := parseInput(input)
	if err != nil {
		fmt.Printf("Error parsing input: %v\n", err)
		return 0
	}

	regionsThatFit := 0
	for _, region := range regions {
		grid := make([][]bool, region.Height)
		for i := range grid {
			grid[i] = make([]bool, region.Width)
		}
		
		// Make a copy of presentsToFit as the backtracking function modifies it
		presentsCopy := make(map[int]int)
		for k, v := range region.PresentsToFit {
			presentsCopy[k] = v
		}

		if canFitPresents(grid, presentsCopy, allOrientedShapes) {
			regionsThatFit++
		}
	}
	return regionsThatFit
}
