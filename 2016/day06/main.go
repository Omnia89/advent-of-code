package main

import (
	"fmt"
	"math"

	"advent2016/util"
)

func main() {
	data := util.GetDataByRow("day06")
	// data := util.GetTestByRow("day06")

	part1(data)
	part2(data)
}

func part1(data []string) {
	chars := make([]map[string]int, len(data))

	for ii, s := range data {
		for i, c := range s {
			if ii == 0 {
				chars[i] = map[string]int{}
			}
			chars[i][string(c)]++
		}
	}

	message := ""

	for _, m := range chars {
		maxVal := 0
		c := ""
		for k, v := range m {
			if v > maxVal {
				maxVal = v
				c = k
			}
		}
		message += c
	}

	fmt.Printf("Part 1: %s\n", message)
}

func part2(data []string) {
	chars := make([]map[string]int, len(data))

	for ii, s := range data {
		for i, c := range s {
			if ii == 0 {
				chars[i] = map[string]int{}
			}
			chars[i][string(c)]++
		}
	}

	message := ""

	for _, m := range chars {
		minVal := math.MaxInt
		c := ""
		for k, v := range m {
			if v < minVal {
				minVal = v
				c = k
			}
		}
		message += c
	}

	fmt.Printf("Part 2: %s\n", message)
}
