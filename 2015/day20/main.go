package main

import (
	"fmt"

	"advent2015/util"
)

func main() {
	data := util.GetRawData("day20")
	// data := util.GetRawTest("day20")

	value := util.ToInt(data)

	part1(value)
	part2(value)
}

func part1(value int) {
	counter := 0

	targetSum := value / 10

	houses := make([]int, targetSum)

	for i := 1; i < targetSum; i++ {
		houses = append(houses, 0)
		for j := i; j < targetSum; j += i {
			houses[j] += i * 10
		}
	}

	for i, n := range houses {
		if n >= value {
			counter = i
			break
		}
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(value int) {
	counter := 0

	targetSum := value / 10

	houses := make([]int, targetSum)

	for i := 1; i < targetSum; i++ {
		houses = append(houses, 0)
		n := 0
		for j := i; j < targetSum; j += i {
			if n >= 50 {
				break
			}
			houses[j] += i * 11
			n++
		}
	}

	for i, n := range houses {
		if n >= value {
			counter = i
			break
		}
	}
	fmt.Printf("Part 2: %d\n", counter)
}
