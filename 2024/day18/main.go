package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"advent2024/util"
)

func main() {
	//data := util.GetTestByRow("day18")
	//size := 6
	//falling := 12

	data := util.GetDataByRow("day18")
	size := 70
	falling := 1024

	list := parse(data, size, falling)

	part1(list)
	part2(list)
}

type Point struct {
	X int
	Y int
}

type Problem struct {
	bytes        []Point
	size         int
	fallingBytes int
}

func parse(data []string, size int, falling int) Problem {
	problem := Problem{
		make([]Point, 0, len(data)),
		size,
		falling,
	}

	for _, r := range data {
		val := util.StringToIntSlice(r, ",")
		problem.bytes = append(problem.bytes, Point{val[0], val[1]})
	}

	return problem
}

func writeChar(grid []string, x int, y int, c string) {
	grid[y] = grid[y][:x] + c + grid[y][x+1:]
}

func printMap(data []string, points []Point, start Point, end Point) {
	var cloned []string
	cloned = append(cloned, data...)

	for _, p := range points {
		writeChar(cloned, p.X, p.Y, "+")
	}

	writeChar(cloned, start.X, start.Y, "S")
	writeChar(cloned, end.X, end.Y, "E")

	fmt.Println("---------------------------------------")
	for _, s := range cloned {
		fmt.Println(s)
	}
	fmt.Println("---------------------------------------")
}

func createMap(problem Problem, fallingBytes int) []string {
	grid := make([]string, 0, problem.size+1)

	for range problem.size + 1 {
		grid = append(grid, strings.Repeat(".", problem.size+1))
	}

	for i := range fallingBytes {
		p := problem.bytes[i]
		writeChar(grid, p.X, p.Y, "#")
	}
	return grid
}

func getNear(grid []string, p Point) []Point {
	ps := []Point{}

	if p.X-1 >= 0 && grid[p.Y][p.X-1] != '#' {
		ps = append(ps, Point{p.X - 1, p.Y})
	}

	if p.Y-1 >= 0 && grid[p.Y-1][p.X] != '#' {
		ps = append(ps, Point{p.X, p.Y - 1})
	}

	if p.X+1 < len(grid[0]) && grid[p.Y][p.X+1] != '#' {
		ps = append(ps, Point{p.X + 1, p.Y})
	}
	if p.Y+1 < len(grid) && grid[p.Y+1][p.X] != '#' {
		ps = append(ps, Point{p.X, p.Y + 1})
	}

	return ps
}

func aStar(grid []string, start Point, end Point) ([]Point, bool) {
	table := map[Point]int{
		start: 0,
	}

	queue := []Point{start}
	var p Point
	for len(queue) > 0 {
		p, queue = queue[0], queue[1:]
		cost := table[p] + 1

		if p == end {
			break
		}

		ps := getNear(grid, p)

		for _, pp := range ps {
			if v, ok := table[pp]; !ok || v > cost {
				table[pp] = cost
				if !ok {
					queue = append(queue, pp)
				}
			}
		}
	}

	path := []Point{}
	last := end
	for range 70 * 70 {
		ps := getNear(grid, last)
		if len(ps) == 0 {
			break
		}

		var p Point
		cost := math.MaxInt
		for _, pp := range ps {
			if v, ok := table[pp]; ok && v < cost {
				cost = v
				p = pp
			}
		}
		if cost == math.MaxInt {
			break
		}
		path = append(path, p)
		last = p
		if p == start {
			break
		}
	}
	slices.Reverse(path)

	return path, last == start
}

func part1(problem Problem) {
	counter := 0

	grid := createMap(problem, problem.fallingBytes)

	start := Point{0, 0}
	end := Point{problem.size, problem.size}

	path, _ := aStar(grid, start, end)

	//printMap(grid, path, start, end)

	counter = len(path)

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(problem Problem) {
	counter := 0

	start := Point{0, 0}
	end := Point{problem.size, problem.size}

	bottom := 0
	upper := len(problem.bytes)

	newIndex := func(bottom, upper int) int {
		return (bottom + upper) / 2
	}

	for {
		toCheck := newIndex(bottom, upper)
		grid := createMap(problem, toCheck)

		_, found := aStar(grid, start, end)
		if found {
			// check after
			grid = createMap(problem, toCheck+1)
			_, foundAfter := aStar(grid, start, end)
			if !foundAfter {
				counter = toCheck
				break
			}

			bottom = toCheck
		} else {
			// check before
			grid = createMap(problem, toCheck-1)
			_, foundBefore := aStar(grid, start, end)
			if foundBefore {
				counter = toCheck - 1
				break
			}

			upper = toCheck
		}

	}
	result := fmt.Sprintf("%d,%d", problem.bytes[counter].X, problem.bytes[counter].Y)

	fmt.Printf("Part 2: %s\n", result)
}
