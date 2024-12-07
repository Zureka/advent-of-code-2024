package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Equation struct {
	answer int
	values []int
}

func main() {
	filename := "input.txt"
	part1(filename)
	part2(filename)
}

func part1(filename string) {
	equations := processInput(filename)
	answer := 0

	for _, equation := range equations {
		if solve1(equation.answer, 0, equation.values, 0) > 0 {
			answer += equation.answer
		}
	}

	fmt.Printf("Part 1: %d\n", answer)
}

func part2(filename string) {
	equations := processInput(filename)
	solvedEquations := []Equation{}
	answer := 0

	resCh := make(chan int)
	var wg sync.WaitGroup
	for _, equation := range equations {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if solve2(equation.answer, 0, equation.values, 0) > 0 {
				solvedEquations = append(solvedEquations, equation)
				resCh <- equation.answer
			}
		}()
	}

	go func() {
		wg.Wait()
		close(resCh)
	}()

	for res := range resCh {
		answer += res
	}

	fmt.Printf("Part 2: %d\n", answer)
}

func solve1(answer int, currentTotal int, values []int, index int) int {
	if index == len(values) && answer == currentTotal {
		return 1
	} else if index == len(values) {
		return 0
	}

	value := values[index]
	return solve1(answer, currentTotal+value, values, index+1) + solve1(answer, currentTotal*value, values, index+1)
}

func solve2(answer int, currentTotal int, values []int, index int) int {
	if index == len(values) && answer == currentTotal {
		return 1
	} else if index == len(values) {
		return 0
	}

	value := values[index]
	concatAnswer := 0
	newValue := 0

	if index != len(values) {
		for i, v := range values {
			if i == index {
				concatValue, _ := strconv.Atoi(fmt.Sprintf("%d%d", currentTotal, v))
				newValue = concatValue
			}
		}

		concatAnswer = solve2(answer, newValue, values, index+1)
	}

	return solve2(answer, currentTotal+value, values, index+1) + solve2(answer, currentTotal*value, values, index+1) + concatAnswer
}

func processInput(filename string) []Equation {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		message := fmt.Sprintf("Error reading file - %s", filename)
		panic(message)
	}

	equations := []Equation{}

	lines := strings.Split(string(bytes), "\n")
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		answer, _ := strconv.Atoi(parts[0])
		valueStrings := strings.Split(parts[1], " ")
		values := []int{}

		for _, v := range valueStrings {
			value, _ := strconv.Atoi(v)
			values = append(values, value)
		}

		equations = append(equations, Equation{answer, values})
	}

	return equations
}
