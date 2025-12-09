package main

import (
	"bufio"
	"log"
	"os"
	//"strings"
	"regexp"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var ans int

	var do bool = true
	for scanner.Scan() {
		line := scanner.Text()
		// Using regexp to get the mul strings
		re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
		muls := re.FindAllStringSubmatch(line, -1)

		for _, mul := range muls {
			log.Println(mul)
			if mul[0] == "don't()" {
				do = false
			}
			if mul[0] == "do()" {
				do = true
			}
			if mul[0] != "do()" && do == true {
				x, err := strconv.Atoi(mul[1])
				y, err := strconv.Atoi(mul[2])
				if err != nil {
					log.Println("Error converting to number", err)
				}
				ans += x * y
			}
		}
	}
	log.Println("Total Sum:", ans)

}
