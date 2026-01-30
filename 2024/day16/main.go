package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"advent2024/util"
)

func main() {
	// data := util.GetTestByRow("day16")
	data := util.GetDataByRow("day16")

	part1(data)
	part2(data)
}

type Point struct {
	X int
	Y int
}

func (p Point) toString() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

func locate(data []string, char string) Point {
	for y, r := range data {
		if x := strings.Index(r, char); x != -1 {
			return Point{x, y}
		}
	}
	return Point{}
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

type Direction int

const (
	NORTH Direction = 1
	EAST  Direction = 2
	SOUTH Direction = 3
	WEST  Direction = 4
)

func (d Direction) move(x, y int) (int, int) {
	switch d {
	case NORTH:
		y--
	case EAST:
		x++
	case SOUTH:
		y++
	case WEST:
		x--
	}
	return x, y
}

type state struct {
	Point
	direction Direction
}

func (s state) goStraight() state {
	x, y := s.direction.move(s.X, s.Y)

	return state{
		Point{x, y},
		s.direction,
	}
}

func (s state) turnLeft() state {
	var newDirection Direction

	switch s.direction {
	case NORTH:
		newDirection = WEST
	case EAST:
		newDirection = NORTH
	case SOUTH:
		newDirection = EAST
	case WEST:
		newDirection = SOUTH
	}

	x, y := newDirection.move(s.X, s.Y)

	return state{
		Point{x, y},
		newDirection,
	}
}

func (s state) turnRight() state {
	var newDirection Direction

	switch s.direction {
	case NORTH:
		newDirection = EAST
	case EAST:
		newDirection = SOUTH
	case SOUTH:
		newDirection = WEST
	case WEST:
		newDirection = NORTH
	}

	x, y := newDirection.move(s.X, s.Y)

	return state{
		Point{x, y},
		newDirection,
	}
}

func isOk(data []string, p Point) bool {
	if data[p.Y][p.X] == '#' {
		return false
	}
	return true
}

func cmpState(distances map[state]int) func(a, b state) int {
	return func(a, b state) int {
		return distances[a] - distances[b]
	}
}

func dijkstra(data []string, start Point, end Point, startDirection Direction) (path []Point, allPoints []Point) {
	startState := state{start, startDirection}

	distances := map[state]int{
		startState: 0,
	}
	parents := map[state][]state{
		startState: {},
	}

	queue := []state{startState}

	var s state
	for len(queue) > 0 {
		slices.SortFunc(queue, cmpState(distances))
		s, queue = queue[0], queue[1:]

		moves := []struct {
			st   state
			cost int
		}{
			{s.goStraight(), 1},
			{s.turnLeft(), 1001},
			{s.turnRight(), 1001},
		}

		for _, m := range moves {
			newCost := m.cost + distances[s]
			st := m.st
			if !isOk(data, st.Point) {
				continue
			}

			oldCost, ok := distances[st]
			if !ok || newCost < oldCost {
				distances[st] = newCost
				parents[st] = []state{s}
				queue = append(queue, st)
			} else if newCost == oldCost {
				// Tie: add parent
				parents[st] = append(parents[st], s)
			}
		}
	}

	path = []Point{}

	var endSt state
	cost := math.MaxInt
	for _, d := range []Direction{NORTH, EAST, SOUTH, WEST} {
		checkSt := state{end, d}
		if c, ok := distances[checkSt]; ok && c < cost {
			cost = c
			endSt = checkSt
		}
	}

	st := endSt
	for {
		path = append(path, st.Point)

		ppSt, ok := parents[st]
		if !ok {
			panic(fmt.Sprintf("Missing entry: p[%s]", st.toString()))
		}
		if len(ppSt) == 0 {
			break
		}

		st = ppSt[0]
	}

	allPointPath := map[Point]bool{
		endSt.Point: true,
	}
	st = endSt
	allPointQueue := []state{endSt}

	for len(allPointQueue) > 0 {
		st, allPointQueue = allPointQueue[0], allPointQueue[1:]

		ppSt, ok := parents[st]

		if !ok {
			panic(fmt.Sprintf("Missing entry: p[%s]", st.toString()))
		}
		for _, pp := range ppSt {
			allPointPath[pp.Point] = true

			allPointQueue = append(allPointQueue, pp)
		}
	}

	slices.Reverse(path)

	allPoints = []Point{}
	for k := range allPointPath {
		allPoints = append(allPoints, k)
	}
	return path, allPoints
}

func part1(data []string) {
	counter := 0

	start := locate(data, "S")
	end := locate(data, "E")

	direction := EAST

	path, _ := dijkstra(data, start, end, direction)

	// printMap(data, path, start, end)

	for i := 0; i < len(path)-1; i++ {
		p := path[i]
		np := path[i+1]

		checkP := p
		switch direction {
		case NORTH:
			checkP.Y--
		case EAST:
			checkP.X++
		case SOUTH:
			checkP.Y++
		case WEST:
			checkP.X--
		}

		counter += 1

		if checkP != np {
			counter += 1000
			switch {
			case p.X-np.X < 0:
				direction = EAST
			case p.X-np.X > 0:
				direction = WEST
			case p.Y-np.Y < 0:
				direction = SOUTH
			case p.Y-np.Y > 0:
				direction = NORTH
			}
		}

	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data []string) {
	counter := 0

	start := locate(data, "S")
	end := locate(data, "E")

	direction := EAST

	_, points := dijkstra(data, start, end, direction)

	// printMap(data, points, start, end)

	counter = len(points)
	fmt.Printf("Part 2: %d\n", counter)
}
