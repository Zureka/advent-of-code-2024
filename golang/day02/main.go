package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "test.txt"
	part1(filename)
}

func part1(filename string) {
	reports := processInput(filename)
	answer := 0

	for _, report := range reports {
		if (verifyAllDecrease(report) || verifyAllIncrease(report)) && validateDifferenceThresholds(report) {
			answer++
		}
	}

	message := fmt.Sprintf("Part 1: %d", answer)
	fmt.Println(message)
}

func processInput(filename string) [][]int {
	fileBytes, err := os.ReadFile(filename)

	if err != nil {
		message := fmt.Sprintf("Error reading file: %v", err)
		panic(message)
	}

	lines := strings.Split(string(fileBytes), "\n")
	reports := [][]int{}

	for _, line := range lines {
		results := strings.Split(line, " ")
		levels := []int{}

		for _, result := range results {
			resultNum, err := strconv.Atoi(result)
			if err != nil {
				fmt.Println("Error converting string to int")
			}

			levels = append(levels, resultNum)
		}

		reports = append(reports, levels)
	}

	message := fmt.Sprintf("Reports: %v", reports)
	fmt.Println(message)

	return reports
}

func verifyAllIncrease(report []int) bool {
	result := true

	for i, level := range report {
		if i != 0 && level >= report[i-1] {
			result = false
			break
		}
	}

	return result
}

func verifyAllDecrease(report []int) bool {
	result := true

	for i, level := range report {
		if i != 0 && level <= report[i-1] {
			result = false
			break
		}
	}

	return result
}

func validateDifferenceThresholds(report []int) bool {
	result := true

	for i, level := range report {
		if i > 0 && i < len(report)-1 {
			before := report[i-1] - level
			after := level - report[i+1]

			if math.Abs(float64(before)) > 3.0 || math.Abs(float64(after)) > 3.0 {
				result = false
				break
			}
		}
	}

	return result
}
