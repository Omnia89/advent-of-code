package main

import (
	"fmt"

	"advent2019/util"
)

func main() {
	data := util.GetRawData("day09")
	// data := util.GetRawTest("day09")

	list := util.StringToIntSlice(data, ",")

	part1(list)
	part2(list)
}

func part1(data []int) {
	counter := 0

	code := NewIntcodeArray(data, []int{1}, true)

	code.run(0)

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data []int) {
	counter := 0
	code := NewIntcodeArray(data, []int{2}, true)

	code.run(0)

	fmt.Printf("Part 2: %d\n", counter)
}
