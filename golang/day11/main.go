package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type StoneBlink struct {
	stone int
	blink int
}

var seen = make(map[StoneBlink]int)

func main() {
	filename := "input.txt"
	part1(filename, 25)
	part2(filename, 75)
}

func part1(filename string, maxBlinks int) {
	start := time.Now()

	stones := processInput(filename)

	for i := 0; i < maxBlinks; i++ {
		stones = blink(stones)
	}

	answer := len(stones)

	fmt.Printf("Part 1 (%d blinks): %d\tExecution time %v\n", maxBlinks, answer, time.Since(start))
}

func part2(filename string, maxBlinks int) {
	start := time.Now()
	answer := 0

	stones := processInput(filename)

	for _, stone := range stones {
		answer += blinkRecursive(stone, maxBlinks)
	}

	fmt.Printf("Part 2 (%d blinks): %d\tExecution time %v\n", maxBlinks, answer, time.Since(start))
}

func processInput(filename string) []int {
	bytes, _ := os.ReadFile(filename)
	stones := strings.Split(string(bytes), " ")
	intStones := []int{}

	for _, stone := range stones {
		intStone, _ := strconv.Atoi(stone)
		intStones = append(intStones, intStone)
	}

	return intStones
}

func blink(stones []int) []int {
	newStones := []int{}

	for _, stone := range stones {
		stoneStr := fmt.Sprintf("%d", stone)

		if stone == 0 {
			newStones = append(newStones, 1)
		} else if len(stoneStr)%2 == 0 {
			chars := strings.Split(stoneStr, "")
			firstNewStone, _ := strconv.Atoi(strings.Join(chars[0:len(stoneStr)/2], ""))
			secondNewStone, _ := strconv.Atoi(strings.Join(chars[len(stoneStr)/2:], ""))
			newStones = append(newStones, firstNewStone)
			newStones = append(newStones, secondNewStone)
		} else {
			newStone := stone * 2024
			newStones = append(newStones, newStone)
		}
	}

	return newStones
}

func blinkRecursive(stone int, currentBlink int) int {
	if currentBlink == 0 {
		return 1
	}

	if count, ok := seen[StoneBlink{stone: stone, blink: currentBlink}]; ok {
		return count
	}

	if stone == 0 {
		return blinkRecursive(1, currentBlink-1)
	} else if stoneStr := strconv.Itoa(stone); len(stoneStr)%2 == 0 {
		stone1, _ := strconv.Atoi(stoneStr[:len(stoneStr)/2])
		stone2, _ := strconv.Atoi(stoneStr[len(stoneStr)/2:])
		count := blinkRecursive(stone1, currentBlink-1) + blinkRecursive(stone2, currentBlink-1)
		seen[StoneBlink{stone: stone, blink: currentBlink}] = count
		return count
	} else {
		newStone := stone * 2024
		count := blinkRecursive(newStone, currentBlink-1)
		seen[StoneBlink{stone: stone, blink: currentBlink}] = count
		return count
	}
}
