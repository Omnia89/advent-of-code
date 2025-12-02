package main

import (
	"advent2025/util"
	"fmt"
	"strings"
)

func main() {

	data := util.GetDataByRow("day01")
	// data := util.GetTestByRow("day01")

	part1(data)
	part2(data)
}

func part1(data []string) {
	counter := 0
	index := 50

	for _, s := range data {
		if s == "" {
			continue
		}
		number := util.ToInt(s[1:])
		if strings.HasPrefix(s, "L") {
			number *= -1
		}

		index = (10000 + index + number) % 100

		if index == 0 {
			counter += 1
		}

	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data []string) {
	counter := 0
	index := 50

	for _, s := range data {
		if s == "" {
			continue
		}
		number := util.ToInt(s[1:])
		if strings.HasPrefix(s, "L") {
			number *= -1
		}

		temp := index + number
		
		direction := -100
		if temp < 0 {
			direction = 100
		}

		// Specific case: if i'm at 0 and go left, I don't have to count the passage, i'm already at 0 and counted previously
		if direction == 100 && index == 0 {
			// I'm countering the plus 1
			counter -= 1
		}

		for temp < 0 || temp > 100 {
			counter += 1
			temp += direction
		}
		index = temp % 100

		if index == 0 {
			counter += 1
		}

	}

	fmt.Printf("Part 2: %d\n", counter)
}
