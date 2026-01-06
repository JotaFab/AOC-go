package day9

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(lines []string) int {
	type tile struct {
		x, y int
	}

	tiles := make([]tile, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			continue
		}
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		tiles = append(tiles, tile{x, y})
	}

	maxArea := 0
	n := len(tiles)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dx := tiles[i].x - tiles[j].x
			if dx < 0 {
				dx = -dx
			}
			dy := tiles[i].y - tiles[j].y
			if dy < 0 {
				dy = -dy
			}

			area := (dx + 1) * (dy + 1)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

func Part2(lines []string) int64 {
	type Point struct {
		X, Y int
	}

	// ---------- Parse red points ----------
	reds := make([]Point, 0, len(lines))
	for _, line := range lines {
		var x, y int
		fmt.Sscanf(strings.TrimSpace(line), "%d,%d", &x, &y)
		reds = append(reds, Point{x, y})
	}

	n := len(reds)

	minX, maxX := reds[0].X, reds[0].X
	minY, maxY := reds[0].Y, reds[0].Y
	for _, p := range reds {
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	type Edge struct {
		X1, Y1, X2, Y2 int
	}

	edges := make([]Edge, 0, n)
	for i := 0; i < n; i++ {
		a := reds[i]
		b := reds[(i+1)%n]
		edges = append(edges, Edge{a.X, a.Y, b.X, b.Y})
	}

	// ---------- Point on segment ----------
	onSegment := func(px, py int, e Edge) bool {
		if e.X1 == e.X2 { // vertical
			if px != e.X1 {
				return false
			}
			return py >= min(e.Y1, e.Y2) && py <= max(e.Y1, e.Y2)
		}
		if e.Y1 == e.Y2 { // horizontal
			if py != e.Y1 {
				return false
			}
			return px >= min(e.X1, e.X2) && px <= max(e.X1, e.X2)
		}
		return false
	}

	// ---------- Point-in-polygon (inclusive) ----------
	inside := func(px, py int) bool {
		// Boundary check first
		for _, e := range edges {
			if onSegment(px, py, e) {
				return true
			}
		}

		crossings := 0
		for _, e := range edges {
			x1, y1 := e.X1, e.Y1
			x2, y2 := e.X2, e.Y2

			if y1 == y2 {
				continue
			}
			if y1 > y2 {
				x1, x2 = x2, x1
				y1, y2 = y2, y1
			}
			if py < y1 || py >= y2 {
				continue
			}
			if x1 == x2 && px < x1 {
				crossings++
			}
		}
		return crossings&1 == 1
	}

	// Precompute inside grid
	grid := make([][]bool, maxX-minX+1)
	for i := range grid {
		grid[i] = make([]bool, maxY-minY+1)
	}
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			grid[x-minX][y-minY] = inside(x, y)
		}
	}
	var maxArea int64 = 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			a := reds[i]
			b := reds[j]

			x1, x2 := a.X, b.X
			y1, y2 := a.Y, b.Y
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			if y1 > y2 {
				y1, y2 = y2, y1
			}

			area := int64(x2-x1+1) * int64(y2-y1+1)
			if area <= maxArea {
				continue
			}

			valid := true
			// Check perimeter
			for xx := x1; xx <= x2 && valid; xx++ {
				if !grid[xx-minX][y1-minY] || !grid[xx-minX][y2-minY] {
					valid = false
				}
			}
			for yy := y1; yy <= y2 && valid; yy++ {
				if !grid[x1-minX][yy-minY] || !grid[x2-minX][yy-minY] {
					valid = false
				}
			}

			if valid {
				maxArea = area
			}
		}
	}

	return maxArea
}
