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
	numHorizontal := countHorizontal(grid)
	numVertical := countVertical(grid)
	numLeftDiagonal := countLeftDiagonal(grid)
	numRightDiagonal := countRightDiagonal(grid)

	fmt.Printf("Horizontal: %d\n", numHorizontal)
	fmt.Printf("Vertical: %d\n", numVertical)
	fmt.Printf("LeftDiagonal: %d\n", numLeftDiagonal)
	fmt.Printf("RightDiagonal: %d\n", numRightDiagonal)

	fmt.Printf("Part 1: %d\n", numHorizontal+numVertical+numLeftDiagonal+numRightDiagonal)
}

func processInput(filename string) [][]string {
	bytes, err := os.ReadFile(filename)

	if err != nil {
		message := fmt.Sprintf("Error reading file: %s", filename)
		panic(message)
	}

	grid := [][]string{}

	lines := strings.Split(string(bytes), "\n")
	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	return grid
}

func countHorizontal(grid [][]string) int {
	return findMatches(grid)
	// fmt.Println("Normal Grid:")
	// printGrid(grid)
}

func countVertical(grid [][]string) int {
	transposedGrid := make([][]string, len(grid[0]))

	for i := range transposedGrid {
		transposedGrid[i] = make([]string, len(grid))
	}

	for i, row := range grid {
		for j, cell := range row {
			transposedGrid[j][i] = cell
		}
	}

	// fmt.Println("Columns:")
	// printGrid(transposedGrid)

	return findMatches(transposedGrid)
}

func countLeftDiagonal(grid [][]string) int {
	numRows := len(grid)
	numCols := len(grid[0])
	diagonals := make([][]string, numRows+numCols-1)

	for i, row := range grid {
		for j, cell := range row {
			diagonals[i+j] = append(diagonals[i+j], cell)
		}
	}

	// fmt.Println("Diagonals:")
	// printGrid(diagonals)

	return findMatches(diagonals)
}

func countRightDiagonal(grid [][]string) int {
	numRows := len(grid)
	numCols := len(grid[0])
	diagonals := make([][]string, numRows+numCols-1)

	for i, row := range grid {
		for j, cell := range row {
			diagonals[numCols-i+j-1] = append(diagonals[numCols-i+j-1], cell)
		}
	}

	// fmt.Println("Diagonals:")
	// printGrid(diagonals)

	return findMatches(diagonals)
}

func findMatches(grid [][]string) int {
	answer := 0

	for _, row := range grid {
		for i := range row {
			if i+3 >= len(row) {
				break
			}

			if row[i] == "X" && row[i+1] == "M" && row[i+2] == "A" && row[i+3] == "S" {
				answer++
			} else if row[i] == "S" && row[i+1] == "A" && row[i+2] == "M" && row[i+3] == "X" {
				answer++
			}
		}
	}

	return answer
}

// func printGrid(grid [][]string) {
// 	for _, row := range grid {
// 		fmt.Println(strings.Join(row, " "))
// 	}
// 	fmt.Println()
// }
