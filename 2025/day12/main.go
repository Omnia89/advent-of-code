package main

import (
	"fmt"
	"strings"

	"advent2025/util"
)

func main() {
	data := util.GetDataByRow("day12")
	// data := util.GetTestByRow("day12")

	shapes, areas := parse(data)

	part1(shapes, areas)
	// part2(data)
}

type Shape [][]bool

func (s Shape) ToString() string {
	st := ""
	for _, rr := range s {
		for _, v := range rr {
			if v {
				st += "#"
			} else {
				st += " "
			}
		}
		st += "\n"
	}
	return st
}

func (s Shape) Rotate(degree int) Shape {
	nn := Shape{
		[]bool{false, false, false},
		[]bool{false, false, false},
		[]bool{false, false, false},
	}
	if degree == 180 {
		for i := len(s) - 1; i >= 0; i-- {
			rev := []bool{}
			for j := len(s[i]) - 1; j >= 0; j-- {
				rev = append(rev, s[i][j])
			}
			nn = append(nn, rev)
		}
	} else if degree == 90 {
		for r := range 3 {
			for c := range 3 {
				nn[c][2-r] = s[r][c]
			}
		}
	} else if degree == 270 {
		for r := range 3 {
			for c := range 3 {
				nn[2-c][r] = s[r][c]
			}
		}
	} else {
		nn = s
	}

	return nn
}

func (s Shape) Flip(direction string) Shape {
	nn := Shape{
		[]bool{false, false, false},
		[]bool{false, false, false},
		[]bool{false, false, false},
	}

	if direction == "h" {
		for r := range 3 {
			for c := range 3 {
				nn[r][c] = s[2-r][c]
			}
		}
	} else if direction == "v" {
		for r := range 3 {
			for c := range 3 {
				nn[r][c] = s[r][2-c]
			}
		}
	} else {
		nn = s
	}
	return nn
}

func (s Shape) AreaValue() int {
	counter := 0

	for r := range 3 {
		for c := range 3 {
			if s[r][c] {
				counter++
			}
		}
	}

	return counter
}

type Area struct {
	Height      int
	Width       int
	ShapesCount []int
}

func (a Area) GetTable() [][]bool {
	table := [][]bool{}
	for range a.Height {
		table = append(table, make([]bool, a.Width))
	}
	return table
}

func parse(rows []string) (shapes []Shape, areas []Area) {
	tmpShape := Shape{}

	for _, row := range rows {
		if row == "" {
			if len(tmpShape) > 0 {
				shapes = append(shapes, tmpShape)
			}
			tmpShape = Shape{}
		} else {
			if strings.HasPrefix(row, ".") || strings.HasPrefix(row, "#") {
				tmpRow := []bool{}
				for _, c := range row {
					tmpRow = append(tmpRow, string(c) == "#")
				}
				tmpShape = append(tmpShape, tmpRow)
			} else {
				if len(row) == 2 {
					continue
				}
				area := Area{}
				parts := strings.Split(row, ":")
				dimensions := util.StringToIntSlice(parts[0], "x")
				area.Height = dimensions[1]
				area.Width = dimensions[0]
				area.ShapesCount = util.StringToIntSlice(parts[1], " ")
				areas = append(areas, area)
			}
		}
	}
	return
}

func checkAreaSize(shapes []Shape, area Area) bool {
	shapeArea := 0
	areaArea := area.Height * area.Width
	for idx, val := range area.ShapesCount {
		shapeArea += shapes[idx].AreaValue() * val
	}
	return shapeArea <= areaArea
}

func looseShapeContained(area Area) bool {
	maxHorizontal := area.Width / 3
	maxVertical := area.Height / 3
	maxPieces := maxHorizontal * maxVertical

	totalPieces := 0
	for _, s := range area.ShapesCount {
		totalPieces += s
	}

	return totalPieces <= maxPieces
}

func part1(shapes []Shape, areas []Area) {
	possible := 0
	maybe := 0
	impossible := 0

	for _, area := range areas {
		check := checkAreaSize(shapes, area)
		loose := looseShapeContained(area)
		if !check {
			impossible++
			continue
		}
		if loose {
			possible++
			continue
		}
		maybe++
	}

	result := fmt.Sprintf("Yes [%d], Maybe [%d], Impossible [%d]", possible, maybe, impossible)

	fmt.Printf("Part 1: %s\n", result)
}
