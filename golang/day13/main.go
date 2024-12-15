package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Machine struct {
	Prize   [2]int
	ButtonA [2]int
	ButtonB [2]int
}

func main() {
	filename := "input.txt"
	part1(filename)
	part2(filename)
}

func part1(filename string) {
	answer := 0
	machines := processInput(filename, false)

	for _, machine := range machines {
		fmt.Printf("A: %v, B: %v, P: %v\n", machine.ButtonA, machine.ButtonB, machine.Prize)

		for aPresses := 0; aPresses < 100; aPresses++ {
			travelX := machine.ButtonA[0] * aPresses
			travelY := machine.ButtonA[1] * aPresses
			remainingX := machine.Prize[0] - travelX
			remainingY := machine.Prize[1] - travelY

			if remainingX%machine.ButtonB[0] == 0 && remainingY%machine.ButtonB[1] == 0 {
				bPressesForX := remainingX / machine.ButtonB[0]
				bPressesForY := remainingY / machine.ButtonB[1]

				fmt.Printf("Found a match %d, %d, %d\n", aPresses, bPressesForX, bPressesForY)

				if bPressesForX == bPressesForY {
					answer += (aPresses * 3) + bPressesForX
					break
				}
			}
		}
	}

	fmt.Printf("Part 1: %d\n", answer)
}

func part2(filename string) {
	answer := 0
	machines := processInput(filename, true)

	for _, machine := range machines {
		// Use Cramer's Rule since I can't brute force: https://en.wikipedia.org/wiki/Cramer%27s_rule
		determinant := (machine.ButtonA[0] * machine.ButtonB[1]) - (machine.ButtonA[1] * machine.ButtonB[0])
		if determinant == 0 {
			continue
		}

		numeratorA := (machine.Prize[0] * machine.ButtonB[1]) - (machine.Prize[1] * machine.ButtonB[0])
		numeratorB := (machine.ButtonA[0] * machine.Prize[1]) - (machine.ButtonA[1] * machine.Prize[0])

		if numeratorA%determinant != 0 || numeratorB%determinant != 0 {
			continue
		}

		aPresses := numeratorA / determinant
		bPresses := numeratorB / determinant

		if aPresses >= 0 && bPresses >= 0 {
			x := machine.ButtonA[0]*aPresses + machine.ButtonB[0]*bPresses
			y := machine.ButtonA[1]*aPresses + machine.ButtonB[1]*bPresses
			if x == machine.Prize[0] && y == machine.Prize[1] {
				answer += (aPresses * 3) + bPresses
			}
		}
	}

	fmt.Printf("Part 2: %d\n", answer)
}

func processInput(filename string, adjustValue bool) []Machine {
	bytes, _ := os.ReadFile(filename)
	file := string(bytes)
	groups := strings.Split(file, "\n\n")
	machines := []Machine{}

	r := regexp.MustCompile(`X\+(\d+)|Y\+(\d+)|X=(\d+)|Y=(\d+)`)
	for _, group := range groups {
		machine := Machine{}

		matches := r.FindAllStringSubmatch(group, -1)
		aX, _ := strconv.Atoi(matches[0][1])
		aY, _ := strconv.Atoi(matches[1][2])
		bX, _ := strconv.Atoi(matches[2][1])
		bY, _ := strconv.Atoi(matches[3][2])
		pX, _ := strconv.Atoi(matches[4][3])
		pY, _ := strconv.Atoi(matches[5][4])

		if adjustValue {
			pX += 10000000000000
			pY += 10000000000000
		}

		machine.ButtonA = [2]int{aX, aY}
		machine.ButtonB = [2]int{bX, bY}
		machine.Prize = [2]int{pX, pY}

		machines = append(machines, machine)
	}

	return machines
}
