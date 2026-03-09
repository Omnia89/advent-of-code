package main

import (
	"fmt"

	"advent2016/util"
)

func main() {
	data := util.GetDataByRow("day18")
	// data := util.GetTestByRow("day18")

	row := parse(data[0])

	part1(row)
	part2(row)
}

func parse(s string) []int {
	r := make([]int, 0, len(s))
	for _, c := range s {
		if c == '^' {
			r = append(r, 1)
		} else {
			r = append(r, 0)
		}
	}
	return r
}

func getNext(r []int) []int {
	n := make([]int, len(r))

	for i := range r {
		left := 0
		right := 0
		if i > 0 {
			left = r[i-1]
		}
		if i < len(r)-1 {
			right = r[i+1]
		}
		n[i] = left ^ right
	}

	return n
}

func countVal(r []int, val int) int {
	s := 0
	for _, n := range r {
		if n == val {
			s++
		}
	}
	return s
}

func part1(row []int) {
	counter := 0

	rows := [][]int{row}
	counter += countVal(row, 0)

	for len(rows) < 40 {
		nR := getNext(rows[len(rows)-1])
		counter += countVal(nR, 0)
		rows = append(rows, nR)
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(row []int) {
	counter := 0

	rows := [][]int{row}
	counter += countVal(row, 0)

	for len(rows) < 400000 {
		nR := getNext(rows[len(rows)-1])
		counter += countVal(nR, 0)
		rows = append(rows, nR)
	}
	fmt.Printf("Part 2: %d\n", counter)
}
