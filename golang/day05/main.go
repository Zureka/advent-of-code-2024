package main

import (
  "fmt"
  "os"
  "slices"
  "strconv"
  "strings"
)

func main() {
  filename := "input.txt"
  part1(filename)
  part2(filename)
}

func part1(filename string) {
  updates, orderings := processInput(filename)
  answer := 0

  for _, update := range updates {
    valid := true

    for _, ordering := range orderings {
      firstIndex := slices.Index(update, ordering[0])
      secondIndex := slices.Index(update, ordering[1])

      if firstIndex != -1 && secondIndex != -1 && firstIndex > secondIndex {
        valid = false
      }
    }

    if valid {
      middleValue := update[((len(update)-1)/2)]
      num, err := strconv.Atoi(middleValue)
      if err == nil {
        answer += num
      }
    }
  }

  fmt.Printf("Part 1: %d\n", answer)
}

func part2(filename string) {
  updates, orderings := processInput(filename)
  answer := 0

  invalidUpdates := [][]string{}

  // Filter out initial "valid" updates and start list of invalid updates that need to be checked again
  for _, update := range updates {
    for _, ordering := range orderings {
      firstIndex := slices.Index(update, ordering[0])
      secondIndex := slices.Index(update, ordering[1])

      if firstIndex != -1 && secondIndex != -1 && firstIndex > secondIndex {
        firstValue := update[firstIndex]
        update[firstIndex] = update[secondIndex]
        update[secondIndex] = firstValue
        invalidUpdates = append(invalidUpdates, update)
        break
      }
    }
  }

  numUpdatesCorrected := 0

  for numUpdatesCorrected < len(invalidUpdates) {
    numUpdatesCorrected = 0

    for _, update := range invalidUpdates {
      valid := true

      for _, ordering := range orderings {
        firstIndex := slices.Index(update, ordering[0])
        secondIndex := slices.Index(update, ordering[1])

        if firstIndex != -1 && secondIndex != -1 && firstIndex > secondIndex {
          valid = false
          firstValue := update[firstIndex]
          update[firstIndex] = update[secondIndex]
          update[secondIndex] = firstValue
        }
      }

      if valid {
        numUpdatesCorrected++
      }
    }
  }

  for _, update := range invalidUpdates {
      middleValue := update[((len(update)-1)/2)]
      num, err := strconv.Atoi(middleValue)
      if err == nil {
        answer += num
      }
  }

  fmt.Printf("Part 2: %d\n", answer)
}

func getMiddleValue(update []string) int {
  middleValue := update[((len(update)-1)/2)]
  num, err := strconv.Atoi(middleValue)
  if err == nil {
    return num
  }
  return 0
}

func processInput(filename string) ([][]string, [][]string) {
  bytes, err := os.ReadFile(filename)
  if err != nil {
    message := fmt.Sprintf("Error reading file: %s", filename)
    panic(message)
  }

  lines := strings.Split(string(bytes), "\n")
  orderings := [][]string{}
  updates := [][]string{}

  for _, line := range lines {
    if line == "" {
      continue
    } else if strings.Contains(line, "|") {
      ordering := strings.Split(line, "|")
      orderings = append(orderings, ordering)
    } else {
      pages := strings.Split(line, ",")
      updates = append(updates, pages)
    }
  }

  return updates, orderings
}
