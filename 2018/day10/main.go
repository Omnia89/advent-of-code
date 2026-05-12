package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"regexp"

	"advent2018/util"
)

func main() {
	data := util.GetDataByRow("day10")
	// data := util.GetTestByRow("day10")

	list := parse(data)
	list2 := parse(data)

	part1(list)
	part2(list2)
}

type Point struct {
	x int
	y int
}

type Light struct {
	p   Point
	dir Point
}

func parse(data []string) []Light {
	ls := []Light{}

	reg := regexp.MustCompile(`position=<\s*(-?\d+),\s*(-?\d+)>.*<\s*(-?\d+),\s*(-?\d+)>`)

	for _, s := range data {
		matchs := reg.FindStringSubmatch(s)
		ls = append(ls, Light{
			Point{util.ToInt(matchs[1]), util.ToInt(matchs[2])},
			Point{util.ToInt(matchs[3]), util.ToInt(matchs[4])},
		})
	}

	return ls
}

func printGrid(w io.Writer, ls []Light) {
	minX, minY, maxX, maxY := ls[0].p.x, ls[0].p.y, ls[0].p.x, ls[0].p.y

	for _, l := range ls {
		if l.p.x > maxX {
			maxX = l.p.x
		}
		if l.p.y > maxY {
			maxY = l.p.y
		}

		if l.p.x < minX {
			minX = l.p.x
		}
		if l.p.y < minY {
			minY = l.p.y
		}
	}

	width := maxX - minX + 1
	heigth := maxY - minY + 1

	grid := make([][]int, heigth)
	for y := range heigth {
		grid[y] = make([]int, width)
	}

	for _, l := range ls {
		grid[l.p.y-minY][l.p.x-minX] = 1
	}

	for _, r := range grid {
		for _, v := range r {
			s := " "
			if v == 1 {
				s = "#"
			}
			fmt.Fprintf(w, "%s", s)
		}
		//l := strings.Join(r, "")
		fmt.Fprintf(w, "\n")
	}
}

func saveFile(ls []Light) {
	basePath := "./day10/out"
	os.MkdirAll(basePath, 0o755)

	path := filepath.Join(basePath, "out.txt")

	f, _ := os.Create(path)
	defer f.Close()

	printGrid(f, ls)
}

func step(ls []Light) {
	for i := range ls {
		p := ls[i].p
		p.x += ls[i].dir.x
		p.y += ls[i].dir.y
		ls[i].p = p
	}
}

func revStep(ls []Light) {
	for i := range ls {
		p := ls[i].p
		p.x -= ls[i].dir.x
		p.y -= ls[i].dir.y
		ls[i].p = p
	}
}

func getArea(ls []Light) int {
	minX, minY, maxX, maxY := ls[0].p.x, ls[0].p.y, ls[0].p.x, ls[0].p.y

	for _, l := range ls {
		if l.p.x > maxX {
			maxX = l.p.x
		}
		if l.p.y > maxY {
			maxY = l.p.y
		}

		if l.p.x < minX {
			minX = l.p.x
		}
		if l.p.y < minY {
			minY = l.p.y
		}
	}

	width := maxX - minX + 1
	heigth := maxY - minY + 1
	return width * heigth
}

func part1(data []Light) {
	lastBox := math.MaxInt

	for {
		step(data)
		area := getArea(data)
		if area > lastBox {
			revStep(data)
			saveFile(data)
			break
		}
		lastBox = area
	}

	fmt.Printf("Part 1: -read out/out.txt-\n")
}

func part2(data []Light) {
	counter := 0

	lastBox := math.MaxInt

	for {
		step(data)
		counter++
		area := getArea(data)
		if area > lastBox {
			counter--
			break
		}
		lastBox = area
	}

	fmt.Printf("Part 2: %d\n", counter)
}
