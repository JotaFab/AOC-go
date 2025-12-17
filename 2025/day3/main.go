package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 2 {
			continue
		}
		total += maxJoltage(line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(total)
}

func maxJoltage(s string) int {
	n := len(s)

	// tens digit: try from 9 down to 1
	for tens := byte('9'); tens >= '1'; tens-- {
		for i := 0; i < n-1; i++ {
			if s[i] == tens {
				// find best ones digit after position i
				bestOnes := byte('0')
				for j := i + 1; j < n; j++ {
					if s[j] > bestOnes {
						bestOnes = s[j]
					}
				}
				if bestOnes != '0' {
					return int(tens-'0')*10 + int(bestOnes-'0')
				}
			}
		}
	}

	return 0
}
