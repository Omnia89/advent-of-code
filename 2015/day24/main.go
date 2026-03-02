package main

import (
	"fmt"
	"math"
	"slices"

	"advent2015/util"
)

func main() {
	data := util.GetRawData("day24")
	// data := util.GetRawTest("day24")

	list := util.StringToIntSlice(data, "\n")

	part1(list)
	part2(list)
}

func arraySum(data []int) int {
	s := 0
	for _, i := range data {
		s += i
	}
	return s
}

func arrayMult(data []int) int {
	m := 1
	for _, i := range data {
		m *= i
	}
	return m
}

func part1(data []int) {
	counter := 0

	slices.Sort(data)

	tot := arraySum(data)
	if tot%3 != 0 {
		panic("Not valid input")
	}

	target := tot / 3

	combinations := make([][][]int, target+1)

	combinations[0] = [][]int{{}}

	for _, n := range data {
		for i := target; i >= n; i-- {

			c := combinations[i-n]
			for _, l := range c {
				ll := append(l, n)
				combinations[i] = append(combinations[i], ll)
			}
		}
	}

	groups := map[int][][]int{}
	minLen := math.MaxInt

	for _, l := range combinations[target] {
		ll := len(l)
		if _, ok := groups[ll]; !ok {
			groups[ll] = [][]int{}
		}
		groups[ll] = append(groups[ll], l)

		if ll < minLen {
			minLen = ll
		}
	}

	quantum := math.MaxInt

	for _, l := range groups[minLen] {
		q := arrayMult(l)
		if q < quantum {
			quantum = q
		}
	}

	counter = quantum

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data []int) {
	counter := 0

	slices.Sort(data)

	tot := arraySum(data)
	if tot%4 != 0 {
		panic("Not valid input")
	}

	target := tot / 4

	combinations := make([][][]int, target+1)

	combinations[0] = [][]int{{}}

	for _, n := range data {
		for i := target; i >= n; i-- {

			c := combinations[i-n]
			for _, l := range c {
				ll := append(l, n)
				combinations[i] = append(combinations[i], ll)
			}
		}
	}

	groups := map[int][][]int{}
	minLen := math.MaxInt

	for _, l := range combinations[target] {
		ll := len(l)
		if _, ok := groups[ll]; !ok {
			groups[ll] = [][]int{}
		}
		groups[ll] = append(groups[ll], l)

		if ll < minLen {
			minLen = ll
		}
	}

	quantum := math.MaxInt

	for _, l := range groups[minLen] {
		q := arrayMult(l)
		if q < quantum {
			quantum = q
		}
	}

	counter = quantum
	fmt.Printf("Part 2: %d\n", counter)
}
