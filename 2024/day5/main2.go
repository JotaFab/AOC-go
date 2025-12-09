package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseRules(lines []string) map[int][]int {
	rules := make(map[int][]int)
	for _, line := range lines {
		parts := strings.Split(line, "|")
		bef, _ := strconv.Atoi(parts[0])
		after, _ := strconv.Atoi(parts[1])
		rules[bef] = append(rules[bef], after)
	}
	return rules
}

func main() {
	flag.Parse()
	file, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var rulesLines []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		rulesLines = append(rulesLines, line)
	}
	parsedRules := parseRules(rulesLines)
	var middlesSum int
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")

		var partsInts []int
		for _, part := range parts {
			partInt, _ := strconv.Atoi(part)
			partsInts = append(partsInts, partInt)
		}
		if !isOrdered(partsInts, parsedRules) {
			updateOrdered := orderUpdate(partsInts, parsedRules)
			middlepart := middlePart(updateOrdered)
			middlesSum += middlepart

		}
	}
	fmt.Println("The sum of all the correct ordered middle numbers is:\n", middlesSum)
}

func middlePart(parts []int) int {
	return parts[len(parts)/2]
}
func orderUpdate(update []int, rules map[int][]int) []int {
	// Create a copy of the update to avoid modifying the original
	orderedUpdate := make([]int, len(update))
	copy(orderedUpdate, update)

	// Attempt to order the pages
	for {
		swapped := false
		for i := 0; i < len(orderedUpdate)-1; i++ {
			for _, after := range rules[orderedUpdate[i]] {
				if contains(orderedUpdate[i+1:], after) {
					// Find the index of the 'after' page
					afterIndex := -1
					for j := i + 1; j < len(orderedUpdate); j++ {
						if orderedUpdate[j] == after {
							afterIndex = j
							break
						}
					}

					// Swap the 'current' and 'after' pages
					orderedUpdate[i], orderedUpdate[afterIndex] = orderedUpdate[afterIndex], orderedUpdate[i]
					swapped = true
					break
				}
			}
			if swapped {
				break
			}
		}

		// If no swaps were made in this pass, the update is ordered
		if !swapped {
			break
		}
	}

	return orderedUpdate
}
func isOrdered(parts []int, rules map[int][]int) bool {
	for i := 1; i < len(parts); i++ {
		for _, after := range rules[parts[i]] {
			fmt.Println(after, parts[:i], i)
			if contains(parts[:i], after) {
				return false
			}
		}
	}
	return true
}

func contains(slice []int, target int) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}
