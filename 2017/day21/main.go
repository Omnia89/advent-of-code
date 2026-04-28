package main

import (
	"fmt"
	"strings"

	"advent2017/util"
)

func main() {
	data := util.GetDataByRow("day21")
	// data := util.GetTestByRow("day21")

	list := parse(data)

	part1(list)
	part2(list)
}

type Pattern struct {
	in  [][]int
	out [][]int
}

type Ord int

const (
	OrStd Ord = iota
	OrStdFlipH
	OrStdFlipV
	Or90
	Or90FlipH
	Or90FlipV
	Or180
	// OrStd180FlipH // same as OrStdFlipV
	// Or180FlipV // same as OrStdFlipH
	Or270
	// Or270FlipH // same as Or90FlipV
	// Or270FlipV // same as Or90
)

type getCoordFunc = func(x, y, size int) (nX int, nY int)

func (p Pattern) compare(d [][]int, orientation Ord) bool {
	// always compare `p` as is

	// rotate/flip `d`
	var getCoord getCoordFunc

	switch orientation {
	case OrStd:
		getCoord = func(x, y, s int) (int, int) {
			return x, y
		}
	case OrStdFlipH:
		getCoord = func(x, y, s int) (int, int) {
			return x, s - y - 1
		}
	case OrStdFlipV:
		getCoord = func(x, y, s int) (int, int) {
			return s - x - 1, y
		}
	case Or90:
		getCoord = func(x, y, s int) (int, int) {
			return s - y - 1, x
		}
	case Or90FlipH:
		getCoord = func(x, y, s int) (int, int) {
			return s - y - 1, s - x - 1
		}
	case Or90FlipV:
		getCoord = func(x, y, s int) (int, int) {
			return y, x
		}
	case Or180:
		getCoord = func(x, y, s int) (int, int) {
			return s - x - 1, s - y - 1
		}
	case Or270:
		getCoord = func(x, y, s int) (int, int) {
			return y, s - x - 1
		}
	}

	size := len(p.in)

	for y := range size {
		for x := range size {
			dX, dY := getCoord(x, y, size)
			if p.in[y][x] != d[dY][dX] {
				return false
			}
		}
	}
	return true
}

func (p Pattern) compareAll(d [][]int) bool {
	if len(p.in) != len(d) {
		return false
	}
	ords := []Ord{OrStd, OrStdFlipH, OrStdFlipV, Or90, Or90FlipH, Or90FlipV, Or180, Or270}

	for _, o := range ords {
		kk := p.compare(d, o)
		if kk {
			return true
		}
	}
	return false
}

func parse(data []string) []Pattern {
	ps := []Pattern{}

	for _, s := range data {
		t := strings.ReplaceAll(s, ".", "0")
		t = strings.ReplaceAll(t, "#", "1")
		in, out, _ := strings.Cut(t, " => ")

		inParts := strings.Split(in, "/")
		outParts := strings.Split(out, "/")

		inData := [][]int{}
		for _, l := range inParts {
			inData = append(inData, util.StringToIntSlice(l, ""))
		}
		outData := [][]int{}
		for _, l := range outParts {
			outData = append(outData, util.StringToIntSlice(l, ""))
		}

		ps = append(ps, Pattern{
			in:  inData,
			out: outData,
		})
	}

	return ps
}

func getSubGrid(grid [][]int, size int, startX int, startY int) (bool, [][]int) {
	if startY+size > len(grid) || startX+size > len(grid[0]) {
		return false, nil
	}

	r := make([][]int, size)
	for y := range size {
		r[y] = make([]int, size)
		for x := range size {
			r[y][x] = grid[y+startY][x+startX]
		}
	}
	return true, r
}

func getEmptyProgram(size int) [][]int {
	p := make([][]int, size)
	for i := range p {
		p[i] = make([]int, size)
	}
	return p
}

func iterate(data []Pattern, times int) int {
	counter := 0

	program := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
	}

	for range times {
		size := len(program)

		outSize := 4
		patchSize := 3
		if size%2 == 0 {
			outSize = 3
			patchSize = 2
		}

		patches := size / patchSize
		newProgram := getEmptyProgram(patches * outSize)

		startY := 0
		for startY < size {
			startX := 0
			for startX < size {
				ok, subG := getSubGrid(program, patchSize, startX, startY)
				if !ok {
					continue
				}

			patternLoop:
				for _, p := range data {
					if ok := p.compareAll(subG); ok {
						oX := (startX / patchSize) * outSize
						oY := (startY / patchSize) * outSize

						for gY, row := range p.out {
							for gX, v := range row {
								newProgram[gY+oY][gX+oX] = v
							}
						}
						break patternLoop
					}
				}
				startX += patchSize
			}
			startY += patchSize
		}
		program = newProgram
	}

	for _, r := range program {
		for _, v := range r {
			counter += v
		}
	}
	return counter
}

func part1(data []Pattern) {
	counter := iterate(data, 5)
	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data []Pattern) {
	counter := iterate(data, 18)

	fmt.Printf("Part 2: %d\n", counter)
}
