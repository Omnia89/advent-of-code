package main

import (
	"fmt"

	"advent2017/util"
)

func main() {
	data := util.GetRawData("day17")
	// data := util.GetRawTest("day17")

	val := util.ToInt(data)

	part1(val)
	part2(val)
}

func insertInArray(vet []int, pos int, val int) []int {
	ret := make([]int, 0, len(vet)+1)
	for i, n := range vet {
		if i == pos {
			ret = append(ret, val)
		}
		ret = append(ret, n)
	}
	if len(vet) == pos {
		ret = append(ret, val)
	}

	return ret
}

func part1(data int) {
	counter := 0

	pos := 0

	vet := []int{0}

	for i := range 2017 {
		n := i + 1

		pos = (pos+data)%len(vet) + 1
		vet = insertInArray(vet, pos, n)
	}

	counter = vet[(pos+1)%len(vet)]

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data int) {
	counter := 0

	// The 0 will always be in first position, so the final result will always be on index `1`.
	// Therefore, it's not needed to keep the whole array, just keep the last result of pos == 1

	pos := 0
	currentLength := 1

	for i := range 50_000_000 {
		pos = (pos+data)%currentLength + 1
		if pos == 1 {
			counter = i + 1
		}
		currentLength++
	}

	fmt.Printf("Part 2: %d\n", counter)
}
