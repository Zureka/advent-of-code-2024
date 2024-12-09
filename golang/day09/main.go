package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := "test.txt"
	part1(filename)
	part2(filename)
}

func part1(filename string) {
	answer := 0
	lines := processInput(filename)
	fmt.Printf("%v\n", lines)
	fmt.Printf("Part 1: %d\n", answer)
}
func part2(filename string) {
	answer := 0
	lines := processInput(filename)
	fmt.Printf("%v\n", lines)
	fmt.Printf("Part 2: %d\n", answer)
}

func processInput(filename string) []string {
	bytes, err := os.ReadFile(filename)

	if err != nil {
		message := fmt.Sprintf("Error reading file: %v", err)
		panic(message)
	}

	lines := strings.Split(string(bytes), "\n")

	return lines
}
