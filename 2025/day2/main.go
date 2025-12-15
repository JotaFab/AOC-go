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

	pos := 50
	count := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		dir := line[0]
		dist, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		// 1) Full cycles (each cycle hits 0 once)
		cycles := dist / 100
		count += cycles

		// 2) Remaining clicks
		steps := dist % 100
		step := 1
		if dir == 'L' {
			step = -1
		}

		for i := 0; i < steps; i++ {
			pos = (pos + step) % 100
			if pos < 0 {
				pos += 100
			}
			if pos == 0 {
				count++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(count)
}
