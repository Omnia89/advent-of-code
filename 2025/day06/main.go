package main

import (
	"fmt"
	"regexp"
	"strings"

	"advent2025/util"
)

func main() {
	data := util.GetDataByRow("day06")
	// data := util.GetTestByRow("day06")

	numbers := data[:len(data)-1]

	trim := regexp.MustCompile(`\s+`)
	trimmed := strings.TrimSpace(data[len(data)-1])
	trimmed = trim.ReplaceAllString(trimmed, " ")
	ops := strings.Split(trimmed, " ")
	operations := ops

	part1(numbers, operations)
	part2(numbers, operations)
}

func part1(numbers []string, operations []string) {
	counter := 0

	intNumbers := [][]int{}

	for _, n := range numbers {
		trim := regexp.MustCompile(`\s+`)
		trimmed := trim.ReplaceAllString(n, " ")
		ns := strings.Split(trimmed, " ")

		for i, nn := range ns {
			if nn == "" {
				continue
			}
			if len(intNumbers) <= i {
				intNumbers = append(intNumbers, []int{})
			}

			intNumbers[i] = append(intNumbers[i], util.ToInt(nn))
		}
	}

	for i, op := range operations {
		if op == "" {
			continue
		}
		c := 0
		if op == "*" {
			c = 1
		}

		for _, n := range intNumbers[i] {
			if op == "*" {
				c *= n
			} else {
				c += n
			}
		}
		counter += c
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(numbers []string, operations []string) {
	counter := 0

	intNumbers := [][]int{}

	index := 0
	for i := 0; i < len(numbers[0]); i++ {
		allSpaces := true
		num := ""

		for r := 0; r < len(numbers); r++ {
			if string(numbers[r][i]) == " " {
				continue
			}
			allSpaces = false

			num += string(numbers[r][i])
		}

		if allSpaces {
			index++
		} else {
			if index >= len(intNumbers) {
				intNumbers = append(intNumbers, []int{})
			}
			intNumbers[index] = append(intNumbers[index], util.ToInt(num))
		}

	}

	for i, op := range operations {
		if op == "" {
			continue
		}
		c := 0
		if op == "*" {
			c = 1
		}

		for _, n := range intNumbers[i] {
			if op == "*" {
				c *= n
			} else {
				c += n
			}
		}
		counter += c
	}
	fmt.Printf("Part 2: %d\n", counter)
}
