package main

import (
	"fmt"

	"advent2016/util"
)

func main() {
	data := util.GetDataByRow("day13")
	// data := util.GetTestByRow("day13")

	num := util.ToInt(data[0])

	part1(num)
	part2(num)
}

type Point struct {
	x int
	y int
}

func printGrid(w, h, num int) {
	fmt.Println("----------------------------------")
	for y := range h {
		for x := range w {
			wall := isWall(num, Point{x, y})
			if wall {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("----------------------------------")
}

func isWall(favNum int, p Point) bool {
	t := p.x*p.x + 3*p.x + 2*p.x*p.y + p.y + p.y*p.y + favNum

	wall := false
	i := 1
	for i <= t {
		bit := t & i
		if bit > 0 {
			wall = !wall
		}
		i *= 2
	}

	return wall
}

func part1(num int) {
	counter := 0

	q := []Point{{1, 1}}
	grid := map[Point]int{q[0]: 0}
	walls := map[Point]bool{q[0]: false}

	target := Point{31, 39}

	var p Point
	for len(q) > 0 {
		p, q = q[0], q[1:]

		dist := grid[p]

		if p == target {
			counter = dist
			break
		}

		// up
		if p.y > 0 {
			up := Point{p.x, p.y - 1}
			if _, kk := grid[up]; !kk {
				var wall bool
				var ok bool
				if wall, ok = walls[up]; !ok {
					wall = isWall(num, up)
					walls[up] = wall
				}
				if !wall {
					grid[up] = dist + 1
					q = append(q, up)
				}
			}
		}

		// down
		down := Point{p.x, p.y + 1}
		if _, kk := grid[down]; !kk {
			var wall bool
			var ok bool
			if wall, ok = walls[down]; !ok {
				wall = isWall(num, down)
				walls[down] = wall
			}
			if !wall {
				grid[down] = dist + 1
				q = append(q, down)
			}
		}

		// left
		if p.x > 0 {
			left := Point{p.x - 1, p.y}
			if _, kk := grid[left]; !kk {
				var wall bool
				var ok bool
				if wall, ok = walls[left]; !ok {
					wall = isWall(num, left)
					walls[left] = wall
				}
				if !wall {
					grid[left] = dist + 1
					q = append(q, left)
				}
			}
		}

		// right
		right := Point{p.x + 1, p.y}
		if _, kk := grid[right]; !kk {
			var wall bool
			var ok bool
			if wall, ok = walls[right]; !ok {
				wall = isWall(num, right)
				walls[right] = wall
			}
			if !wall {
				grid[right] = dist + 1
				q = append(q, right)
			}
		}
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(num int) {
	counter := 0

	q := []Point{{1, 1}}
	grid := map[Point]int{q[0]: 0}
	walls := map[Point]bool{q[0]: false}

	var p Point
	for len(q) > 0 {
		p, q = q[0], q[1:]

		dist := grid[p]

		if dist+1 > 50 {
			continue
		}

		// up
		if p.y > 0 {
			up := Point{p.x, p.y - 1}
			if _, kk := grid[up]; !kk {
				var wall bool
				var ok bool
				if wall, ok = walls[up]; !ok {
					wall = isWall(num, up)
					walls[up] = wall
				}
				if !wall {
					grid[up] = dist + 1
					q = append(q, up)
				}
			}
		}

		// down
		down := Point{p.x, p.y + 1}
		if _, kk := grid[down]; !kk {
			var wall bool
			var ok bool
			if wall, ok = walls[down]; !ok {
				wall = isWall(num, down)
				walls[down] = wall
			}
			if !wall {
				grid[down] = dist + 1
				q = append(q, down)
			}
		}

		// left
		if p.x > 0 {
			left := Point{p.x - 1, p.y}
			if _, kk := grid[left]; !kk {
				var wall bool
				var ok bool
				if wall, ok = walls[left]; !ok {
					wall = isWall(num, left)
					walls[left] = wall
				}
				if !wall {
					grid[left] = dist + 1
					q = append(q, left)
				}
			}
		}

		// right
		right := Point{p.x + 1, p.y}
		if _, kk := grid[right]; !kk {
			var wall bool
			var ok bool
			if wall, ok = walls[right]; !ok {
				wall = isWall(num, right)
				walls[right] = wall
			}
			if !wall {
				grid[right] = dist + 1
				q = append(q, right)
			}
		}
	}

	counter = len(grid)
	fmt.Printf("Part 2: %d\n", counter)
}
