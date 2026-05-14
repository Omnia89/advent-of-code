package main

import (
	"fmt"

	"advent2018/util"
)

func main() {
	data := util.GetDataByRow("day11")
	// data := util.GetTestByRow("day11")

	sn := util.ToInt(data[0])

	part1(sn)
	part2(sn)
}

func sumX(grid [][]int, x int, y int, size int) int {
	s := 0
	for dy := range size {
		for dx := range size {
			s += grid[y+dy][x+dx]
		}
	}
	return s
}

func part1(sn int) {
	counter := 0

	grid := make([][]int, 300)
	for y := range grid {
		grid[y] = make([]int, 300)
	}

	for y := range grid {
		for x := range grid {
			rY := y + 1
			rX := x + 1

			id := rX + 10
			pl := (id*rY + sn) * id

			pl = (pl/100)%10 - 5
			grid[y][x] = pl
		}
	}

	pX, pY := 0, 0

	for y := range len(grid) - 3 {
		for x := range len(grid[0]) - 3 {
			s := sumX(grid, x, y, 3)
			if s > counter {
				counter = s
				pX = x + 1
				pY = y + 1
			}
		}
	}

	res := fmt.Sprintf("%d,%d", pX, pY)

	fmt.Printf("Part 1: %s\n", res)
}

// Get the summed area summed area table
//   - every point is the sum of the relative grid point plus all the points at his left and above
//   - to get the sum of an area with size of K, get the value of the bottom-right corner, minus up-right, minus bottom-left, plus up-left
func summedAreaTable(grid [][]int) [][]int {
	summed := make([][]int, len(grid))
	for y := range summed {
		summed[y] = make([]int, len(grid[0]))
	}

	for y := range len(grid) {
		for x := range len(grid[0]) {
			s := grid[y][x]
			if x > 0 {
				s += summed[y][x-1]
			}
			if y > 0 {
				s += summed[y-1][x]
			}
			if x > 0 && y > 0 {
				s -= summed[y-1][x-1]
			}
			summed[y][x] = s
		}
	}
	return summed
}

func part2(sn int) {
	counter := 0

	grid := make([][]int, 300)
	for y := range grid {
		grid[y] = make([]int, 300)
	}

	for y := range grid {
		for x := range grid {
			rY := y + 1
			rX := x + 1

			id := rX + 10
			pl := (id*rY + sn) * id

			pl = (pl/100)%10 - 5
			grid[y][x] = pl
		}
	}

	summed := summedAreaTable(grid)

	pX, pY, pSize := 0, 0, 0

	for size := 1; size <= 300; size++ {
		for y := size; y < 300; y++ {
			for x := size; x < 300; x++ {
				sum := summed[y][x]
				if y-size >= 0 {
					sum -= summed[y-size][x]
				}
				if x-size >= 0 {
					sum -= summed[y][x-size]
				}
				if y-size >= 0 && x-size >= 0 {
					sum += summed[y-size][x-size]
				}
				if sum > counter {
					counter = sum
					// minus size plus 2 to get upper left
					pX = x - size + 2
					pY = y - size + 2
					pSize = size
				}
			}
		}
	}

	res := fmt.Sprintf("%d,%d,%d", pX, pY, pSize)
	fmt.Printf("Part 2: %s\n", res)
}
