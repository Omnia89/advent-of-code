package main

import (
	"fmt"
	"slices"
	"strconv"

	"advent2024/util"
)

func main() {
	// data := util.GetRawTest("day11")
	data := util.GetRawData("day11")

	list := util.StringToIntSlice(data, " ")

	part1(list)
	part2(list)
}

func compute(stone int, blinks int, memory map[string]int) int {
	if blinks == 0 {
		return 1
	}
	key := fmt.Sprintf("%d_%d", stone, blinks)
	if v, ok := memory[key]; ok {
		return v
	}

	count := 0
	if stone == 0 {
		count = compute(1, blinks-1, memory)
	} else if str := fmt.Sprintf("%d", stone); len(str)%2 == 0 {
		n1, _ := strconv.Atoi(str[:len(str)/2])
		n2, _ := strconv.Atoi(str[len(str)/2:])
		count = compute(n1, blinks-1, memory) + compute(n2, blinks-1, memory)
	} else {
		count = compute(stone*2024, blinks-1, memory)
	}
	memory[key] = count

	return count
}

func elaborate(data []int, times int) int {
	list := slices.Clone(data)

	memory := map[string]int{}

	count := 0
	for _, d := range list {
		count += compute(d, times, memory)
	}
	return count
}

func part1(data []int) {
	sum := 0

	sum = elaborate(data, 25)

	fmt.Printf("Part 1: %d\n", sum)
}

func part2(data []int) {
	sum := 0

	sum = elaborate(data, 75)

	fmt.Printf("Part 2: %d\n", sum)
}
