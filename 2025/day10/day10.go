package day10

import (
	"math/bits"
	"strings"
)

func Part1(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	total := 0

	for _, line := range lines {
		target, buttons := parseLine(line)
		minPresses := solveMachine(target, buttons)
		total += minPresses
	}

	return total
}

func solveMachine(target uint64, buttons []uint64) int {
	nButtons := len(buttons)

	// Build augmented matrix: each row is a light equation
	var rows []uint64
	for i := 0; i < 64; i++ {
		if (target>>i)&1 == 0 {
			continue
		}
		var row uint64
		for j, b := range buttons {
			if (b>>i)&1 == 1 {
				row |= 1 << j
			}
		}
		rows = append(rows, row)
	}

	// Gaussian elimination
	pivot := 0
	where := make([]int, nButtons)
	for i := range where {
		where[i] = -1
	}

	for col := 0; col < nButtons && pivot < len(rows); col++ {
		sel := -1
		for i := pivot; i < len(rows); i++ {
			if (rows[i]>>col)&1 == 1 {
				sel = i
				break
			}
		}
		if sel == -1 {
			continue
		}
		rows[pivot], rows[sel] = rows[sel], rows[pivot]
		where[col] = pivot

		for i := 0; i < len(rows); i++ {
			if i != pivot && ((rows[i]>>col)&1 == 1) {
				rows[i] ^= rows[pivot]
			}
		}
		pivot++
	}

	// Particular solution
	var x0 uint64
	for i, w := range where {
		if w != -1 {
			x0 |= 1 << i
		}
	}

	// Nullspace basis
	var basis []uint64
	for i, w := range where {
		if w == -1 {
			vec := uint64(1 << i)
			for j, w2 := range where {
				if w2 != -1 && ((rows[w2]>>i)&1 == 1) {
					vec |= 1 << j
				}
			}
			basis = append(basis, vec)
		}
	}

	// Minimize presses
	best := bits.OnesCount64(x0)
	k := len(basis)

	for mask := 1; mask < (1 << k); mask++ {
		x := x0
		for i := 0; i < k; i++ {
			if (mask>>i)&1 == 1 {
				x ^= basis[i]
			}
		}
		best = min(best, bits.OnesCount64(x))
	}

	return best
}

func parseLine(line string) (uint64, []uint64) {
	// indicator lights
	start := strings.Index(line, "[")
	end := strings.Index(line, "]")
	lights := line[start+1 : end]

	var target uint64
	for i, c := range lights {
		if c == '#' {
			target |= 1 << i
		}
	}

	// buttons
	var buttons []uint64
	for {
		open := strings.Index(line, "(")
		if open == -1 {
			break
		}
		close := strings.Index(line, ")")
		content := line[open+1 : close]
		line = line[close+1:]

		var mask uint64
		for _, v := range strings.Split(content, ",") {
			if v == "" {
				continue
			}
			idx := atoi(v)
			mask |= 1 << idx
		}
		buttons = append(buttons, mask)
	}

	return target, buttons
}

func atoi(s string) int {
	n := 0
	for _, c := range s {
		n = n*10 + int(c-'0')
	}
	return n
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
