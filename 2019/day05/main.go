package main

import (
	"fmt"
	"slices"

	"advent2019/util"
)

func main() {
	data := util.GetRawData("day05")
	//data := util.GetRawTest("day05")

	list := util.StringToIntSlice(data, ",")

	part1(list)
	part2(list)
}

func part1(list []int) {
	counter := 0

	code := IntCode{
		0,
		slices.Clone(list),
		[]int{1},
	}
	counter = code.run(0)

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(list []int) {
	counter := 0

	code := IntCode{
		0,
		slices.Clone(list),
		[]int{5},
	}
	counter = code.run(0)

	fmt.Printf("Part 2: %d\n", counter)
}
