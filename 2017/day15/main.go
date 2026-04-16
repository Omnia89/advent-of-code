package main

import (
	"fmt"
	"strings"

	"advent2017/util"
)

func main() {
	data := util.GetDataByRow("day15")
	// data := util.GetTestByRow("day15")

	a, b := parse(data)

	part1(a, b)
	part2(a, b)
}

func parse(data []string) (genA int, genB int) {
	genA = util.ToInt(strings.Split(data[0], " ")[4])
	genB = util.ToInt(strings.Split(data[1], " ")[4])
	return
}

const (
	aFactor   = 16807
	bFactor   = 48271
	remainder = 2147483647
)

func compare16Bits(a, b int) int {
	aBits := a & 0xFFFF
	bBits := b & 0xFFFF

	if aBits == bBits {
		return 1
	}
	return 0
}

func part1(a, b int) {
	counter := 0

	aValue := a
	bValue := b

	for range 40000000 {
		aValue = (aValue * aFactor) % remainder
		bValue = (bValue * bFactor) % remainder

		counter += compare16Bits(aValue, bValue)
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(a, b int) {
	counter := 0

	aValue := a
	bValue := b

	for range 5000000 {
		aValue = (aValue * aFactor) % remainder
		for aValue%4 != 0 {
			aValue = (aValue * aFactor) % remainder
		}

		bValue = (bValue * bFactor) % remainder
		for bValue%8 != 0 {
			bValue = (bValue * bFactor) % remainder
		}

		counter += compare16Bits(aValue, bValue)
	}
	fmt.Printf("Part 2: %d\n", counter)
}
