package main

import (
	"fmt"
	"slices"
	"strings"

	"advent2017/util"
)

func main() {
	data := util.GetDataByRow("day04")
	// data := util.GetTestByRow("day04")

	part1(data)
	part2(data)
}

func part1(data []string) {
	counter := 0

	for _, s := range data {
		values := strings.Split(s, " ")

		passw := map[string]bool{}
		double := false

		for _, v := range values {
			if _, ok := passw[v]; ok {
				double = true
			}
			passw[v] = true
		}
		if !double {
			counter++
		}
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func orderString(s string) string {
	o := strings.Split(s, "")
	slices.Sort(o)
	return strings.Join(o, "")
}

func part2(data []string) {
	counter := 0

	for _, s := range data {
		values := strings.Split(s, " ")

		passw := map[string]bool{}
		double := false

		for _, v := range values {
			anag := orderString(v)
			if _, ok := passw[anag]; ok {
				double = true
			}
			passw[anag] = true
		}
		if !double {
			counter++
		}
	}
	fmt.Printf("Part 2: %d\n", counter)
}
