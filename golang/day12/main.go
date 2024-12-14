package main

import (
	"fmt"
	"os"
	"strings"
)

type Region struct {
	Letter string
	Points [][2]int
}

var directions = [][]int{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

func main() {
	filename := "input.txt"
	part1(filename)
}

func part1(filename string) {
	grid := processInput(filename)
	regions := findRegions(grid)
	answer := 0

	for _, region := range regions {
		area := len(region.Points)
		perimeter := calculatePerimeter(grid, region)
		answer += area * perimeter
	}

	fmt.Printf("Part 1: %d\n", answer)
}

func calculatePerimeter(grid [][]string, region Region) int {
	perimeter := len(region.Points) * 4
	for _, point := range region.Points {
		for _, direction := range directions {
			newRow := point[0] + direction[0]
			newCol := point[1] + direction[1]

			if newRow >= 0 && newRow < len(grid) && newCol >= 0 && newCol < len(grid[0]) {
				if grid[newRow][newCol] == region.Letter {
					perimeter--
				}
			}
		}
	}
	return perimeter
}

func findRegions(grid [][]string) []Region {
	rows := len(grid)
	cols := len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	var regions []Region

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if !visited[r][c] { // If the cell is not visited
				region := depthFirstSearch(grid, visited, r, c, grid[r][c])
				regions = append(regions, region)
			}
		}
	}

	return regions
}

func depthFirstSearch(grid [][]string, visited [][]bool, row, col int, letter string) Region {
	visited[row][col] = true
	region := Region{Letter: letter, Points: [][2]int{{row, col}}}

	for _, direction := range directions {
		newRow := row + direction[0]
		newCol := col + direction[1]

		if newRow >= 0 && newRow < len(grid) && newCol >= 0 && newCol < len(grid[0]) && !visited[newRow][newCol] && grid[newRow][newCol] == letter {
			region.Points = append(region.Points, depthFirstSearch(grid, visited, newRow, newCol, letter).Points...)
		}
	}

	return region
}

func processInput(filename string) [][]string {
	bytes, _ := os.ReadFile(filename)
	lines := strings.Split(string(bytes), "\n")

	grid := [][]string{}

	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	return grid
}
