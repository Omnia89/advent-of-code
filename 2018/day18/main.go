package main

import (
	"crypto/sha256"
	"fmt"
	"slices"
	"strings"

	"advent2018/util"
)

func main() {
	data := util.GetDataByRow("day18")
	// data := util.GetTestByRow("day18")

	area := parse(data)

	part1(area)
	part2(area)
}

type Cell int

const (
	Open Cell = iota
	Tree
	Lumberjard
)

func newGrid(h, w int) [][]Cell {
	g := make([][]Cell, h)
	for i := range h {
		g[i] = make([]Cell, w)
	}
	return g
}

func parse(data []string) [][]Cell {
	g := newGrid(len(data), len(data[0]))

	for y, r := range data {
		for x, ch := range r {
			var c Cell
			switch ch {
			case '.':
				c = Open
			case '|':
				c = Tree
			case '#':
				c = Lumberjard
			}
			g[y][x] = c
		}
	}

	return g
}

func becomeTree(area [][]Cell, x, y int) bool {
	if area[y][x] != Open {
		return false
	}
	treeCount := 0

	for iy := -1; iy <= 1; iy++ {
		for ix := -1; ix <= 1; ix++ {
			if ix == 0 && iy == 0 {
				continue
			}
			ry := y + iy
			rx := x + ix
			if rx < 0 || ry < 0 || ry >= len(area) || rx >= len(area[0]) {
				continue
			}
			if area[ry][rx] == Tree {
				treeCount++
			}
		}
	}

	return treeCount >= 3
}

func becomeLumberjard(area [][]Cell, x, y int) bool {
	if area[y][x] != Tree {
		return false
	}
	lumberCount := 0

	for iy := -1; iy <= 1; iy++ {
		for ix := -1; ix <= 1; ix++ {
			if ix == 0 && iy == 0 {
				continue
			}
			ry := y + iy
			rx := x + ix
			if rx < 0 || ry < 0 || ry >= len(area) || rx >= len(area[0]) {
				continue
			}
			if area[ry][rx] == Lumberjard {
				lumberCount++
			}
		}
	}

	return lumberCount >= 3
}

func becomeOpen(area [][]Cell, x, y int) bool {
	if area[y][x] != Lumberjard {
		return false
	}
	treeCount := 0
	lumberCount := 0

	for iy := -1; iy <= 1; iy++ {
		for ix := -1; ix <= 1; ix++ {
			if ix == 0 && iy == 0 {
				continue
			}
			ry := y + iy
			rx := x + ix
			if rx < 0 || ry < 0 || ry >= len(area) || rx >= len(area[0]) {
				continue
			}
			if area[ry][rx] == Tree {
				treeCount++
			}
			if area[ry][rx] == Lumberjard {
				lumberCount++
			}
		}
	}

	return treeCount < 1 || lumberCount < 1
}

func printArea(area [][]Cell) {
	var sb strings.Builder

	sb.WriteString("-----------------------------------------\n")

	for y, row := range area {
		for x := range row {
			switch area[y][x] {
			case Open:
				sb.WriteString(".")
			case Tree:
				sb.WriteString("|")
			case Lumberjard:
				sb.WriteString("#")
			}
		}
		sb.WriteString("\n")
	}
	sb.WriteString("-----------------------------------------\n")
	fmt.Print(sb.String())
}

func step(area [][]Cell) [][]Cell {
	newArea := newGrid(len(area), len(area[0]))

	for y, row := range area {
		for x := range row {
			if becomeTree(area, x, y) {
				newArea[y][x] = Tree
			} else if becomeLumberjard(area, x, y) {
				newArea[y][x] = Lumberjard
			} else if becomeOpen(area, x, y) {
				newArea[y][x] = Open
			} else {
				newArea[y][x] = area[y][x]
			}
		}
	}
	return newArea
}

func part1(area [][]Cell) {
	counter := 0

	// printArea(area)
	for range 10 {
		area = step(area)
		// printArea(area)
	}

	trees := 0
	lumbers := 0
	for y, row := range area {
		for x := range row {
			if area[y][x] == Tree {
				trees++
			}
			if area[y][x] == Lumberjard {
				lumbers++
			}
		}
	}

	counter = trees * lumbers

	fmt.Printf("Part 1: %d\n", counter)
}

func areaHash(area [][]Cell) [32]byte {
	var sb strings.Builder
	for y, r := range area {
		for x := range r {
			fmt.Fprintf(&sb, "%d", area[y][x])
		}
	}
	return sha256.Sum256([]byte(sb.String()))
}

// 183040 too low
func part2(area [][]Cell) {
	counter := 0

	cache := [][][]Cell{}
	hashCache := [][32]byte{}

	TOT := 1_000_000_000

	// printArea(area)
	for i := range TOT {
		area = step(area)
		hash := areaHash(area)
		if firstI := slices.Index(hashCache, hash); firstI >= 0 {
			index := (TOT-1-firstI)%(i-firstI) + firstI
			area = cache[index]
			break
		}
		cache = append(cache, area)
		hashCache = append(hashCache, hash)

		// printArea(area)
	}

	trees := 0
	lumbers := 0
	for y, row := range area {
		for x := range row {
			if area[y][x] == Tree {
				trees++
			}
			if area[y][x] == Lumberjard {
				lumbers++
			}
		}
	}

	counter = trees * lumbers
	fmt.Printf("Part 2: %d\n", counter)
}
