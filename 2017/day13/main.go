package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"advent2017/util"
)

func main() {
	data := util.GetDataByRow("day13")
	// data := util.GetTestByRow("day13")

	list := parse(data)

	part1(list)
	part2(list)
}

func parse(data []string) map[int]int {
	layers := map[int]int{}

	for _, str := range data {
		s := strings.ReplaceAll(str, " ", "")
		k, v, _ := strings.Cut(s, ":")
		layers[util.ToInt(k)] = util.ToInt(v)
	}
	return layers
}

func part1(layers map[int]int) {
	counter := 0

	maxLayer := slices.Max(slices.Collect(maps.Keys(layers)))

	for i := range maxLayer + 1 {
		if i == 0 {
			continue
		}
		if l, ok := layers[i]; ok {
			step := l*2 - 2
			if i%step == 0 {
				counter += i * l
			}
		}
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(layers map[int]int) {
	counter := 0

	maxLayer := slices.Max(slices.Collect(maps.Keys(layers)))

	seconds := 1
	found := false

	for !found {
		found = true
		for i := range maxLayer + 1 {
			if l, ok := layers[i]; ok {
				step := l*2 - 2
				if (i+seconds)%step == 0 {
					found = false
					break
				}
			}
		}
		if !found {
			seconds++
		}
	}
	counter = seconds

	fmt.Printf("Part 2: %d\n", counter)
}
