package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strings"

	"advent2018/util"
)

func main() {
	// data := util.GetTestByRow("day17")
	data := util.GetDataByRow("day17")

	area, minY, maxY := parse(data)

	part1_2(area, minY, maxY)
}

type Point struct {
	x int
	y int
}

type Ground int

const (
	Sand Ground = iota
	Spring
	Clay
	Wet
	Water
)

func parse(data []string) (map[Point]Ground, int, int) {
	grounds := map[Point]Ground{
		{500, 0}: Spring,
	}

	minY := math.MaxInt
	maxY := 0

	regMatch := regexp.MustCompile(`^([xy])=(\d+),\s[xy]=(\d+)\.\.(\d+)$`)
	for _, s := range data {
		parts := regMatch.FindStringSubmatch(s)
		if parts[1] == "x" {
			x := util.ToInt(parts[2])
			l := util.ToInt(parts[3])
			u := util.ToInt(parts[4])
			for i := range u - l + 1 {
				grounds[Point{x, l + i}] = Clay
				if l+i < minY {
					minY = l + i
				}
				if l+i > maxY {
					maxY = l + i
				}
			}
		} else {
			y := util.ToInt(parts[2])
			l := util.ToInt(parts[3])
			u := util.ToInt(parts[4])
			for i := range u - l + 1 {
				grounds[Point{l + i, y}] = Clay
			}
			if y < minY {
				minY = y
			}
			if y > maxY {
				maxY = y
			}
		}
	}

	return grounds, minY, maxY
}

func printAreaFile(area map[Point]Ground) {
	minX := math.MaxInt
	maxX := 0
	minY := math.MaxInt
	maxY := 0

	for p := range area {
		if p.x < minX {
			minX = p.x
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	printableArea := make([][]string, maxY+1)
	for y := range maxY + 1 {
		row := slices.Repeat([]string{"."}, maxX-minX+1)
		printableArea[y] = row
	}

	for p, g := range area {
		var s string
		switch g {
		case Water:
			s = "~"
		case Wet:
			s = "|"
		case Spring:
			s = "+"
		case Clay:
			s = "#"
		default:
			s = "."
		}
		printableArea[p.y][p.x-minX] = s
	}

	var sb strings.Builder
	for _, r := range printableArea {
		sb.WriteString(strings.Join(r, ""))
		sb.WriteString("\n")
	}
	f, err := os.OpenFile("out.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString(sb.String())
}

func part1_2(area map[Point]Ground, minY int, maxY int) {
	counter := 0

	queue := make(PointQueue, 0)
	heap.Init(&queue)

	queue.Push(&Point{500, 0})

	for queue.Len() > 0 {
		p := queue.Pop().(*Point)
		if p.y == maxY {
			area[*p] = Wet
			continue
		}

		downPoint := Point{p.x, p.y + 1}
		down, ok := area[downPoint]
		if !ok || down == Sand {
			queue.Push(p)
			queue.Push(&downPoint)
			continue
		}

		if down == Clay || down == Water {
			leftWall := false
			rightWall := false

			// search left
			leftPoint := Point{p.x, p.y}
			for {
				leftPoint.x--
				if area[leftPoint] == Clay {
					leftWall = true
					leftPoint.x++
					break
				}
				if area[Point{leftPoint.x, leftPoint.y + 1}] == Sand || area[Point{leftPoint.x, leftPoint.y + 1}] == Wet {
					break
				}
			}
			// search right
			rightPoint := Point{p.x, p.y}
			for {
				rightPoint.x++
				if area[rightPoint] == Clay {
					rightPoint.x--
					rightWall = true
					break
				}
				if area[Point{rightPoint.x, rightPoint.y + 1}] == Sand || area[Point{rightPoint.x, rightPoint.y + 1}] == Wet {
					break
				}
			}
			// Fill with water
			if leftWall && rightWall {
				for i := range rightPoint.x - leftPoint.x + 1 {
					area[Point{leftPoint.x + i, p.y}] = Water
				}
			} else if rightWall {
				for i := range rightPoint.x - leftPoint.x + 1 {
					area[Point{leftPoint.x + i, p.y}] = Wet
				}
				//queue.Push(&Point{leftPoint.x, p.y + 1})
				queue.Push(&leftPoint)
			} else if leftWall {
				for i := range rightPoint.x - leftPoint.x + 1 {
					area[Point{leftPoint.x + i, p.y}] = Wet
				}
				//queue.Push(&Point{rightPoint.x, p.y + 1})
				queue.Push(&rightPoint)
			} else {
				// go both ways
				for i := range rightPoint.x - leftPoint.x + 1 {
					area[Point{leftPoint.x + i, p.y}] = Wet
				}
				//queue.Push(&Point{leftPoint.x, p.y + 1})
				//queue.Push(&Point{rightPoint.x, p.y + 1})
				queue.Push(&leftPoint)
				queue.Push(&rightPoint)
			}
		}

		if down == Wet {
			area[*p] = Wet
		}
	}

	water := 0
	for p, g := range area {
		if p.y < minY || p.y > maxY {
			continue
		}
		if g == Water || g == Wet {
			counter++
		}
		if g == Water {
			water++
		}
	}

	// printAreaFile(area)

	fmt.Printf("Part 1: %d\n", counter)
	fmt.Printf("Part 2: %d\n", water)
}
