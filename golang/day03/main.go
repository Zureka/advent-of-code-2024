package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	filename := "input.txt"
	part1(filename)
	part2(filename)
}

func part1(filename string) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		message := fmt.Sprintf("Error reading file: %v\n", err)
		panic(message)
	}

	commandRegex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	digitRegex := regexp.MustCompile(`\d{1,3}`)

	matches := commandRegex.FindAllString(string(bytes), -1)
	answer := 0

	for _, match := range matches {
		digits := digitRegex.FindAllString(match, -1)
		a, _ := strconv.Atoi(digits[0])
		b, _ := strconv.Atoi(digits[1])

		answer += a * b
	}

	fmt.Printf("Part 1: %d\n", answer)
}

func part2(filename string) {
	answer := 0
	enabled := true

	bytes, err := os.ReadFile(filename)
	if err != nil {
		message := fmt.Sprintf("Error reading file: %v\n", err)
		panic(message)
	}

	commandsRegex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
	digitRegex := regexp.MustCompile(`\d{1,3}`)

	matches := commandsRegex.FindAllString(string(bytes), -1)

	for _, match := range matches {
		if match == "do()" {
			enabled = true
		} else if match == "don't()" {
			enabled = false
		} else if enabled {
			digits := digitRegex.FindAllString(match, -1)
			a, _ := strconv.Atoi(digits[0])
			b, _ := strconv.Atoi(digits[1])

			answer += a * b
		}
	}

	fmt.Printf("Part 2: %d\n", answer)
}
