package main

import (
	"fmt"

	"advent2019/util"
)

func main() {
	data := util.GetRawData("day01")
	// data := util.GetRawTest("day01")

	list := util.StringToIntSlice(data, "\n")

	part1(list)
	part2(list)
}

func calcFuel(n int) int {
	return n/3 - 2
}

func part1(data []int) {
	counter := 0

	for _, n := range data {
		counter += calcFuel(n)
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data []int) {
	counter := 0

	for _, n := range data {
		t := n
		for t > 0 {
			t = util.IntMax(calcFuel(t), 0)
			counter += t
		}
	}

	fmt.Printf("Part 2: %d\n", counter)
}
