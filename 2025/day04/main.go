package main

import (
	"fmt"

	"advent2025/util"
)

func main() {
	rows := util.GetDataByRow("day04")
	// rows := util.GetTestByRow("day04")

	part1(rows)
	part2(rows)
}

func neighbornCount(data []string, r int, c int) int {
	count := 0
	for ri := -1; ri <= 1; ri++ {
		for ci := -1; ci <= 1; ci++ {
			if ri == 0 && ci == 0 {
				continue
			}
			nR := r + ri
			nC := c + ci

			if nR < 0 || nC < 0 {
				continue
			}
			if nR >= len(data) || nC >= len(data[0]) {
				continue
			}

			if string(data[nR][nC]) == "@" {
				count += 1
			}
		}
	}
	return count
}

func part1(data []string) {
	counter := 0

	for r := 0; r < len(data); r++ {
		for c := 0; c < len(data[0]); c++ {
			if string(data[r][c]) != "@" {
				continue
			}
			neigh := neighbornCount(data, r, c)
			if neigh < 4 {
				counter += 1
			}
		}
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data []string) {
	counter := 0

	removed := -1

	for removed != 0 {
		removed = 0
		for r := 0; r < len(data); r++ {
			for c := 0; c < len(data[0]); c++ {
				if string(data[r][c]) != "@" {
					continue
				}
				neigh := neighbornCount(data, r, c)
				if neigh < 4 {
					removed += 1
					// TODO: forse bisogna rimuovere alla fine della mandata
					data[r] = data[r][:c] + "." + data[r][c+1:]
				}
			}
		}
		counter += removed
	}

	fmt.Printf("Part 2: %d\n", counter)
}
