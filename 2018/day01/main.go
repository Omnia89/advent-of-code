package main

import (
	"fmt"

	"advent2018/util"
)

func main() {
	data := util.GetRawData("day01")
	// data := util.GetRawTest("day01")

	list := util.StringToIntSlice(data, "\n")

	part1(list)
	part2(list)
}

func part1(data []int) {
	counter := 0

	for _, n := range data {
		counter += n
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data []int) {
	counter := 0

	alreadySeen := map[int]bool{0: true}
	s := 0

externalLoop:
	for {
		for _, n := range data {
			s += n
			if alreadySeen[s] {
				counter = s
				break externalLoop
			}
			alreadySeen[s] = true
		}
	}

	fmt.Printf("Part 2: %d\n", counter)
}
