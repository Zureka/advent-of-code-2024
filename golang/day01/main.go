package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	filename := "input.txt"
	part1(filename)
	part2(filename)
}

func part1(filename string) {
	leftNums, rightNums := processFile(filename)
	answer := 0

	for i, left := range leftNums {
		right := rightNums[i]
		answer += int(math.Abs(float64(left - right)))
	}

	message := fmt.Sprintf("Part 1: %d", answer)
	fmt.Println(message)
}

func part2(filename string) {
	leftNums, rightNums := processFile(filename)
	answer := 0

	for _, left := range leftNums {
		count := 0
		for _, right := range rightNums {
			if left == right {
				count++
			}
		}
		answer += (left * count)
	}

	message := fmt.Sprintf("Part 2: %d", answer)
	fmt.Println(message)
}

func processFile(filename string) ([]int, []int) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		message := fmt.Sprintf("Error reading file: %s", filename)
		panic(message)
	}

	content := string(bytes)
	leftNums := []int{}
	rightNums := []int{}

	for _, line := range strings.Split(content, "\n") {
		parts := strings.Split(line, "   ")
		left, err := strconv.Atoi(parts[0])

		if err != nil {
			message := fmt.Sprintf("Error converting string to int: %s", parts[0])
			panic(message)
		}
		leftNums = append(leftNums, left)

		right, err := strconv.Atoi(parts[1])
		if err != nil {
			message := fmt.Sprintf("Error converting string to int: %s", parts[0])
			panic(message)
		}
		rightNums = append(rightNums, right)
	}

	sort.Ints(leftNums)
	sort.Ints(rightNums)

	return leftNums, rightNums
}
