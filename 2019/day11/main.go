package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"advent2019/util"
)

func main() {
	data := util.GetRawData("day11")
	// data := util.GetRawTest("day11")

	list := util.StringToIntSlice(data, ",")

	part1(list)
	part2(list)
}

type Point struct {
	x int
	y int
}

var turnLeft = map[Point]Point{
	{0, -1}: {-1, 0},
	{-1, 0}: {0, 1},
	{0, 1}:  {1, 0},
	{1, 0}:  {0, -1},
}

var turnRight = map[Point]Point{
	{0, -1}: {1, 0},
	{1, 0}:  {0, 1},
	{0, 1}:  {-1, 0},
	{-1, 0}: {0, -1},
}

func runRobot(data []int, hull map[Point]int) {
	inChan := make(chan int, 3)
	outChan := make(chan int, 3)

	intcode := NewIntcodeChannel(data, inChan, outChan, false)

	robot := Point{0, 0}
	facing := Point{0, -1}

	go intcode.run(0)

	for {
		color := hull[robot]

		inChan <- color

		paint, ok := <-outChan
		if !ok {
			break
		}

		direction, ok := <-outChan
		if !ok {
			break
		}

		hull[robot] = paint
		if direction == 0 {
			facing = turnLeft[facing]
		} else {
			facing = turnRight[facing]
		}
		robot.x += facing.x
		robot.y += facing.y
	}
}

func part1(data []int) {
	counter := 0

	hull := map[Point]int{}

	runRobot(data, hull)

	counter = len(hull)

	fmt.Printf("Part 1: %d\n", counter)
}

func getMatrixString(values map[Point]int) string {
	minX, minY := math.MaxInt, math.MaxInt
	maxX, maxY := math.MinInt, math.MinInt

	for p, v := range values {
		if v == 0 {
			continue
		}
		if p.x < minX {
			minX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	maxX -= (minX - 1)
	maxY -= (minY - 1)

	grid := make([][]string, maxY)
	for i := range grid {
		grid[i] = make([]string, maxX)
		for j := range grid[i] {
			grid[i][j] = " "
		}
	}

	for p, v := range values {
		if v == 0 {
			continue
		}
		nx, ny := p.x-minX, p.y-minY

		grid[ny][nx] = "#"
	}

	var sb strings.Builder

	for _, s := range grid {
		sb.WriteString(strings.Join(s, ""))
		sb.WriteString("\n")
	}

	return sb.String()
}

func outputFile(val string) {
	f, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	f.WriteString(val)
	f.Sync()
}

func part2(data []int) {
	counter := 0

	hull := map[Point]int{{0, 0}: 1}

	runRobot(data, hull)

	outputFile(getMatrixString(hull))

	fmt.Printf("Part 2: %d\n", counter)
}
