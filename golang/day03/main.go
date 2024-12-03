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
}

func part1(filename string) {
	matches := processInput(filename)
	r := regexp.MustCompile(`\d{1,3}`)
	answer := 0

	for _, match := range matches {
		digits := r.FindAllString(match, -1)
		a, _ := strconv.Atoi(digits[0])
		b, _ := strconv.Atoi(digits[1])

		answer += a * b
	}

	fmt.Printf("Part 1: %d\n", answer)
}

func processInput(filename string) []string {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		message := fmt.Sprintf("Error reading file: %v\n", err)
		panic(message)
	}

	r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	return r.FindAllString(string(bytes), -1)
}
