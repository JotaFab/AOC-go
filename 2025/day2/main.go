package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan() // entire input is one line
	line := scanner.Text()

	ranges := strings.Split(line, ",")
	var total int64 = 0

	for _, r := range ranges {
		parts := strings.Split(r, "-")
		lo, _ := strconv.ParseInt(parts[0], 10, 64)
		hi, _ := strconv.ParseInt(parts[1], 10, 64)

		total += sumInvalidInRange(lo, hi)
	}

	fmt.Println(total)
}

func sumInvalidInRange(lo, hi int64) int64 {
	var sum int64 = 0

	// Determine digit lengths
	minLen := numDigits(lo)
	maxLen := numDigits(hi)

	for length := minLen; length <= maxLen; length++ {
		if length%2 != 0 {
			continue
		}

		half := length / 2
		start := pow10(half - 1)
		end := pow10(half) - 1

		for i := start; i <= end; i++ {
			s := strconv.FormatInt(i, 10)
			full, _ := strconv.ParseInt(s+s, 10, 64)

			if full >= lo && full <= hi {
				sum += full
			}
		}
	}

	return sum
}

func numDigits(n int64) int {
	count := 0
	for n > 0 {
		n /= 10
		count++
	}
	return count
}

func pow10(n int) int64 {
	res := int64(1)
	for i := 0; i < n; i++ {
		res *= 10
	}
	return res
}
