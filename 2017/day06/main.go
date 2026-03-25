package main

import (
	"fmt"
	"slices"

	"advent2017/util"
)

func main() {
	data := util.GetRawData("day06")
	// data := util.GetRawTest("day06")

	list := util.StringToIntSlice(data, "\t")
	list2 := slices.Clone(list)

	part1(list)
	part2(list2)
}

func getBiggerBankIndex(data []int) int {
	big := 0
	idx := -1

	for i, v := range data {
		if v > big {
			idx = i
			big = v
		}
	}
	return idx
}

func part1(data []int) {
	counter := 0

	banks := len(data)
	patterns := map[string]bool{}

	patterns[util.IntJoin(data, "-")] = true

	for {
		idx := getBiggerBankIndex(data)
		v := data[idx]
		data[idx] = 0
		for range v {
			idx = (idx + 1) % banks
			data[idx]++
		}
		counter++

		k := util.IntJoin(data, "-")
		if _, ok := patterns[k]; ok {
			break
		}
		patterns[k] = true
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data []int) {
	counter := 0

	banks := len(data)
	patterns := map[string]int{}

	patterns[util.IntJoin(data, "-")] = 0

	for {
		idx := getBiggerBankIndex(data)
		v := data[idx]
		data[idx] = 0
		for range v {
			idx = (idx + 1) % banks
			data[idx]++
		}
		counter++

		k := util.IntJoin(data, "-")
		if c, ok := patterns[k]; ok {
			counter -= c
			break
		}
		patterns[k] = counter
	}
	fmt.Printf("Part 2: %d\n", counter)
}
