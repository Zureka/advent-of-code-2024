package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Robot struct {
	Position [2]int
	Velocity [2]int
}

type Variant struct {
	filename string
	width    int
	height   int
}

func main() {
	// variant := Variant{"test.txt", 11, 7}
	variant := Variant{"input.txt", 101, 103}
	part1(variant)
	part2(variant)
}

func part1(variant Variant) {
	start := time.Now()
	numSeconds := 100
	answer := 0
	robots := processInput(variant.filename)

	for i := 0; i < numSeconds; i++ {
		for j := 0; j < len(robots); j++ {
			applyVelocity(&robots[j], variant.width, variant.height)
		}
	}

	q1, q2, q3, q4 := countQuadrants(robots, variant.width, variant.height)
	answer = q1 * q2 * q3 * q4

	fmt.Printf("Part 1: %d, Execution time: %v\n", answer, time.Since(start))
}

func part2(variant Variant) {
	start := time.Now()
	maxSeconds := 10000
	answer := 0
	robots := processInput(variant.filename)

	for i := 0; i < maxSeconds; i++ {
		for j := 0; j < len(robots); j++ {
			applyVelocity(&robots[j], variant.width, variant.height)
		}
		found := findEasterEgg(robots, variant.width, variant.height)
		if found {
			answer = i + 1
			break
		}
	}

	fmt.Printf("Part 2: %d, Execution time: %v\n", answer, time.Since(start))
}

func countQuadrants(robots []Robot, width int, height int) (int, int, int, int) {
	quadrants := [4]int{0, 0, 0, 0}

	for _, robot := range robots {
		if robot.Position[0] == width/2 || robot.Position[1] == height/2 {
			continue
		}
		if robot.Position[0] < width/2 && robot.Position[1] < height/2 {
			quadrants[0]++
		} else if robot.Position[0] >= width/2 && robot.Position[1] < height/2 {
			quadrants[1]++
		} else if robot.Position[0] < width/2 && robot.Position[1] >= height/2 {
			quadrants[2]++
		} else {
			quadrants[3]++
		}
	}

	fmt.Println(quadrants)

	return quadrants[0], quadrants[1], quadrants[2], quadrants[3]
}

func applyVelocity(robot *Robot, width int, height int) {
	robot.Position[0] += robot.Velocity[0]
	robot.Position[1] += robot.Velocity[1]

	if robot.Position[0] < 0 {
		robot.Position[0] = width + robot.Position[0]
	}

	if robot.Position[0] >= width {
		robot.Position[0] = robot.Position[0] - width
	}

	if robot.Position[1] < 0 {
		robot.Position[1] = height + robot.Position[1]
	}

	if robot.Position[1] >= height {
		robot.Position[1] = robot.Position[1] - height
	}
}

func processInput(filename string) []Robot {
	bytes, _ := os.ReadFile(filename)
	lines := strings.Split(string(bytes), "\n")
	robots := []Robot{}

	r := regexp.MustCompile(`^p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)$`)

	for _, line := range lines {
		robot := Robot{}
		matches := r.FindAllStringSubmatch(line, -1)

		pX, _ := strconv.Atoi(matches[0][1])
		pY, _ := strconv.Atoi(matches[0][2])
		vX, _ := strconv.Atoi(matches[0][3])
		vY, _ := strconv.Atoi(matches[0][4])

		robot.Position = [2]int{pX, pY}
		robot.Velocity = [2]int{vX, vY}

		robots = append(robots, robot)
	}

	return robots
}

func findEasterEgg(robots []Robot, width int, height int) bool {
	grid := make([][]string, height)
	content := ""
	found := false
	r := regexp.MustCompile(`\#{8,}`)

	for i := 0; i < height; i++ {
		grid[i] = make([]string, width)
		for j := 0; j < width; j++ {
			grid[i][j] = "."
		}
	}

	for _, robot := range robots {
		grid[robot.Position[1]][robot.Position[0]] = "#"
	}

	for i := 0; i < height; i++ {
		row := strings.Join(grid[i], "")
		if r.MatchString(row) {
			found = true
		}
		content += row + "\n"
	}

	// if found {
	// 	content = strings.ReplaceAll(content, ".", "🟥")
	// 	content = strings.ReplaceAll(content, "#", "🤖")
	// 	fmt.Println(content)
	// }

	return found
}
