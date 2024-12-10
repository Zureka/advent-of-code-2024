package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	filename := "input.txt"
	part1(filename)
	part2(filename)
}

func part1(filename string) {
	start := time.Now()
	grid, trailheads := processInput(filename)
	answer := processTrails(trailheads, grid, true)

	fmt.Printf("Part 1: %d\tExecution time: %v\n", answer, time.Since(start))
}

func part2(filename string) {
	start := time.Now()
	grid, trailheads := processInput(filename)
	answer := processTrails(trailheads, grid, false)

	fmt.Printf("Part 2: %d\tExecution time: %v\n", answer, time.Since(start))
}

func processTrails(trailheads [][]int, grid [][]int, distinct bool) int {
	width := len(grid[0])
	height := len(grid)
	answer := 0

	resCh := make(chan int)
	var wg sync.WaitGroup

	for _, trailhead := range trailheads {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var seen map[string]bool
			var score int

			if distinct {
				seen = map[string]bool{}
			} else {
				score = 0
			}

			queue := [][]int{trailhead}
			for len(queue) > 0 {
				point := queue[0]
				queue = queue[1:]
				currentElevation := grid[point[0]][point[1]]
				nextElevation := currentElevation + 1

				if currentElevation == 9 {
					if distinct {
						key := fmt.Sprintf("%v", point)
						if !seen[key] {
							seen[key] = true
						}
					} else {
						score++
					}
					continue
				}

				if point[1] > 0 && grid[point[0]][point[1]-1] == nextElevation {
					queue = append(queue, []int{point[0], point[1] - 1})
				}
				if point[1] < width-1 && grid[point[0]][point[1]+1] == nextElevation {
					queue = append(queue, []int{point[0], point[1] + 1})
				}
				if point[0] > 0 && grid[point[0]-1][point[1]] == nextElevation {
					queue = append(queue, []int{point[0] - 1, point[1]})
				}
				if point[0] < height-1 && grid[point[0]+1][point[1]] == nextElevation {
					queue = append(queue, []int{point[0] + 1, point[1]})
				}
			}

			if distinct {
				resCh <- len(seen)
			} else {
				resCh <- score
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

	return answer
}

func processInput(filename string) ([][]int, [][]int) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		message := fmt.Sprintf("Error reading file: %s\n", filename)
		panic(message)
	}

	grid := [][]int{}

	lines := strings.Split(string(bytes), "\n")

	trailheads := [][]int{}

	for i, line := range lines {
		cells := strings.Split(line, "")
		cellInts := []int{}
		for _, cell := range cells {
			num, _ := strconv.Atoi(cell)
			cellInts = append(cellInts, num)
		}

		grid = append(grid, cellInts)

		for j, cell := range cells {
			if cell == "0" {
				trailheads = append(trailheads, []int{i, j})
			}
		}
	}

	return grid, trailheads
}
