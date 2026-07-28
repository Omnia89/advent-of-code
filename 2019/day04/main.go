package main

import (
	"fmt"

	"advent2019/util"
)

func main() {
	data := util.GetRawData("day04")
	// data := util.GetRawTest("day04")

	parts := util.StringToIntSlice(data, "-")

	part1(parts)
	part2(parts)
}

func isValid(n int) bool {
	double := false

	parts := util.StringToIntSlice(fmt.Sprintf("%d", n), "")

	for i := range len(parts) - 1 {
		if !double {
			double = parts[i] == parts[i+1]
		}

		if parts[i] > parts[i+1] {
			return false
		}
	}

	return double
}

func part1(parts []int) {
	counter := 0

	lower := parts[0]
	upper := parts[1]

	for i := lower; i <= upper; i++ {
		if isValid(i) {
			counter++
		}
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func isValid2(n int) bool {
	double := false

	strN := fmt.Sprintf("%d", n)

	parts := util.StringToIntSlice(strN, "")

	same := 0

	for i := range len(parts) - 1 {
		if parts[i] != parts[i+1] {
			if same == 2 {
				double = true
			}
			same = 0
		} else {
			if same == 0 {
				same++
			}
			same++
		}

		if parts[i] > parts[i+1] {
			return false
		}
	}
	if same == 2 {
		double = true
	}

	return double
}

// 738 too low
func part2(parts []int) {
	counter := 0

	lower := parts[0]
	upper := parts[1]

	for i := lower; i <= upper; i++ {
		if isValid2(i) {
			counter++
		}
	}
	fmt.Printf("Part 2: %d\n", counter)
}
