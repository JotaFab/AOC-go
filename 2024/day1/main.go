package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Open the file
	file, err := os.Open("numbers.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Initialize arrays to store the two lists
	var list1, list2 []int

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split each line into two numbers
		line := strings.TrimSpace(scanner.Text())
		numbers := strings.Fields(line)
		if len(numbers) != 2 {
			fmt.Println("Invalid line format:", line)
			continue
		}

		// Parse the numbers and add them to their respective lists
		num1, err1 := strconv.Atoi(numbers[0])
		num2, err2 := strconv.Atoi(numbers[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Error parsing numbers in line:", line)
			continue
		}
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	// Check for errors in reading the file
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Sort both lists
	sort.Ints(list1)
	sort.Ints(list2)

	// Display the sorted lists
	fmt.Println("Sorted List 1:", list1)
	fmt.Println("Sorted List 2:", list2)
	var sim int
	for i := 0; i < len(list1); i++ {
		for j:= 0; j < len(list2); j++ {
			if list1[i] == list2[j] {
				sim += list1[i]
			}
		}
	}	
	fmt.Printf("The total similarity is: %d\n", sim)

}

