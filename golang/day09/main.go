package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "input.txt"
	part1(filename)
	part2(filename)
}

func part1(filename string) {
	chars := processInput(filename)
	blocks := constructBlocks(chars)

	for i := 0; i < len(blocks); i++ {
		index := len(blocks) - 1 - i
		if blocks[index] == "." {
			continue
		}

		block := blocks[index]
		firstFreeBlock := -1

		for i, block := range blocks {
			if block == "." {
				firstFreeBlock = i
				break
			}
		}

		if firstFreeBlock > index {
			break
		}

		blocks[firstFreeBlock] = block
		blocks[index] = "."
	}

	answer := calcChecksum(blocks)
	fmt.Printf("Part 1: %d\n", answer)
}

func part2(filename string) {
	chars := processInput(filename)
	blocks := constructBlocks(chars)

	blockSize := map[string][]int{}

	for i, block := range blocks {
		if block == "." {
			continue
		}
		if blockSize[block] == nil {
			blockSize[block] = []int{i, 1}
		} else {
			blockSize[block] = []int{blockSize[block][0], blockSize[block][1] + 1}
		}
	}

	for i := 0; i < len(blockSize); i++ {
		block := strconv.Itoa(len(blockSize) - 1 - i)
		blockIndex := blockSize[block][0]
		size := blockSize[block][1]
		freeBlockIndex := -1
		freeBlockSize := 0

		for i := 0; i < len(blocks); i++ {
			if blocks[i] == "." {
				if freeBlockIndex == -1 {
					freeBlockIndex = i

					if freeBlockIndex > blockIndex {
						break
					}
				}

				freeBlockSize++

				if freeBlockSize == size {
					for j, b := range blocks {
						if b == block {
							blocks[j] = "."
						}
					}

					for i := 0; i < freeBlockSize; i++ {
						blocks[freeBlockIndex+i] = block
					}
					break
				}
			} else if blocks[i] == block {
				blockIndex = i
				freeBlockIndex = -1
				freeBlockSize = 0
			} else {
				freeBlockIndex = -1
				freeBlockSize = 0
			}
		}
	}

	answer := calcChecksum(blocks)
	fmt.Printf("Part 2: %d\n", answer)
}

func calcChecksum(blocks []string) int {
	answer := 0
	for i, block := range blocks {
		id, _ := strconv.Atoi(block)
		answer += i * id
	}
	return answer
}

func constructBlocks(chars []string) []string {
	id := 0
	blocks := []string{}

	for i, char := range chars {
		if i%2 == 0 {
			num, _ := strconv.Atoi(char)
			for i := 0; i < num; i++ {
				blocks = append(blocks, strconv.Itoa(id))
			}
			id++
		} else {
			num, _ := strconv.Atoi(char)
			for i := 0; i < num; i++ {
				blocks = append(blocks, ".")
			}
		}
	}

	return blocks
}

func processInput(filename string) []string {
	bytes, err := os.ReadFile(filename)

	if err != nil {
		message := fmt.Sprintf("Error reading file: %v", err)
		panic(message)
	}

	chars := strings.Split(string(bytes), "")

	return chars
}