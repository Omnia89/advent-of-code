package main

import (
	"fmt"
	"math"
	"slices"

	"advent2016/util"
)

func main() {
	data := util.GetDataByRow("day24")
	// data := util.GetTestByRow("day24")

	p := parse(data)

	part1(p)
	part2(p)
}

type Point struct {
	x int
	y int
}

type Problem struct {
	grid  [][]int
	start Point
	poi   []Point
}

func parse(data []string) Problem {
	p := Problem{
		grid: [][]int{},
		poi:  []Point{},
	}

	for y, s := range data {
		p.grid = append(p.grid, []int{})
		for x, c := range s {
			switch c {
			case '.':
				p.grid[y] = append(p.grid[y], 1)
			case '#':
				p.grid[y] = append(p.grid[y], 0)
			case '0':
				p.grid[y] = append(p.grid[y], 1)
				p.start = Point{x, y}
			default:
				p.grid[y] = append(p.grid[y], 1)
				p.poi = append(p.poi, Point{x, y})
			}
		}
	}

	return p
}

func getNear(p Point) []Point {
	return []Point{
		{p.x, p.y - 1},
		{p.x, p.y + 1},
		{p.x - 1, p.y},
		{p.x + 1, p.y},
	}
}

func getDistances(start Point, grid [][]int, poi []Point) map[Point]int {
	dist := map[Point]int{}

	distances := map[Point]int{start: 0}

	q := []Point{start}
	var p Point

	for len(q) > 0 {
		p, q = q[0], q[1:]
		distance := distances[p] + 1

		if slices.Contains(poi, p) {
			dist[p] = distances[p]
		}

		for _, n := range getNear(p) {
			if n.x < 0 || n.y < 0 || n.x >= len(grid[0]) || n.y >= len(grid) {
				continue
			}
			_, ok := distances[n]
			if !ok && grid[n.y][n.x] == 1 {
				distances[n] = distance
				q = append(q, n)
			}
		}
	}

	return dist
}

func getAllDistances(problem Problem) map[Point]map[Point]int {
	distances := map[Point]map[Point]int{} // start - end - distance
	allPoints := append(problem.poi, problem.start)
	var p Point
	for len(allPoints) > 0 {
		p, allPoints = allPoints[0], allPoints[1:]

		d := getDistances(p, problem.grid, allPoints)

		for pp, v := range d {
			if _, o := distances[p]; !o {
				distances[p] = map[Point]int{}
			}
			distances[p][pp] = v

			if _, o := distances[pp]; !o {
				distances[pp] = map[Point]int{}
			}
			distances[pp][p] = v
		}
	}
	return distances
}

type state struct {
	visited []Point
	last    Point
	steps   int
}

func part1(problem Problem) {
	counter := 0

	// breadth first search per determinare tutte le distanze di tutti i punti e start tra loro
	distances := getAllDistances(problem)

	cmpState := func(a, b state) int {
		return a.steps - b.steps
	}

	// Dijkstra per trovare il percorso minore
	q := []state{{
		visited: []Point{problem.start},
		last:    problem.start,
		steps:   0,
	}}
	var s state

	minSteps := math.MaxInt

	for len(q) > 0 {
		s, q = q[0], q[1:]

		if s.steps > minSteps {
			continue
		}

		for p, v := range distances[s.last] {
			if slices.Contains(s.visited, p) {
				continue
			}
			visited := append(slices.Clone(s.visited), p)

			if len(visited) == len(problem.poi)+1 {
				if s.steps+v < minSteps {
					minSteps = s.steps + v
				}
				continue
			}

			q = append(q, state{
				visited: visited,
				last:    p,
				steps:   s.steps + v,
			})
		}
		slices.SortFunc(q, cmpState)
	}
	counter = minSteps

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(problem Problem) {
	counter := 0

	// breadth first search per determinare tutte le distanze di tutti i punti e start tra loro
	distances := getAllDistances(problem)

	cmpState := func(a, b state) int {
		return a.steps - b.steps
	}

	// Dijkstra per trovare il percorso minore
	q := []state{{
		visited: []Point{problem.start},
		last:    problem.start,
		steps:   0,
	}}
	var s state

	minSteps := math.MaxInt

	for len(q) > 0 {
		s, q = q[0], q[1:]

		if s.steps > minSteps {
			continue
		}

		for p, v := range distances[s.last] {
			if slices.Contains(s.visited, p) {
				continue
			}
			visited := append(slices.Clone(s.visited), p)

			if len(visited) == len(problem.poi)+1 {
				steps := s.steps + v + distances[p][problem.start]
				if steps < minSteps {
					minSteps = steps
				}
				continue
			}

			q = append(q, state{
				visited: visited,
				last:    p,
				steps:   s.steps + v,
			})
		}
		slices.SortFunc(q, cmpState)
	}
	counter = minSteps

	fmt.Printf("Part 2: %d\n", counter)
}
