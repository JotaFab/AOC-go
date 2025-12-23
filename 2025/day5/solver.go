package day5

import (
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func ParseRanges(input string) []Range {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var ranges []Range

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}

		parts := strings.Split(line, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		ranges = append(ranges, Range{start, end})
	}

	return ranges
}

/* ---------- Part 1 ---------- */

func CountFreshIngredients(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	var ranges []Range
	var values []int
	parsingRanges := true

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" {
			parsingRanges = false
			continue
		}

		if parsingRanges {
			parts := strings.Split(line, "-")
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])
			ranges = append(ranges, Range{start, end})
		} else {
			val, _ := strconv.Atoi(line)
			values = append(values, val)
		}
	}

	count := 0
	for _, v := range values {
		if isFresh(v, ranges) {
			count++
		}
	}

	return count
}

func isFresh(v int, ranges []Range) bool {
	for _, r := range ranges {
		if v >= r.start && v <= r.end {
			return true
		}
	}
	return false
}

/* ---------- Part 2 ---------- */

func CountFreshIDsFromRanges(input string) int {
	ranges := ParseRanges(input)
	if len(ranges) == 0 {
		return 0
	}

	// Sort by start
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	total := 0
	curStart := ranges[0].start
	curEnd := ranges[0].end

	for _, r := range ranges[1:] {
		if r.start <= curEnd+1 {
			// overlap or adjacency
			if r.end > curEnd {
				curEnd = r.end
			}
		} else {
			total += curEnd - curStart + 1
			curStart = r.start
			curEnd = r.end
		}
	}

	total += curEnd - curStart + 1
	return total
}
