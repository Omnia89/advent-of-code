package main

import (
	"fmt"
	"regexp"

	"advent2015/util"
)

func main() {
	data := util.GetDataByRow("day25")
	// data := util.GetTestByRow("day25")

	row, col := parse(data[0])

	part1(row, col)
	part2(row, col)
}

func parse(data string) (row int, col int) {
	numberRegex := regexp.MustCompile(`(\d+)`)
	numbers := numberRegex.FindAllString(data, -1)

	row = util.ToInt(numbers[0])
	col = util.ToInt(numbers[1])
	return
}

func getNext(n int) int {
	return (n * 252533) % 33554393
}

// 13041972 too high
func part1(row, col int) {
	counter := 0

	numberIdx := ((row + col - 2) * (row + col - 1) / 2) + col

	// 0 based
	numberIdx--

	counter = 20151125

	for range numberIdx {
		counter = getNext(counter)
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(row, col int) {
	counter := 0

	fmt.Printf("Part 2: %d\n", counter)
}
