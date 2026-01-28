package main

import (
	"fmt"
	"math"
	"strings"

	"advent2024/util"
)

func main() {
	// data := util.GetTestByRow("day15")
	data := util.GetDataByRow("day15")

	problem := parse(data)
	problem2 := parse(data)

	part1(problem)
	part2(problem2)
}

type Direction int

const (
	UP    Direction = 1
	RIGHT Direction = 2
	DOWN  Direction = 3
	LEFT  Direction = 4
)

type Problem struct {
	robotX int
	robotY int
	grid   []string
	moves  []Direction
}

type Point struct {
	x int
	y int
}

func parse(data []string) Problem {
	doingMap := true

	grid := []string{}
	moves := []Direction{}

	var robotX, robotY int

	for y, r := range data {
		if r == "" {
			doingMap = false
			continue
		}
		if doingMap {
			grid = append(grid, r)
			if idx := strings.Index(r, "@"); idx != -1 {
				robotX = idx
				robotY = y
			}
		} else {
			for _, c := range r {
				switch c {
				case '^':
					moves = append(moves, UP)
				case '>':
					moves = append(moves, RIGHT)
				case 'v':
					moves = append(moves, DOWN)
				case '<':
					moves = append(moves, LEFT)
				}
			}
		}
	}

	return Problem{robotX, robotY, grid, moves}
}

func isFree(grid []string, x int, y int) bool {
	return grid[y][x] == '.'
}

func next(x int, y int, dir Direction) (int, int) {
	newX, newY := x, y
	switch dir {
	case UP:
		newY--
	case DOWN:
		newY++
	case RIGHT:
		newX++
	case LEFT:
		newX--
	}
	return newX, newY
}

func findNextFree(grid []string, x int, y int, direction Direction) (freeX int, freeY int, found bool) {
	freeX = x
	freeY = y

	for range len(grid) + len(grid[0]) {
		freeX, freeY = next(freeX, freeY, direction)
		if freeX < 0 || freeY < 0 || freeX >= len(grid[0]) || freeY >= len(grid) {
			return x, y, false
		}
		if grid[freeY][freeX] == '#' {
			return x, y, false
		}
		if isFree(grid, freeX, freeY) {
			return freeX, freeY, true
		}
	}
	return x, y, false
}

func writeChar(grid []string, x int, y int, c string) {
	grid[y] = grid[y][:x] + c + grid[y][x+1:]
}

func moveRobot(grid []string, x int, y int, direction Direction) (newX int, newY int) {
	emptyX, emptyY, found := findNextFree(grid, x, y, direction)

	newX, newY = x, y
	if found {
		newX, newY = next(newX, newY, direction)

		if math.Abs(float64(x-emptyX))+math.Abs(float64(y-emptyY)) > 1 {
			writeChar(grid, emptyX, emptyY, string(grid[newY][newX]))
		}
		writeChar(grid, newX, newY, "@")
		writeChar(grid, x, y, ".")
	}
	return
}

type WideBox struct {
	left  Point
	right Point
}

func moveOnWide(grid []string, x int, y int, direction Direction) (freeX int, freeY int) {
	toCheck := []Point{
		{
			x, y,
		},
	}
	switch direction {
	case UP:
		toCheck[0].y--
	case DOWN:
		toCheck[0].y++
	case RIGHT:
		toCheck[0].x++
	case LEFT:
		toCheck[0].x--
	}

	var p Point
	first := true
	for len(toCheck) > 0 {
		p, toCheck = toCheck[0], toCheck[1:]

		if isFree(grid, p.x, p.y) {
			continue
		}
		if grid[p.y][p.x] == '#' {
			return x, y
		}

		tX, tY := next(p.x, p.y, direction)
		toCheck = append(toCheck, Point{tX, tY})
		if direction == UP || direction == DOWN {
			hDir := RIGHT
			if grid[p.y][p.x] == ']' {
				hDir = LEFT
			}
			tX, tY = next(tX, tY, hDir)
			toCheck = append(toCheck, Point{tX, tY})
		}

		first = false
	}
	if first {
		writeChar(grid, p.x, p.y, "@")
		writeChar(grid, x, y, ".")
		return p.x, p.y
	}
	if direction == LEFT || direction == RIGHT {
		opposite := LEFT
		if direction == LEFT {
			opposite = RIGHT
		}
		oldX, oldY := p.x, p.y
		nextX, nextY := next(p.x, p.y, opposite)
		for nextX != x || nextY != y {
			writeChar(grid, oldX, oldY, string(grid[nextY][nextX]))
			oldX, oldY = nextX, nextY
			nextX, nextY = next(nextX, nextY, opposite)
		}
		writeChar(grid, oldX, oldY, "@")
		writeChar(grid, x, y, ".")
		return oldX, oldY
	}

	nextX, nextY := next(x, y, direction)
	boxes := []WideBox{}
	if grid[nextY][nextX] == '[' {
		boxes = append(boxes, WideBox{
			Point{nextX, nextY},
			Point{nextX + 1, nextY},
		})
		writeChar(grid, x, y, ".")
		writeChar(grid, nextX, nextY, "@")
		writeChar(grid, nextX+1, nextY, ".")
	} else {
		boxes = append(boxes, WideBox{
			Point{nextX - 1, nextY},
			Point{nextX, nextY},
		})
		writeChar(grid, x, y, ".")
		writeChar(grid, nextX, nextY, "@")
		writeChar(grid, nextX-1, nextY, ".")
	}
	var box WideBox
	for len(boxes) > 0 {
		box, boxes = boxes[0], boxes[1:]

		nextLeftX, nextLeftY := next(box.left.x, box.left.y, direction)
		nextRightX, nextRightY := next(box.right.x, box.right.y, direction)
		if grid[nextLeftY][nextLeftX] == '[' {
			// only one box in the exact position
			boxes = append(boxes, WideBox{
				Point{nextLeftX, nextLeftY},
				Point{nextRightX, nextRightY},
			})
			continue
		}
		if grid[nextLeftY][nextLeftX] == ']' {
			boxes = append(boxes, WideBox{
				Point{nextLeftX - 1, nextLeftY},
				Point{nextLeftX, nextLeftY},
			})
			writeChar(grid, nextLeftX-1, nextLeftY, ".")
		}
		if grid[nextRightY][nextRightX] == '[' {
			boxes = append(boxes, WideBox{
				Point{nextRightX, nextRightY},
				Point{nextRightX + 1, nextRightY},
			})
			writeChar(grid, nextRightX+1, nextRightY, ".")
		}
		writeChar(grid, nextLeftX, nextLeftY, "[")
		writeChar(grid, nextRightX, nextRightY, "]")
	}
	return nextX, nextY
}

func printGrid(grid []string) {
	fmt.Println("------")
	for _, s := range grid {
		fmt.Println(s)
	}
	fmt.Println("------")
}

func part1(problem Problem) {
	counter := 0

	for _, dir := range problem.moves {
		problem.robotX, problem.robotY = moveRobot(problem.grid, problem.robotX, problem.robotY, dir)
	}

	for y := range problem.grid {
		for x := range problem.grid[0] {
			if problem.grid[y][x] == 'O' {
				counter += y*100 + x
			}
		}
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func enlarge(grid []string) []string {
	newGrid := []string{}

	translations := map[rune]string{
		'#': "##",
		'O': "[]",
		'.': "..",
		'@': "@.",
	}

	for _, s := range grid {
		line := ""
		for _, c := range s {
			line += translations[c]
		}
		newGrid = append(newGrid, line)
	}

	return newGrid
}

func part2(problem Problem) {
	counter := 0

	problem.grid = enlarge(problem.grid)

	// find new robot position
	for y, r := range problem.grid {
		if x := strings.Index(r, "@"); x != -1 {
			problem.robotX = x
			problem.robotY = y
			break
		}
	}
	// più complesso del previsto: in orizontale, posso controllare solo la linea,
	// ma in verticale devo controllare che se sposto una scatola, l'altra metà
	// non sia bloccata, o che non ci sia una catena di scatole che impediscono il movimento, es:
	// ............
	// .##.........
	// ..[]........
	// ...[].......
	// ....[]......
	// .....[].....
	// ......@.....
	// ............

	for _, dir := range problem.moves {
		problem.robotX, problem.robotY = moveOnWide(problem.grid, problem.robotX, problem.robotY, dir)
	}

	for y := range problem.grid {
		for x := range problem.grid[0] {
			if problem.grid[y][x] == '[' {
				counter += y*100 + x
			}
		}
	}

	fmt.Printf("Part 2: %d\n", counter)
}
