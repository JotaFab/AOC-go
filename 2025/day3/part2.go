package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

const PICK = 12

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := big.NewInt(0)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < PICK {
			continue
		}

		val := maxSubsequence(line, PICK)
		total.Add(total, val)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(total.String())
}

func maxSubsequence(s string, k int) *big.Int {
	n := len(s)
	toRemove := n - k
	stack := make([]byte, 0, n)

	for i := 0; i < n; i++ {
		d := s[i]
		for toRemove > 0 && len(stack) > 0 && stack[len(stack)-1] < d {
			stack = stack[:len(stack)-1]
			toRemove--
		}
		stack = append(stack, d)
	}

	// If removals remain, drop from end
	stack = stack[:k]

	result := new(big.Int)
	result.SetString(string(stack), 10)
	return result
}
