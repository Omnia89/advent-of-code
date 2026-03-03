package main

import (
	"fmt"

	"advent2016/util"
)

func main() {
	data := util.GetDataByRow("day02")
	// data := util.GetTestByRow("day02")

	part1(data)
	part2(data)
}

func part1(data []string) {
	response := ""

	row := 1
	col := 1

	getPadNum := func(r, c int) int {
		return r*3 + c + 1
	}

	moves := map[string][]int{
		"U": {-1, 0},
		"D": {1, 0},
		"R": {0, 1},
		"L": {0, -1},
	}

	for _, s := range data {
		for _, d := range s {
			dir := string(d)
			row += moves[dir][0]
			row = util.IntMax(util.IntMin(row, 2), 0)

			col += moves[dir][1]
			col = util.IntMax(util.IntMin(col, 2), 0)
		}
		response = fmt.Sprintf("%s%d", response, getPadNum(row, col))
	}

	fmt.Printf("Part 1: %s\n", response)
}

func part2(data []string) {
	response := ""

	pad := []string{
		"  1  ",
		" 234 ",
		"56789",
		" ABC ",
		"  D  ",
	}

	moves := map[string][]int{
		"U": {-1, 0},
		"D": {1, 0},
		"R": {0, 1},
		"L": {0, -1},
	}

	getNext := func(r, c, dR, dC int) (int, int) {
		newR := r + dR
		newC := c + dC

		if newR < 0 || newR > 4 || newC < 0 || newC > 4 {
			return r, c
		}

		if pad[newR][newC] == ' ' {
			return r, c
		}
		return newR, newC
	}

	row := 2
	col := 0

	for _, s := range data {
		for _, d := range s {
			dir := string(d)
			row, col = getNext(row, col, moves[dir][0], moves[dir][1])
		}
		response = fmt.Sprintf("%s%s", response, string(pad[row][col]))
	}

	fmt.Printf("Part 2: %s\n", response)
}
