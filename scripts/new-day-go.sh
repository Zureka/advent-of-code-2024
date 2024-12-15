if [ -z "$1" ]; then
    echo "Usage: new-day-go.sh <day>"
    exit 1
fi

mkdir golang/day$1
cd golang/day$1
go mod init day$1
touch main.go test.txt input.txt

FILE="
package main

import(
  \"fmt\"
)

func main() {
  filename := \"test.txt\"
  part1(filename)
  part2(filename)
}

func part1(filename string) {
  answer := 0
  fmt.Printf(\"Part 1: %d\", answer)
}

func part2(filename string) {
  answer := 0
  fmt.Printf(\"Part 2: %d\", answer)
}"

echo "$FILE" > main.go
