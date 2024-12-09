package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	filename := "input.txt"
	part1(filename)
	part2(filename)
}

func part1(filename string) {
	antennas, numRows, numColumns := processInput(filename)

	antinodes := [][]int{}
	seen := map[string]bool{}

	for _, frequency := range antennas {
		for i, v1 := range frequency {
			for j, v2 := range frequency {
				if i != j {
					xSlope := v2[0] - v1[0]
					ySlope := v2[1] - v1[1]
					antinode := []int{(v2[0] + xSlope), (v2[1] + ySlope)}
					if antinode[0] >= 0 && antinode[0] < numColumns && antinode[1] >= 0 && antinode[1] < numRows {
						key := fmt.Sprintf("%d,%d", antinode[0], antinode[1])
						if !seen[key] {
							seen[key] = true
							antinodes = append(antinodes, antinode)
						}
					}
				}
			}
		}
	}

	answer := len(antinodes)

	fmt.Printf("Part 1: %d\n", answer)
}

func part2(filename string) {
	processInput(filename)
	answer := 0
	fmt.Printf("Part 2: %d\n", answer)
}

func processInput(filename string) (antennas map[string][][]int, numRows int, numColumns int) {
	bytes, err := os.ReadFile(filename)

	if err != nil {
		message := fmt.Sprintf("Error reading file: %s", filename)
		panic(message)
	}

	lines := strings.Split(string(bytes), "\n")
	numRows = len(lines)
	numColumns = len(lines[0])
	antennas = map[string][][]int{}

	r := regexp.MustCompile(`\d|[a-z]|[A-Z]`)

	for i, row := range lines {
		columns := strings.Split(row, "")
		for j, column := range columns {
			if r.MatchString(column) {
				antennas[column] = append(antennas[column], []int{i, j})
			}
		}
	}

	return antennas, numRows, numColumns
}
