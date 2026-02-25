package main

import (
	"fmt"
	"math"
	"slices"

	"advent2015/util"
)

func main() {
	data := util.GetRawData("day17")
	// data := util.GetRawTest("day17")

	list := util.StringToIntSlice(data, "\n")

	part1(list)
	part2(list)
}

func part1(data []int) {
	counter := 0

	slices.Sort(data)

	objective := 150

	jumps := make([]int, objective+1)

	// one way to get 0 => not choose anything
	jumps[0] = 1

	for _, num := range data {
		for i := objective; i >= num; i-- {
			jumps[i] += jumps[i-num]
		}
	}

	counter = jumps[objective]

	fmt.Printf("Part 1: %d\n", counter)
}

// not 4
func part2(data []int) {
	counter := 0

	slices.Sort(data)

	objective := 150

	jumps := make([][][]int, objective+1)

	// one way to get 0 => not choose anything
	jumps[0] = [][]int{{}}

	for _, num := range data {
		for i := objective; i >= num; i-- {
			js := jumps[i-num]
			for _, j := range js {
				jj := append(j, num)
				jumps[i] = append(jumps[i], jj)
			}
			// jumps[i] = append(jumps[i], jumps[i-num]...)
		}
	}
	groups := map[int]int{}

	minLen := math.MaxInt
	for _, l := range jumps[objective] {
		groups[len(l)]++
		if len(l) < minLen {
			minLen = len(l)
		}
	}

	counter = groups[minLen]
	fmt.Printf("Part 2: %d\n", counter)
}
