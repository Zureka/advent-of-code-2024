package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := "input.txt"
	part1(filename)
	part2(filename)
}

func part1(filename string) {
	grid := processInput(filename)
	currentPosition := getCurrentPosition(grid)
	processPath(currentPosition, grid, false)

	answer := 0

	for _, row := range grid {
		for _, cell := range row {
			if cell == "X" {
				answer++
			}
		}
	}

	fmt.Printf("Part 1: %d\n", answer)
}

func part2(filename string) {
	originalGrid := processInput(filename)
	gridCopy := copyGrid(originalGrid)
	startPosition := getCurrentPosition(gridCopy)
	processPath(startPosition, gridCopy, false)

	traveledPositions := [][]int{}

	for i, row := range gridCopy {
		for j, cell := range row {
			if cell == "X" && !(i == startPosition[0] && j == startPosition[1]) {
				traveledPositions = append(traveledPositions, []int{i, j})
			}
		}
	}

	answer := 0

	for _, position := range traveledPositions {
		newGrid := copyGrid(originalGrid)
		newGrid[position[0]][position[1]] = "#"
		reachedEnd := processPath(startPosition, newGrid, false)

		if !reachedEnd {
			answer++
		}
	}

	fmt.Printf("Part 2: %d\n", answer)
}

func processPath(startPosition []int, grid [][]string, shouldPrint bool) bool {
	position := startPosition
	direction := grid[position[0]][position[1]]
	reachedEnd := false
	numSteps := 0
	maxSteps := 10000

	for !reachedEnd {
		if numSteps >= maxSteps {
			break
		}

		if direction == "^" {
			if position[0]-1 < 0 {
				grid[position[0]][position[1]] = "X"
				reachedEnd = true
			} else if grid[position[0]-1][position[1]] == "#" {
				direction = turnRight(direction)
				grid[position[0]][position[1]] = direction
			} else {
				grid[position[0]-1][position[1]] = "^"
				grid[position[0]][position[1]] = "X"
				position = []int{position[0] - 1, position[1]}
				numSteps++
			}
		} else if direction == ">" {
			if position[1]+1 > len(grid[0])-1 {
				grid[position[0]][position[1]] = "X"
				reachedEnd = true
			} else if grid[position[0]][position[1]+1] == "#" {
				direction = turnRight(direction)
				grid[position[0]][position[1]] = direction
			} else {
				grid[position[0]][position[1]+1] = ">"
				grid[position[0]][position[1]] = "X"
				position = []int{position[0], position[1] + 1}
				numSteps++
			}
		} else if direction == "v" {
			if position[0]+1 > len(grid)-1 {
				grid[position[0]][position[1]] = "X"
				reachedEnd = true
			} else if grid[position[0]+1][position[1]] == "#" {
				direction = turnRight(direction)
				grid[position[0]][position[1]] = direction
			} else {
				grid[position[0]+1][position[1]] = "v"
				grid[position[0]][position[1]] = "X"
				position = []int{position[0] + 1, position[1]}
				numSteps++
			}
		} else if direction == "<" {
			if position[1]-1 < 0 {
				grid[position[0]][position[1]] = "X"
				reachedEnd = true
			} else if grid[position[0]][position[1]-1] == "#" {
				direction = turnRight(direction)
				grid[position[0]][position[1]] = direction
			} else {
				grid[position[0]][position[1]-1] = "v"
				grid[position[0]][position[1]] = "X"
				position = []int{position[0], position[1] - 1}
				numSteps++
			}
		}
		if shouldPrint {
			printGrid(grid)
		}
	}

	return reachedEnd
}

func turnRight(direction string) string {
	switch direction {
	case "^":
		return ">"
	case ">":
		return "v"
	case "v":
		return "<"
	case "<":
		return "^"
	default:
		// Should never happen
		return ""
	}
}

func getCurrentPosition(grid [][]string) []int {
	for i, row := range grid {
		for j, cell := range row {
			if cell == "^" || cell == ">" || cell == "v" || cell == "<" {
				return []int{i, j}
			}
		}
	}

	return []int{0, 0}
}

func processInput(filename string) [][]string {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		message := fmt.Sprintf("Error reading file: %v", err)
		panic(message)
	}

	lines := strings.Split(string(bytes), "\n")
	grid := [][]string{}

	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	return grid
}

func printGrid(grid [][]string) {
	for _, row := range grid {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println()
	}
	fmt.Println()
}

func copyGrid(grid [][]string) [][]string {
	duplicate := make([][]string, len(grid))
	for i := range grid {
		duplicate[i] = make([]string, len(grid[i]))
		copy(duplicate[i], grid[i])
	}
	return duplicate
}
