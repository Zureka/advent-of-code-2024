package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "input.txt"
	part1(filename)
	part2(filename)
}

func part1(filename string) {
	reports := processInput(filename)
	answer := 0

	for _, report := range reports {
		valid, _ := validateReport(report)
		if valid {
			answer++
		}
	}

	message := fmt.Sprintf("Part 1: %d", answer)
	fmt.Println(message)
}

func part2(filename string) {
	reports := processInput(filename)
	answer := 0

	for _, report := range reports {
		valid, invalidIndex := validateReport(report)

		if valid {
			answer++
		} else {
			newReport1 := generateNewReport(report, int(math.Max(float64(invalidIndex-1), 0)))
			newReport2 := generateNewReport(report, invalidIndex)
			newReport3 := generateNewReport(report, invalidIndex+1)
			valid1, _ := validateReport(newReport1)
			valid2, _ := validateReport(newReport2)
			valid3, _ := validateReport(newReport3)

			if valid1 || valid2 || valid3 {
				answer++
			}
		}
	}

	message := fmt.Sprintf("Part 2: %d", answer)
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

	return reports
}

func validateReport(report []int) (bool, int) {
	var levelType string
	if report[0] > report[1] {
		levelType = "increase"
	} else {
		levelType = "decrease"
	}

	if levelType == "increase" {
		for i, level := range report {
			validIndex := i != len(report)-1

			if validIndex && level <= report[i+1] {
				return false, i
			} else if validIndex && math.Abs(float64(level-report[i+1])) > 3.0 {
				return false, i
			}
		}
	} else {
		for i, level := range report {
			validIndex := i != len(report)-1
			if validIndex && level >= report[i+1] {
				return false, i
			} else if validIndex && math.Abs(float64(level-report[i+1])) > 3.0 {
				return false, i
			}
		}
	}

	return true, -1
}

func generateNewReport(report []int, invalidIndex int) []int {
	newReport := []int{}

	for i, level := range report {
		if i != invalidIndex {
			newReport = append(newReport, level)
		}
	}

	return newReport
}
