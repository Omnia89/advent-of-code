package main

import (
	"fmt"

	"advent2016/util"
)

func main() {
	data := util.GetDataByRow("day19")
	// data := util.GetTestByRow("day19")

	part1(util.ToInt(data[0]))
	part2(util.ToInt(data[0]))
}

func buildArrayIdx(n int) []int {
	r := make([]int, 0, n)

	for i := range n {
		r = append(r, i+1)
	}
	return r
}

func part1(n int) {
	counter := 0

	indexes := buildArrayIdx(n)

	for len(indexes) > 1 {
		start := 0
		if len(indexes)%2 != 0 {
			start = 2
		}
		temp := []int{}
		for i := start; i < len(indexes); i += 2 {
			temp = append(temp, indexes[i])
		}
		indexes = temp
	}
	counter = indexes[0]

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(n int) {
	counter := 0

	p := 1

	// find power of 3 less than n
	for p*3 <= n {
		p *= 3
	}

	if n == p {
		// if whole power of 3, the last elf is correct
		counter = n
	} else if n <= 2*p {
		// if p is over the half of n
		counter = n - p
	} else {
		// get the 2/3 position
		counter = 2*n - 3*p
	}

	fmt.Printf("Part 2: %d\n", counter)
}
