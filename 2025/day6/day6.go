package day6

import (
	"strconv"
	"strings"
	"unicode"
)

func Solve(input string) int {
	lines := strings.Split(strings.TrimRight(input, "\n"), "\n")

	height := len(lines)
	width := 0
	for _, l := range lines {
		if len(l) > width {
			width = len(l)
		}
	}

	// Pad lines
	grid := make([][]rune, height)
	for i, l := range lines {
		row := []rune(l)
		for len(row) < width {
			row = append(row, ' ')
		}
		grid[i] = row
	}

	var problems [][]int
	current := []int{}

	for col := 0; col < width; col++ {
		nonSpace := false
		for row := 0; row < height; row++ {
			if grid[row][col] != ' ' {
				nonSpace = true
				break
			}
		}

		if nonSpace {
			current = append(current, col)
		} else if len(current) > 0 {
			problems = append(problems, current)
			current = []int{}
		}
	}

	if len(current) > 0 {
		problems = append(problems, current)
	}

	total := 0

	for _, cols := range problems {
		var numbers []int
		var op rune

		for row := 0; row < height; row++ {
			var sb strings.Builder
			for _, c := range cols {
				sb.WriteRune(grid[row][c])
			}
			chunk := strings.TrimSpace(sb.String())
			if chunk == "" {
				continue
			}

			if chunk == "+" || chunk == "*" {
				op = rune(chunk[0])
			} else {
				n, _ := strconv.Atoi(chunk)
				numbers = append(numbers, n)
			}
		}

		result := numbers[0]
		for i := 1; i < len(numbers); i++ {
			if op == '+' {
				result += numbers[i]
			} else {
				result *= numbers[i]
			}
		}

		total += result
	}

	return total
}



func SolvePartTwo(input string) int {
	lines := strings.Split(strings.TrimRight(input, "\n"), "\n")

	height := len(lines)
	width := 0
	for _, l := range lines {
		if len(l) > width {
			width = len(l)
		}
	}

	grid := make([][]rune, height)
	for i, l := range lines {
		row := []rune(l)
		for len(row) < width {
			row = append(row, ' ')
		}
		grid[i] = row
	}

	// detect problem column blocks
	var problems [][]int
	var current []int

	for col := 0; col < width; col++ {
		hasChar := false
		for row := 0; row < height; row++ {
			if grid[row][col] != ' ' {
				hasChar = true
				break
			}
		}

		if hasChar {
			current = append(current, col)
		} else if len(current) > 0 {
			problems = append(problems, current)
			current = nil
		}
	}
	if len(current) > 0 {
		problems = append(problems, current)
	}

	total := 0

	for _, cols := range problems {
		var nums []int
		var op rune

		// operator is found at bottom of any column
		for _, c := range cols {
			ch := grid[height-1][c]
			if ch == '+' || ch == '*' {
				op = ch
				break
			}
		}

		// read columns right → left
		for i := len(cols) - 1; i >= 0; i-- {
			col := cols[i]
			var digits []rune

			for row := 0; row < height; row++ {
				ch := grid[row][col]
				if unicode.IsDigit(ch) {
					digits = append(digits, ch)
				}
			}

			if len(digits) == 0 {
				continue
			}

			n, _ := strconv.Atoi(string(digits))
			nums = append(nums, n)
		}

		result := nums[0]
		for i := 1; i < len(nums); i++ {
			if op == '+' {
				result += nums[i]
			} else {
				result *= nums[i]
			}
		}

		total += result
	}

	return total
}
