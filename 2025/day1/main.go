package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	n := 50

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split each line into two numbers
		line := strings.TrimSpace(scanner.Text())

		fmt.Println(line)
		for l := range line {
			fmt.Println(l)

		}
	}

}


func solver()


