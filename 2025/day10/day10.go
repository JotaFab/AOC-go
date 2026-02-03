package day10

import (
	"fmt"
	"log"
	"math/bits"
	"sort"
	"strconv"
	"strings"
)

// machine represents a single machine with its target light state, buttons, and joltage requirements
type machine struct {
	desired uint  // For Part 1 (lights on/off)
	buttons []uint // Each uint is a bitmask of lights/counters toggled by a button
	joltage []int  // For Part 2 (target joltage levels)
	num     int    // Machine index in input (e.g., 0 for first machine, 1 for second, etc.)
}

// parseMachine parses a single line of input into a machine struct
func parseMachine(line string) machine {
	var m machine
	words := strings.Fields(line)
	first := words[0]
	
	// Determine the number of lights/counters from the diagram length
	// This will be used for bounds checking in parseButton and for array sizing.
	numIndicators := len(first) - 2 // -2 for the '[' and ']'
	
	// Parse desired light pattern for Part 1
	for i, c := range first[1 : len(first)-1] {
		if c == '#' {
			m.desired |= 1 << i
		}
	}
	
	// Parse button wiring schematics
	for _, w := range words[1 : len(words)-1] {
		if strings.HasPrefix(w, "(") && strings.HasSuffix(w, ")") {
			// Pass numIndicators to parseButton for bounds checking
			m.buttons = append(m.buttons, parseButton(w, numIndicators))
		}
	}
	
	last := words[len(words)-1]
	// Parse joltage requirements for Part 2
	if strings.HasPrefix(last, "{") && strings.HasSuffix(last, "}") {
		for _, s := range strings.Split(last[1:len(last)-1], ",") {
			i, err := strconv.Atoi(s)
			if err != nil {
				// Use log.Fatalf which is fine in main, but for a package, consider returning error
				// For AOC, it's typically fine for parsing errors to be fatal if input is guaranteed valid.
				log.Fatalf("Non-numeric joltage in %s: %v", last, err)
			}
			m.joltage = append(m.joltage, i)
		}
	}
	return m
}

// parseButton now takes numLights to validate indices
func parseButton(word string, numIndicators int) uint {
	var b uint
	for _, s := range strings.Split(word[1:len(word)-1], ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Non-integer button index in %s: %v", word, err)
		}
		if i < 0 || i >= numIndicators {
			log.Fatalf("Button index %d out of bounds (0-%d) for button %s", i, numIndicators-1, word)
		}
		b |= 1 << i
	}
	return b
}

// machinePresses1 solves Part 1 for a single machine using BFS
func machinePresses1(m machine) int {
	// pressState stores the bitmask of buttons pressed and the resulting light set
	type pressState struct{ pressed, set uint }
	
	// If the desired state is all lights off, and we start all lights off, 0 presses.
	if m.desired == 0 {
		return 0
	}
	
	// Queue for BFS
	queue := []pressState{{pressed: 0, set: 0}}
	// Visited map: key is the current 'set' of lights, value is the minimum presses to reach that 'set'
	visited := map[uint]int{0: 0} 

	head := 0
	for head < len(queue) {
		s := queue[head]
		head++

		for i, b := range m.buttons {
			// If this button has already been part of the 'pressed' combination for this state, skip.
			// This check is for cycle detection and ensuring minimum path.
			if (s.pressed>>i)&1 == 1 { 
				continue
			}

			nextSet := s.set ^ b // Toggle lights with this button
			nextPressed := s.pressed | (1 << i) // Mark this button as pressed
			nextPressCount := bits.OnesCount(nextPressed) // Count presses

			// If we reached the desired state, return the press count
			if nextSet == m.desired {
				return nextPressCount
			}

			// If not visited, or found a shorter path to it
			if currentMin, exists := visited[nextSet]; !exists || nextPressCount < currentMin {
				visited[nextSet] = nextPressCount
				queue = append(queue, pressState{pressed: nextPressed, set: nextSet})
			}
		}
	}
	return -1 // Should not happen for solvable problems
}

// SolvePart1 processes all machines for Part 1
func SolvePart1(lines []string) string {
	var machines []machine
	for i, l := range lines {
		m := parseMachine(l)
		m.num = i // Store original index if needed, not num lights
		machines = append(machines, m)
	}
	var sum int
	for _, m := range machines {
		x := machinePresses1(m)
		sum += x
	}
	return strconv.Itoa(sum)
}

// --- Part 2 Logic (from your provided code) ---

type matrix [][]int

func newMatrix(rows, cols int) matrix {
	m := make([][]int, rows)
	for i := 0; i < rows; i++ {
		m[i] = make([]int, cols)
	}
	return m
}

func (m matrix) String() string {
	var b strings.Builder
	for i, r := range m {
		if i > 0 {
			b.WriteRune('\n')
		}
		for j, c := range r {
			if j > 0 {
				b.WriteRune('\t')
			}
			b.WriteString(strconv.Itoa(c))
		}
	}
	return b.String()
}

func (m matrix) rows() int { return len(m) }
func (m matrix) cols() int { return len(m[0]) }

func (m matrix) swap(i, j int) matrix {
	n := make(matrix, m.rows())
	copy(n, m)
	n[i], n[j] = m[j], m[i]
	return n
}

func (m matrix) colswap(a, b int) matrix {
	n := newMatrix(m.rows(), m.cols())
	for i := range n {
		copy(n[i], m[i])
		n[i][a], n[i][b] = n[i][b], n[i][a]
	}
	return n
}

func (m matrix) negate(i int) matrix {
	n := make(matrix, m.rows())
	copy(n, m)
	n[i] = make([]int, m.cols())
	for c := 0; c < m.cols(); c++ {
		n[i][c] = m[i][c] * -1
	}
	return n
}

func (m matrix) add(dest, src, factor int) matrix {
	n := make(matrix, m.rows())
	copy(n, m)
	n[dest] = make([]int, m.cols())
	for c := 0; c < m.cols(); c++ {
		n[dest][c] = m[dest][c] + m[src][c]*factor
	}
	return n
}

func (m matrix) rightmost(row int) int {
	for i := m.cols() - 1; i >= 0; i-- {
		if m[row][i] != 0 {
			return i
		}
	}
	return -1
}

type freevars struct {
	vars map[int]bool
	vals [13]int // Fixed size array based on your code
}

func (f freevars) set(i int) bool { return f.vars[i] }

func (f freevars) numset() int { return len(f.vars) }

func (f freevars) successors(m matrix, worst int, seen map[[13]int]bool) []freevars {
	res := make([]freevars, 0, f.numset())
	if sumSlice(f.vals[:]) >= worst {
		return res
	}
	for i := range f.vars {
		ok := false
		allneg := true
		for _, r := range m {
			if r[i] != 0 {
				v := r[len(r)-1]
				for c := 0; c < m.cols()-1; c++ {
					v -= r[c] * f.vals[c]
				}
				u := v - r[i]
				if absInt(v) > absInt(u) {
					ok = true
				}
				if r[i] > 0 && r[len(r)-1] > 0 {
					allneg = false
				}
			}
		}
		if !ok && !allneg {
			continue
		}
		x := freevars{vars: f.vars, vals: f.vals}
		x.vals[i]++
		if seen[x.vals] {
			continue
		}
		res = append(res, x)
	}
	return res
}

func (f freevars) backSubstitute(m matrix) (sum, rem int, ok bool) {
	ok = true
	all := f.vals
	for r := m.rows() - 1; r >= 0; r-- {
		v := m[r][m.cols()-1]
		for c := m.cols() - 2; c > r; c-- { // Corrected inner loop to c--
			v -= all[c] * m[r][c]
		}
		if r > m.cols()-1 || m[r][r] == 0 {
			if v != 0 {
				rem += absInt(v)
				ok = false
			}
		} else {
			if v < 0 {
				ok = false
				rem += -1 * v
			}
			if v%m[r][r] == 0 {
				all[r] = v / m[r][r]
			} else {
				ok = false
			}
		}
	}
	sum = sumSlice(all[:])
	return
}

func reduce(m matrix) matrix {
	col := 0
outer:
	for col < minInt(m.rows(), m.cols()) {
		for i := col + 1; i < m.rows(); i++ {
			if m[col][col] == 0 && m[i][col] != 0 {
				m = m.swap(col, i)
			} else if m[i][col] != 0 && absInt(m[i][col]) < absInt(m[col][col]) {
				m = m.swap(col, i)
			}
		}
		if m[col][col] == 0 {
			for c := col + 1; c < m.cols()-1; c++ {
				if m[col][c] != 0 {
					m = m.colswap(col, c)
					continue outer
				}
			}
		}
		if m[col][col] == 0 {
			for i := col; i < m.rows(); i++ {
				if m.rightmost(i) == -1 {
					if i != col {
						m = m.swap(col, i)
					}
					break
				}
			}
			col++
			continue
		}
		if m[col][col] < 0 {
			m = m.negate(col)
		}
		for i := 0; i < m.rows(); i++ {
			if i == col {
				continue
			}
			if m[i][col] != 0 {
				factor := m[i][col] / m[col][col]
				m = m.add(i, col, factor*-1)
				if i > col && m[i][col] != 0 {
					continue outer 
				}
			}
		}
		col++
	}
	return m
}

func machinePart2(x machine) int {
	buts := make([]uint, len(x.buttons))
	copy(buts, x.buttons)
	sort.SliceStable(buts, func(i, j int) bool {
		return bits.OnesCount(buts[i]) > bits.OnesCount(buts[j])
	})
	numrows := len(x.joltage)
	numcols := len(buts) + 1
	m := newMatrix(numrows, numcols)
	for i := range x.joltage {
		for j, b := range buts {
			if b&(1<<i) != 0 {
				m[i][j] = 1
			}
		}
	}
	for i, v := range x.joltage {
		m[i][numcols-1] = v
	}
	sort.SliceStable(m, func(i, j int) bool {
		var a, b int
		for c := 0; c < numcols-1; c++ {
			a += m[i][c]
			b += m[j][c]
		}
		if a == b {
			return m[i][numcols-1] < m[j][numcols-1]
		}
		return a < b
	})
	m = reduce(m)
	initial := freevars{vars: map[int]bool{}}
	for i := 0; i < m.cols()-1; i++ {
		if i >= m.rows() || m[i][i] == 0 {
			initial.vars[i] = true
		}
	}
	initsum, initrem, initok := initial.backSubstitute(m)
	if initial.numset() == 0 {
		if initok && initrem == 0 {
			return initsum
		}
		panic(fmt.Sprintf("no freevars sum=%d rem=%d ok=%t free=%+v", initsum, initrem, initok, initial))
	}
	worst := sumSlice(x.joltage)
	best := worst
	if initok && initrem == 0 {
		best = initsum
	}
	seen := map[[13]int]bool{initial.vals: true}
	q := make([]freevars, 0, 1024)
	q = append(q, initial)
	for len(q) != 0 {
		cur := q[0]
		q = q[1:]
		for _, f := range cur.successors(m, worst, seen) {
			seen[f.vals] = true
			sum, rem, ok := f.backSubstitute(m)
			if ok && rem == 0 {
				best = minInt(best, sum)
			}
			if sum < worst && sumSlice(f.vals[:]) < best {
				q = append(q, f)
			}
		}
	}
	if best == worst {
		log.Printf("%d: UNLIKELY!!! best==worst %d", x.num, best)
	}
	return best
}

// SolvePart2 processes all machines for Part 2
func SolvePart2(lines []string) string {
	var machines []machine
	for i, l := range lines {
		m := parseMachine(l)
		m.num = i + 1 // Original code used i+1 for part2 num
		machines = append(machines, m)
	}
	var sum int
	for _, m := range machines {
		x := machinePart2(m)
		sum += x
	}
	return strconv.Itoa(sum)
}

func absInt(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}

func maxInt(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func sumSlice(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}
