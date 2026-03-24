package main

import (
	"fmt"

	"advent2017/util"
)

func main() {
	data := util.GetDataByRow("day01")
	// data := util.GetTestByRow("day01")

	list := util.StringToIntSlice(data[0], "")

	part1(list)
	part2(list)
}

func part1(data []int) {
	counter := 0

	for i := range data {
		n := data[i]
		m := data[(i+1)%len(data)]

		if n == m {
			counter += n
		}
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data []int) {
	counter := 0

	half := len(data) / 2

	for i := range data {
		n := data[i]
		m := data[(i+half)%len(data)]

		if n == m {
			counter += n
		}
	}
	fmt.Printf("Part 2: %d\n", counter)
}
