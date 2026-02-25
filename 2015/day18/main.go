package main

import (
	"fmt"
	"strconv"
	"strings"

	"advent2015/util"
)

func main() {
	data := util.GetDataByRow("day18")
	// data := util.GetTestByRow("day18")

	matrix := parse(data)

	part1(matrix)
	part2(matrix)
}

func parse(data []string) [][]int {
	r := [][]int{}
	for _, s := range data {
		t := strings.ReplaceAll(s, ".", "0")
		t = strings.ReplaceAll(t, "#", "1")
		r = append(r, util.StringToIntSlice(t, ""))
	}
	return r
}

func newMatrix(w int, h int) [][]int {
	n := make([][]int, h)
	for i := range n {
		n[i] = make([]int, w)
	}
	return n
}

func countNear(data [][]int, x int, y int) int {
	n := 0

	op := []int{-1, 0, 1}

	for _, dy := range op {
		for _, dx := range op {
			if dx == 0 && dy == 0 {
				continue
			}
			ny := y + dy
			nx := x + dx
			if ny < 0 || nx < 0 || ny >= len(data) || nx >= len(data[0]) {
				continue
			}
			n += data[ny][nx]
		}
	}
	return n
}

func tick(matrix [][]int) [][]int {
	nMatrix := newMatrix(len(matrix[0]), len(matrix))

	for y := range matrix {
		for x, v := range matrix[y] {
			n := countNear(matrix, x, y)
			if n == 3 || (n == 2 && v == 1) {
				nMatrix[y][x] = 1
			}
		}
	}
	return nMatrix
}

func countOn(matrix [][]int) int {
	n := 0

	for y := range matrix {
		for _, v := range matrix[y] {
			n += v
		}
	}
	return n
}

func printMatrix(matrix [][]int) {
	fmt.Println("----------------------------------")
	for _, r := range matrix {
		ss := make([]string, len(r))
		for i := range r {
			ss[i] = strconv.Itoa(r[i])
		}
		fmt.Printf("%s\n", strings.Join(ss, ""))
	}
	fmt.Println("----------------------------------")
}

func part1(data [][]int) {
	counter := 0

	for range 100 {
		data = tick(data)
	}

	counter = countOn(data)

	fmt.Printf("Part 1: %d\n", counter)
}

func stuckCorners(matrix [][]int) {
	matrix[0][0] = 1
	matrix[0][len(matrix[0])-1] = 1
	matrix[len(matrix)-1][0] = 1
	matrix[len(matrix)-1][len(matrix[0])-1] = 1
}

func part2(data [][]int) {
	counter := 0

	stuckCorners(data)

	for range 100 {
		data = tick(data)
		stuckCorners(data)
	}

	counter = countOn(data)
	fmt.Printf("Part 2: %d\n", counter)
}
