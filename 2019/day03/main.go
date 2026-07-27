package main

import (
	"fmt"
	"math"
	"strings"

	"advent2019/util"
)

func main() {
	data := util.GetDataByRow("day03")
	//data := util.GetTestByRow("day03")

	list := parse(data)

	part1(list)
	part2(list)
}

type Point struct {
	x int
	y int
}

func parse(data []string) (wires [][]Point) {
	for _, s := range data {
		w := []Point{}
		parts := strings.Split(s, ",")
		p := Point{0, 0}
		w = append(w, p)

		for _, move := range parts {
			distance := util.ToInt(move[1:])
			switch move[0] {
			case 'R':
				p.x += distance
			case 'L':
				p.x -= distance
			case 'D':
				p.y += distance
			case 'U':
				p.y -= distance
			}
			w = append(w, p)
		}
		wires = append(wires, w)
	}
	return wires
}

func isVertical(a, b Point) bool {
	return a.x == b.x
}

func distance(p1, p2 Point) int {
	return util.IntAbs(p1.x-p2.x) + util.IntAbs(p1.y-p2.y)
}

func part1(wires [][]Point) {
	counter := 0

	minDistance := math.MaxInt

	for i1 := range len(wires[0]) - 1 {
		for i2 := range len(wires[1]) - 1 {
			a1, b1 := wires[0][i1], wires[0][i1+1]
			a2, b2 := wires[1][i2], wires[1][i2+1]

			v1 := isVertical(a1, b1)
			v2 := isVertical(a2, b2)

			if v1 == v2 {
				continue
			}

			xL, xM, xH := 0, 0, 0
			yL, yM, yH := 0, 0, 0

			if v1 {
				xL = util.IntMin(a2.x, b2.x)
				xM = a1.x
				xH = util.IntMax(a2.x, b2.x)

				yL = util.IntMin(a1.y, b1.y)
				yM = a2.y
				yH = util.IntMax(a1.y, b1.y)
			} else {

				xL = util.IntMin(a1.x, b1.x)
				xM = a2.x
				xH = util.IntMax(a1.x, b1.x)

				yL = util.IntMin(a2.y, b2.y)
				yM = a1.y
				yH = util.IntMax(a2.y, b2.y)
			}

			// Skip origin
			if (xM != 0 || yM != 0) && xL <= xM && xM <= xH && yL <= yM && yM <= yH {
				crossPoint := Point{xM, yM}
				if d := distance(crossPoint, Point{0, 0}); d < minDistance {
					minDistance = d
				}
			}

		}
	}
	counter = minDistance

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(wires [][]Point) {
	counter := 0

	minDistance := math.MaxInt

	walked1 := 0
	walked2 := 0
	for i1 := range len(wires[0]) - 1 {
		walked2 = 0
		a1, b1 := wires[0][i1], wires[0][i1+1]
		for i2 := range len(wires[1]) - 1 {
			a2, b2 := wires[1][i2], wires[1][i2+1]

			v1 := isVertical(a1, b1)
			v2 := isVertical(a2, b2)

			if v1 == v2 {
				walked2 += distance(a2, b2)
				continue
			}

			xL, xM, xH := 0, 0, 0
			yL, yM, yH := 0, 0, 0

			if v1 {
				xL = util.IntMin(a2.x, b2.x)
				xM = a1.x
				xH = util.IntMax(a2.x, b2.x)

				yL = util.IntMin(a1.y, b1.y)
				yM = a2.y
				yH = util.IntMax(a1.y, b1.y)
			} else {

				xL = util.IntMin(a1.x, b1.x)
				xM = a2.x
				xH = util.IntMax(a1.x, b1.x)

				yL = util.IntMin(a2.y, b2.y)
				yM = a1.y
				yH = util.IntMax(a2.y, b2.y)
			}

			// Skip origin
			if (xM != 0 || yM != 0) && xL <= xM && xM <= xH && yL <= yM && yM <= yH {
				crossPoint := Point{xM, yM}
				walked := walked1 + walked2
				walked += distance(crossPoint, a1)
				walked += distance(crossPoint, a2)
				if walked < minDistance {
					minDistance = walked
				}
			}
			walked2 += distance(a2, b2)
		}
		walked1 += distance(a1, b1)
	}
	counter = minDistance
	fmt.Printf("Part 2: %d\n", counter)
}
