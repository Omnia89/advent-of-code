package main

import (
	"advent2025/util"
	"fmt"
)

func main() {

	data := util.GetDataByRow("day03")
	// data := util.GetTestByRow("day03")


	part1(data)
	part2(data)
}

func part1(data []string) {
	counter := 0

	for _, r := range data {
		if r == "" {
			continue
		}

		num := 0
		for i := 0; i < len(r) - 1; i++ {
			for j := i + 1; j < len(r); j++ {
				n := util.ToInt(fmt.Sprintf("%s%s", string(r[i]), string(r[j])))
				if n > num {
					num = n
				}
			}
		}
		counter += num
	}


	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data []string) {
	counter := 0

	for _, r := range data {
		if r == "" {
			continue
		}

		num := ""
		for len(num) < 12 {
			greatestIndex := 0
			// sliding window: I get the higher number per section
			for i := 0; i < len(r) - 11 + len(num); i++ {
				if r[i] > r[greatestIndex] {
					greatestIndex = i
				}
			}
			num += string(r[greatestIndex])
			r = r[greatestIndex+1:]
		}
		counter += util.ToInt(num)
	}


	fmt.Printf("Part 2: %d\n", counter)
}
