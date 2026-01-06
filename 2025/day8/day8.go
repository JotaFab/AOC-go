package day8

import (
	"sort"
)

type Point struct {
	X, Y, Z int
}

type edge struct {
	a, b int
	dist int
}

func SolveDay8(points []Point) int {
	n := len(points)
	if n < 3 {
		return 0
	}

	// build all edges
	edges := make([]edge, 0, n*(n-1)/2)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dx := points[i].X - points[j].X
			dy := points[i].Y - points[j].Y
			dz := points[i].Z - points[j].Z
			dist := dx*dx + dy*dy + dz*dz

			edges = append(edges, edge{i, j, dist})
		}
	}

	// sort by distance
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dist < edges[j].dist
	})

	// DSU
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(a, b int) {
		ra := find(a)
		rb := find(b)
		if ra == rb {
			return
		}
		if size[ra] < size[rb] {
			ra, rb = rb, ra
		}
		parent[rb] = ra
		size[ra] += size[rb]
	}

	// process first 1000 edges (successful or not)
	limit := 1000
	if len(edges) < limit {
		limit = len(edges)
	}

	for i := 0; i < limit; i++ {
		union(edges[i].a, edges[i].b)
	}

	// collect component sizes
	comp := make(map[int]int)
	for i := 0; i < n; i++ {
		root := find(i)
		comp[root]++
	}

	sizes := make([]int, 0, len(comp))
	for _, s := range comp {
		sizes = append(sizes, s)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	return sizes[0] * sizes[1] * sizes[2]
}


func SolveDay8Part2(points []Point) int {
	n := len(points)
	if n < 2 {
		return 0
	}

	type edge struct {
		a, b int
		dist int
	}

	edges := make([]edge, 0, n*(n-1)/2)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dx := points[i].X - points[j].X
			dy := points[i].Y - points[j].Y
			dz := points[i].Z - points[j].Z
			dist := dx*dx + dy*dy + dz*dz

			edges = append(edges, edge{i, j, dist})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dist < edges[j].dist
	})

	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	components := n

	for _, e := range edges {
		ra := find(e.a)
		rb := find(e.b)
		if ra == rb {
			continue
		}

		// union
		if size[ra] < size[rb] {
			ra, rb = rb, ra
		}
		parent[rb] = ra
		size[ra] += size[rb]
		components--

		// this edge made the graph fully connected
		if components == 1 {
			return points[e.a].X * points[e.b].X
		}
	}

	return 0
}