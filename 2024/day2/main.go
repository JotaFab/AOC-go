package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
)

func isReallySafe(nums []int) bool{
	if isSafe(nums){
		return true
	}

	for i := 0; i < len(nums); i++ {
		newNums := make([]int, 0, len(nums)-1)
		newNums = append(newNums, nums[:i]...) // Add elements before the current index
		newNums = append(newNums, nums[i+1:]...) // Add elements after the current index

		if isSafe(newNums) {
			return true
		}
	}
	return false
}

func isSafe(nums []int) bool {
	// Calculate the first difference (level)
	prevDiff := nums[1] - nums[0]
	isIncreasing := prevDiff > 0 // Determine if the sequence starts increasing

	for i := 1; i < len(nums); i++ {
		currDiff := nums[i] - nums[i-1]
		// Check if the sequence remains either increasing or decreasing
		if (isIncreasing && currDiff <= 0) || (!isIncreasing && currDiff >= 0) {
			return false
		}

		// Ensure the difference between consecutive levels is between 1 and 3
		levelDiff := int(math.Abs(float64(currDiff)))
		if levelDiff < 1 || levelDiff > 3 {
			return false
		}

		// Update the previous difference
		prevDiff = currDiff
	}

	return true
}

func main() {
	// Open the file
	file, err := os.Open("numbers.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Initialize arrays to store the two lists

	// Read the file line by line
	scanner := bufio.NewScanner(file)

	var safes int = 0 

	for scanner.Scan() {
		// Split each line into two numbers
		line := strings.TrimSpace(scanner.Text())
		numberStrings := strings.Fields(line)

		var numbers []int
		for _, numStr := range numberStrings {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Error converting number:", numStr, "in line:", line)
				continue
			}
			numbers = append(numbers, num)
		}
		if isReallySafe(numbers) {
			safes++
		}
	}
	
	fmt.Println("Safes: ", safes)

	
}

