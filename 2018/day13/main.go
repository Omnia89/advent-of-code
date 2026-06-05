package main

import (
	"fmt"
	"slices"
	"strings"

	"advent2018/util"
)

func main() {
	data := util.GetDataByRow("day13")
	//data := util.GetTestByRow("day13")

	circuit, carts := parse(data)

	carts2 := make([]Cart, len(carts))
	copy(carts2, carts)

	part1(circuit, carts)
	part2(circuit, carts2)
}

type Point struct {
	x int
	y int
}

type Cart struct {
	pos       Point
	direction Point
	turn      string
}

// [direction][turn]newDirection
var turns = map[Point]map[string]Point{
	// up
	{0, -1}: {
		"left":  {-1, 0},
		"right": {1, 0},
	},
	// down
	{0, 1}: {
		"left":  {1, 0},
		"right": {-1, 0},
	},
	// left
	{-1, 0}: {
		"left":  {0, 1},
		"right": {0, -1},
	},
	// right
	{1, 0}: {
		"left":  {0, -1},
		"right": {0, 1},
	},
}

var turnOrder = map[string]string{
	"left":     "straight",
	"straight": "right",
	"right":    "left",
}

func (c *Cart) makeTurn() {
	// left -> straight -> right -> ...

	if c.turn != "straight" {
		c.direction = turns[c.direction][c.turn]
	}

	c.turn = turnOrder[c.turn]
}

var curves = map[Point]map[rune]Point{
	// up
	{0, -1}: {
		'/':  {1, 0},
		'\\': {-1, 0},
	},
	// down
	{0, 1}: {
		'/':  {-1, 0},
		'\\': {1, 0},
	},
	// left
	{-1, 0}: {
		'/':  {0, 1},
		'\\': {0, -1},
	},
	// right
	{1, 0}: {
		'/':  {0, -1},
		'\\': {0, 1},
	},
}

func (c *Cart) move(circuit []string, carts []Cart) (collision bool, crashPos Point) {
	oldPosition := c.pos
	// move forward
	c.pos.x += c.direction.x
	c.pos.y += c.direction.y
	// check collision
	for _, nc := range carts {
		if oldPosition == nc.pos {
			continue
		}
		if c.pos == nc.pos {
			return true, c.pos
		}
	}

	// turn on curves/crossway
	switch p := circuit[c.pos.y][c.pos.x]; p {
	case '/', '\\':
		c.direction = curves[c.direction][rune(p)]
	case '+':
		c.makeTurn()
	}
	return false, Point{}
}

func parse(data []string) (circuit []string, carts []Cart) {
	for y, s := range data {
		t := s

		if strings.Contains(t, "^") {
			x := strings.Index(t, "^")
			for x != -1 {
				carts = append(carts, Cart{
					pos:       Point{x, y},
					direction: Point{0, -1},
					turn:      "left",
				})
				t = strings.Replace(t, "^", "|", 1)
				x = strings.Index(t, "^")
			}
		}

		if strings.Contains(t, "v") {
			x := strings.Index(t, "v")
			for x != -1 {
				carts = append(carts, Cart{
					pos:       Point{x, y},
					direction: Point{0, 1},
					turn:      "left",
				})
				t = strings.Replace(t, "v", "|", 1)
				x = strings.Index(t, "v")
			}
		}

		if strings.Contains(t, "<") {
			x := strings.Index(t, "<")
			for x != -1 {
				carts = append(carts, Cart{
					pos:       Point{x, y},
					direction: Point{-1, 0},
					turn:      "left",
				})
				t = strings.Replace(t, "<", "-", 1)
				x = strings.Index(t, "<")
			}
		}

		if strings.Contains(t, ">") {
			x := strings.Index(t, ">")
			for x != -1 {
				carts = append(carts, Cart{
					pos:       Point{x, y},
					direction: Point{1, 0},
					turn:      "left",
				})
				t = strings.Replace(t, ">", "-", 1)
				x = strings.Index(t, ">")
			}
		}
		circuit = append(circuit, t)
	}

	return
}

func part1(circuit []string, carts []Cart) {
	sortCarts := func(a Cart, b Cart) int {
		if a.pos.y < b.pos.y {
			return -1
		} else if b.pos.y < a.pos.y {
			return 1
		}
		return a.pos.x - b.pos.x
	}

	var collide Point
external:
	for {
		slices.SortFunc(carts, sortCarts)
		for i := range carts {
			c := carts[i]

			collision, _ := c.move(circuit, carts)

			if collision {
				collide = c.pos
				break external
			}

			carts[i] = c
		}
	}

	res := fmt.Sprintf("%d,%d", collide.x, collide.y)

	fmt.Printf("Part 1: %s\n", res)
}

func part2(circuit []string, carts []Cart) {
	sortCarts := func(a Cart, b Cart) int {
		if a.pos.y < b.pos.y {
			return -1
		} else if b.pos.y < a.pos.y {
			return 1
		}
		return a.pos.x - b.pos.x
	}

	for len(carts) > 1 {
		slices.SortFunc(carts, sortCarts)
		crashes := []Point{}
		for i := range carts {
			c := carts[i]
			if slices.Contains(crashes, c.pos) {
				continue
			}

			collision, crash := c.move(circuit, carts)

			if collision {
				crashes = append(crashes, crash)
			}

			carts[i] = c
		}
		carts = slices.DeleteFunc(carts, func(c Cart) bool {
			return slices.Contains(crashes, c.pos)
		})
	}

	res := fmt.Sprintf("%d,%d", carts[0].pos.x, carts[0].pos.y)

	fmt.Printf("Part 2: %s\n", res)
}
