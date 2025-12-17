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
	scanner.Scan() // input is one long line
	line := scanner.Text()

	ranges := strings.Split(line, ",")
	var total int64 = 0

	for _, r := range ranges {
		parts := strings.Split(r, "-")
		lo, _ := strconv.ParseInt(parts[0], 10, 64)
		hi, _ := strconv.ParseInt(parts[1], 10, 64)

		total += sumInvalid(lo, hi)
	}

	fmt.Println(total)
}

func sumInvalid(lo, hi int64) int64 {
	seen := make(map[int64]struct{})
	var sum int64 = 0

	minLen := digits(lo)
	maxLen := digits(hi)

	for L := minLen; L <= maxLen; L++ {
		for d := 1; d <= L/2; d++ {
			if L%d != 0 {
				continue
			}

			repeats := L / d
			if repeats < 2 {
				continue
			}

			start := pow10(d - 1)
			end := pow10(d) - 1

			for base := start; base <= end; base++ {
				s := strconv.FormatInt(base, 10)
				var b strings.Builder

				for i := 0; i < repeats; i++ {
					b.WriteString(s)
				}

				val, err := strconv.ParseInt(b.String(), 10, 64)
				if err != nil {
					continue
				}

				if val >= lo && val <= hi {
					if _, ok := seen[val]; !ok {
						seen[val] = struct{}{}
						sum += val
					}
				}
			}
		}
	}

	return sum
}

func digits(n int64) int {
	c := 0
	for n > 0 {
		n /= 10
		c++
	}
	return c
}

func pow10(n int) int64 {
	res := int64(1)
	for i := 0; i < n; i++ {
		res *= 10
	}
	return res
}
