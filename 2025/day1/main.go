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

		switch dir {
		case 'L':
			pos = (pos - dist) % 100
		case 'R':
			pos = (pos + dist) % 100
		}

		// Go can produce negative modulo results
		if pos < 0 {
			pos += 100
		}

		if pos == 0 {
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(count)
}
