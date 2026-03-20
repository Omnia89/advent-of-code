package main

import (
	"fmt"
	"strings"

	"advent2016/util"
)

func main() {
	data := util.GetDataByRow("day22")
	// data := util.GetTestByRow("day22")

	list := parse(data)

	part1(list)
	part2(list)
}

type Node struct {
	Point
	size      int
	used      int
	available int
}

type Point struct {
	x int
	y int
}

func (n Node) toString() string {
	return fmt.Sprintf("[%03d,%03d] s[%04d] u[%04d] a[%04d]", n.x, n.y, n.size, n.used, n.available)
}

func parse(data []string) []Node {
	nodes := []Node{}

	for i := 2; i < len(data); i++ {
		s := strings.ReplaceAll(data[i], "T", "")
		s = strings.ReplaceAll(s, "y", "")
		s = strings.ReplaceAll(s, "x", "")

		parts := strings.Fields(s)
		coord := strings.Split(parts[0], "-")

		nodes = append(nodes, Node{
			Point{util.ToInt(coord[1]), util.ToInt(coord[2])},
			util.ToInt(parts[1]),
			util.ToInt(parts[2]),
			util.ToInt(parts[3]),
		})
	}
	return nodes
}

func part1(nodes []Node) {
	counter := 0

	for i := range len(nodes) - 1 {
		n1 := nodes[i]
		for j := i + 1; j < len(nodes); j++ {
			n2 := nodes[j]
			if n1.used <= n2.available && n1.used > 0 {
				counter++
			}
			if n2.used <= n1.available && n2.used > 0 {
				counter++
			}
		}
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func printNodeMap(nodes [][]Node) {
	for y := range nodes {
		for _, n := range nodes[y] {
			fmt.Printf("[%3d/%3d] ", n.used, n.size)
		}
		fmt.Println()
	}
}

func printDistanceMap(points map[Point]int, w int, h int) {
	for y := range w {
		for x := range h {
			dist := -1
			if d, ok := points[Point{x, y}]; ok {
				dist = d
			}
			fmt.Printf("[%3d] ", dist)
		}
		fmt.Println()
	}
}

func printPath(points []Node, empty Point, pack Point, dest Point, w int, h int) {
	grid := [][]string{}
	for y := range h {
		grid = append(grid, make([]string, w))
		for x := range w {
			grid[y][x] = " "
		}
	}

	for _, p := range points {
		grid[p.y][p.x] = "o"
	}

	grid[empty.y][empty.x] = "E"
	grid[pack.y][pack.x] = "P"
	grid[dest.y][dest.x] = "X"

	fmt.Printf("%s\n", strings.Repeat("-", w+2))
	for _, s := range grid {
		fmt.Printf("|%s|\n", strings.Join(s, ""))
	}
	fmt.Printf("%s\n", strings.Repeat("-", w+2))
}

func neighbors(p Point) []Point {
	return []Point{
		{p.x - 1, p.y},
		{p.x + 1, p.y},
		{p.x, p.y - 1},
		{p.x, p.y + 1},
	}
}

func findEmpty(start Node, nonTraversable Point, grid [][]Node, w int, h int) []Node {
	tracks := map[Point]int{start.Point: 0}

	q := []Node{start}
	var n Node

	var emptyPoint Point

	for len(q) > 0 {
		n, q = q[0], q[1:]
		distance := tracks[n.Point] + 1

		// up
		if n.y > 0 {
			p := Point{n.x, n.y - 1}
			_, ok := tracks[p]
			pp := grid[p.y][p.x]
			if nonTraversable != p && pp.size >= n.used && !ok {
				q = append(q, pp)
				tracks[p] = distance

				if pp.used == 0 {
					emptyPoint = p
					break
				}
			}
		}
		// right
		if n.x < w-1 {
			p := Point{n.x + 1, n.y}
			_, ok := tracks[p]
			pp := grid[p.y][p.x]
			if nonTraversable != p && pp.size >= n.used && !ok {
				q = append(q, pp)
				tracks[p] = distance

				if pp.used == 0 {
					emptyPoint = p
					break
				}
			}
		}

		// down
		if n.y < h-1 {
			p := Point{n.x, n.y + 1}
			_, ok := tracks[p]
			pp := grid[p.y][p.x]
			if nonTraversable != p && pp.size >= n.used && !ok {
				q = append(q, pp)
				tracks[p] = distance

				if pp.used == 0 {
					emptyPoint = p
					break
				}
			}
		}
		// left
		if n.x > 0 {
			p := Point{n.x - 1, n.y}
			_, ok := tracks[p]
			pp := grid[p.y][p.x]
			if nonTraversable != p && pp.size >= n.used && !ok {
				q = append(q, pp)
				tracks[p] = distance

				if pp.used == 0 {
					emptyPoint = p
					break
				}
			}
		}
		distance++
	}

	track := []Node{grid[emptyPoint.y][emptyPoint.x]}

	lastPoint := emptyPoint

	currentDist := tracks[lastPoint]
	for currentDist > 0 {
		for _, neighbor := range neighbors(lastPoint) {
			if d, ok := tracks[neighbor]; ok && d == currentDist-1 {
				lastPoint = neighbor
				track = append(track, grid[lastPoint.y][lastPoint.x])
				currentDist--
				break
			}
		}
	}

	return track
}

func moveData(grid [][]Node, track []Node) {
	for i := range len(track) - 1 {
		curr := track[i]
		next := track[i+1]

		curr.used = next.used
		next.used = 0

		curr.available = curr.size - curr.used
		next.available = next.size

		grid[curr.y][curr.x] = curr
		grid[next.y][next.x] = next
	}
}

func part2(nodes []Node) {
	counter := 0

	grid := [][]Node{}

	maxX := 0
	maxY := 0
	for _, n := range nodes {
		if n.x > maxX {
			maxX = n.x
		}
		if n.y > maxY {
			maxY = n.y
		}
	}

	for range maxY + 1 {
		grid = append(grid, make([]Node, maxX+1))
	}

	for _, n := range nodes {
		grid[n.y][n.x] = n
	}

	currentX := maxX

	for currentX > 0 {
		toBeEmptySpot := grid[0][currentX-1]
		packPoint := grid[0][currentX]
		track := findEmpty(toBeEmptySpot, packPoint.Point, grid, maxX+1, maxY+1)

		counter += len(track)
		moveData(grid, track) // append(track, packPoint)
		moveData(grid, []Node{toBeEmptySpot, packPoint})

		currentX--
	}

	fmt.Printf("Part 2: %d\n", counter)
}
