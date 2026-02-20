package main

import (
	"fmt"
	"strings"

	"advent2015/util"
)

func main() {
	data := util.GetDataByRow("day01")
	// data := util.GetTestByRow("day01")

	part1(data)
	part2(data)
}

func part1(data []string) {
	counter := 0

	line := data[0]

	openLen := len(strings.ReplaceAll(line, ")", ""))
	closeLen := len(strings.ReplaceAll(line, "(", ""))

	counter = openLen - closeLen

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data []string) {
	counter := 0

	line := data[0]

	floor := 0

	for i, c := range line {
		if c == '(' {
			floor++
		} else {
			floor--
		}
		if floor == -1 {
			counter = i + 1
			break
		}
	}

	fmt.Printf("Part 2: %d\n", counter)
}
