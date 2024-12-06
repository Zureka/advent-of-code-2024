package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := "input.txt"
	part1(filename)
}

func part1(filename string) {
	grid := processInput(filename)
	currentPosition := []int{0, 0}

	for i, row := range grid {
		for j, cell := range row {
			if cell == "^" {
				currentPosition = []int{i, j}
				break
			}
		}
	}

	processPath(currentPosition, "^", grid, false)

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

func processPath(startPosition []int, startDirection string, grid [][]string, shouldPrint bool) {
	position := startPosition
	direction := startDirection
	reachedEnd := false

	for !reachedEnd {
		if direction == "^" {
			if position[0]-1 < 0 {
				grid[position[0]][position[1]] = "X"
				reachedEnd = true
				printGrid(grid, shouldPrint)
			} else if grid[position[0]-1][position[1]] == "#" {
				direction = turnRight(direction)
				grid[position[0]][position[1]] = direction
				printGrid(grid, shouldPrint)
			} else {
				grid[position[0]-1][position[1]] = "^"
				grid[position[0]][position[1]] = "X"
				position = []int{position[0] - 1, position[1]}
				printGrid(grid, shouldPrint)
			}
		} else if direction == ">" {
			if position[1]+1 > len(grid[0])-1 {
				grid[position[0]][position[1]] = "X"
				reachedEnd = true
				printGrid(grid, shouldPrint)
			} else if grid[position[0]][position[1]+1] == "#" {
				direction = turnRight(direction)
				grid[position[0]][position[1]] = direction
				printGrid(grid, shouldPrint)
			} else {
				grid[position[0]][position[1]+1] = ">"
				grid[position[0]][position[1]] = "X"
				position = []int{position[0], position[1] + 1}
				printGrid(grid, shouldPrint)
			}
		} else if direction == "v" {
			if position[0]+1 > len(grid)-1 {
				grid[position[0]][position[1]] = "X"
				reachedEnd = true
				printGrid(grid, shouldPrint)
			} else if grid[position[0]+1][position[1]] == "#" {
				direction = turnRight(direction)
				grid[position[0]][position[1]] = direction
				printGrid(grid, shouldPrint)
			} else {
				grid[position[0]+1][position[1]] = "v"
				grid[position[0]][position[1]] = "X"
				position = []int{position[0] + 1, position[1]}
				printGrid(grid, shouldPrint)
			}
		} else if direction == "<" {
			if position[1]-1 < 0 {
				grid[position[0]][position[1]] = "X"
				reachedEnd = true
				printGrid(grid, shouldPrint)
			} else if grid[position[0]][position[1]-1] == "#" {
				direction = turnRight(direction)
				grid[position[0]][position[1]] = direction
				printGrid(grid, shouldPrint)
			} else {
				grid[position[0]][position[1]-1] = "v"
				grid[position[0]][position[1]] = "X"
				position = []int{position[0], position[1] - 1}
				printGrid(grid, shouldPrint)
			}
		}
	}
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

func printGrid(grid [][]string, shouldPrint bool) {
	if shouldPrint {
		for _, row := range grid {
			for _, cell := range row {
				fmt.Print(cell)
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
