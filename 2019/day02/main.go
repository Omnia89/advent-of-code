package main

import (
	"fmt"
	"slices"

	"advent2019/util"
)

func main() {
	data := util.GetRawData("day02")
	// data := util.GetRawTest("day02")

	list := util.StringToIntSlice(data, ",")

	part1(list)
	part2(list)
}

func part1(list []int) {
	counter := 0

	code := IntCode{
		0,
		slices.Clone(list),
	}
	code.program[1] = 12
	code.program[2] = 2

	// fmt.Printf("[%d]\t[%v]\n", code.index, code.program)

	counter = code.run(0)

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(list []int) {
	counter := 0

	for verb := range 100 {
		for noun := range 100 {
			code := IntCode{
				0,
				slices.Clone(list),
			}
			code.program[1] = verb
			code.program[2] = noun

			val := code.run(0)
			if val == 19690720 {
				counter = 100*verb + noun
				break
			}
		}
	}

	fmt.Printf("Part 2: %d\n", counter)
}
