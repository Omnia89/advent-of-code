package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"advent2024/util"
)

func main() {
	// data := util.GetTestByRow("day14")
	// height, width := 7, 11
	data := util.GetDataByRow("day14") // pt1 216145083 too low
	height, width := 103, 101

	list := parse(data)

	part1(list, height, width)
	part2(list, height, width)
}

type Point struct {
	X int
	Y int
}

type Robot struct {
	Position Point
	Movement Point
}

func (r *Robot) Move(times int, width int, height int) {
	if times <= 0 {
		return
	}
	newX := (r.Position.X + times*r.Movement.X) % width
	newY := (r.Position.Y + times*r.Movement.Y) % height

	for newX < 0 {
		newX += width
	}
	for newY < 0 {
		newY += height
	}

	r.Position = Point{newX, newY}
}

func (r *Robot) MoveOnTheGrid(grid [][]int, width int, height int) [][]int {
	oldX, oldY := r.Position.X, r.Position.Y

	r.Move(1, width, height)

	grid[oldY][oldX] = 0
	grid[r.Position.Y][r.Position.X] = 1
	return grid
}

func (r *Robot) Quadrant(width, height int) int {
	halfW := width / 2
	halfH := height / 2

	if r.Position.Y < halfH {
		if r.Position.X < halfW {
			return 1
		} else if r.Position.X > halfW {
			return 2
		}
	} else if r.Position.Y > halfH {
		if r.Position.X < halfW {
			return 3
		} else if r.Position.X > halfW {
			return 4
		}
	}

	// on the line, no quadrant
	return 0
}

func parse(data []string) []Robot {
	robots := []Robot{}

	reg, _ := regexp.Compile(`p=(\d+),(\d+)\s+v=([0-9\-]+),([0-9\-]+)`)

	for i, s := range data {
		values := reg.FindStringSubmatch(s)
		if len(values) == 0 {
			fmt.Printf("====== ERROR LINE [%d] ======", i)
			continue
		}
		robots = append(robots, Robot{Point{util.ToInt(values[1]), util.ToInt(values[2])}, Point{util.ToInt(values[3]), util.ToInt(values[4])}})
	}

	return robots
}

func (m Robot) ToString() string {
	return fmt.Sprintf("Pos[%d, %d] Mov[%d, %d]", m.Position.X, m.Position.Y, m.Movement.X, m.Movement.Y)
}

func part1(robots []Robot, h int, w int) {
	counter := 1

	quads := map[int]int{}

	for _, r := range robots {
		r.Move(100, w, h)
		if q := r.Quadrant(w, h); q != 0 {
			quads[q]++
		}
		// fmt.Printf("  %s\n", r.ToString())
	}

	for _, q := range quads {
		counter *= q
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func getGrid(robots []Robot, h int, w int) [][]int {
	grid := [][]int{}

	for range h {
		grid = append(grid, make([]int, w))
	}

	for _, r := range robots {
		grid[r.Position.Y][r.Position.X] = 1
	}

	return grid
}

func sumLine(line []int) int {
	c := 0
	for _, l := range line {
		c += l
	}
	return c
}

func countConsecutives(line []int) int {
	max := 0
	c := 0
	for _, n := range line {
		if n == 1 {
			c++
		} else {
			if c > max {
				max = c
			}
			c = 0
		}
	}
	return max
}

func checkTree(grid [][]int) bool {
	for _, l := range grid {
		n := countConsecutives(l)
		if n == len("###############################") {
			return true
		}
	}
	return false
}

func printLine(line []int) string {
	str := ""
	for _, l := range line {
		if l == 0 {
			str += " "
		} else {
			str += "#"
		}
	}
	return str
}

func saveTree(grid [][]int) {
	filename := "tree.txt"
	f, _ := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o600)
	defer f.Close()
	w := bufio.NewWriter(f)

	for _, l := range grid {
		fmt.Fprintf(w, "%s\n", printLine(l))
	}
	fmt.Fprintf(w, "---\n")
	w.Flush()
}

func part2(robots []Robot, h int, w int) {
	counter := 0

	grid := getGrid(robots, h, w)

	for n := range 100000 {
		for i := range robots {
			grid = robots[i].MoveOnTheGrid(grid, w, h)
		}
		// saveTree(grid)
		if checkTree(grid) {
			counter = n + 1
			break
		}
	}

	fmt.Printf("Part 2: %d\n", counter)
}
