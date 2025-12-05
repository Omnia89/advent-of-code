package main

import (
	"fmt"
	"strings"

	"advent2025/util"
)

type Range struct {
	min int
	max int
}

func main() {
	data := util.GetDataByRow("day05")
	// data := util.GetTestByRow("day05")

	ranges := make([]Range, 0)
	values := make([]int, 0)

	change := false

	for _, r := range data {
		if r == "" {
			change = true
			continue
		}
		if change {
			values = append(values, util.ToInt(r))
		} else {
			val := strings.Split(r, "-")
			ranges = append(ranges, Range{
				min: util.ToInt(val[0]),
				max: util.ToInt(val[1]),
			})
		}
	}

	part1(ranges, values)
	part2(ranges)
}

func part1(ranges []Range, values []int) {
	counter := 0

	for _, n := range values {
		for _, r := range ranges {
			if n >= r.min && n <= r.max {
				counter += 1
				break
			}
		}
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(ranges []Range) {
	uniqueRanges := make([]Range, 0)

	source := ranges

	mergiati := true
	for mergiati {
		mergiati = false
		uniqueRanges = make([]Range, 0)
		for _, r := range source {
			if len(uniqueRanges) == 0 {
				// fmt.Printf("inserisco r[%d ~ %d]\n", r.min, r.max)
				uniqueRanges = append(uniqueRanges, r)
				continue
			}
			done := false
			for i, u := range uniqueRanges {
				if r.min <= u.max && r.max >= u.min {
					// fmt.Printf("mergio r[%d ~ %d] u[%d ~ %d] [%d ~ %d]\n", r.min, r.max, u.min, u.max, min(u.min, r.min), max(u.max, r.max))
					uniqueRanges[i].min = min(u.min, r.min)
					uniqueRanges[i].max = max(u.max, r.max)
					done = true
					mergiati = true
					break
				}
			}
			if !done {
				// fmt.Printf("inserisco r[%d ~ %d]\n", r.min, r.max)
				uniqueRanges = append(uniqueRanges, r)
			}
		}
		source = uniqueRanges
	}

	counter := 0

	for _, uu := range uniqueRanges {
		// fmt.Printf("[%d ~ %d]\n", uu.min, uu.max)
		counter += uu.max - uu.min + 1
	}

	fmt.Printf("Part 2: %d\n", counter)
}
