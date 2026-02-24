package main

import (
	"fmt"
	"slices"
	"strings"

	"advent2015/util"
)

func main() {
	data := util.GetDataByRow("day16")
	// data := util.GetTestByRow("day16")

	list := parse(data)

	part1(list)
	part2(list)
}

type aunt struct {
	index int
	// children    *int
	// cats        *int
	// samoyeds    *int
	// pomeranians *int
	// akitas      *int
	// vizlas      *int
	// goldfish    *int
	// trees       *int
	// cars        *int
	// perfumes    *int
	values map[string]int
}

func parse(data []string) []aunt {
	aunts := []aunt{}

	for _, s := range data {
		normalize := strings.ReplaceAll(s, ",", "")
		normalize = strings.ReplaceAll(normalize, ":", "")

		pieces := strings.Split(normalize, " ")

		a := aunt{index: util.ToInt(pieces[1])}

		values := map[string]int{}
		for i := 2; i < len(pieces); i += 2 {
			values[pieces[i]] = util.ToInt(pieces[i+1])
		}
		a.values = values
		aunts = append(aunts, a)
	}

	return aunts
}

func compare(scan map[string]int, memory map[string]int) (int, bool) {
	n := 0

	for k, v := range memory {
		if scan[k] != v {
			return 0, false
		}
		n++
	}
	return n, true
}

func compareRetroencabulator(scan map[string]int, memory map[string]int) (int, bool) {
	n := 0

	greater := []string{"cats", "trees"}
	lesser := []string{"pomeranians", "goldfish"}

	for k, v := range memory {
		if slices.Contains(greater, k) {
			if scan[k] >= v {
				return 0, false
			}
			n++
			continue
		}
		if slices.Contains(lesser, k) {
			if scan[k] <= v {
				return 0, false
			}
			n++
			continue
		}

		if scan[k] != v {
			return 0, false
		}
		n++
	}
	return n, true
}

func part1(data []aunt) {
	counter := 0

	scan := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizlas":      0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}

	maxScore := 0
	auntIndex := 0
	for _, a := range data {
		score, ok := compare(scan, a.values)
		if ok && score > maxScore {
			maxScore = score
			auntIndex = a.index
		}
	}

	counter = auntIndex

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data []aunt) {
	counter := 0

	scan := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizlas":      0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}

	maxScore := 0
	auntIndex := 0
	for _, a := range data {
		score, ok := compareRetroencabulator(scan, a.values)
		if ok && score > maxScore {
			maxScore = score
			auntIndex = a.index
		}
	}

	counter = auntIndex
	fmt.Printf("Part 2: %d\n", counter)
}
