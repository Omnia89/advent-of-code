package main

import (
	"fmt"
	"strings"

	"advent2024/util"
)

func main() {
	// data := util.GetTestByRow("day08")
	data := util.GetDataByRow("day08") // 912 too low

	part1(data)
	part2(data)
}

type Point struct {
	X int
	Y int
}

func (p Point) ToString() string {
	return fmt.Sprintf("%d-%d", p.X, p.Y)
}

func printTable(antennas map[string][]Point, points []Point, h int, w int) {
	table := []string{}

	for range h {
		table = append(table, strings.Repeat(".", w))
	}

	for _, p := range points {
		table[p.Y] = table[p.Y][:p.X] + "#" + table[p.Y][p.X+1:]
	}

	for k, ps := range antennas {
		for _, p := range ps {
			table[p.Y] = table[p.Y][:p.X] + k + table[p.Y][p.X+1:]
		}
	}

	for _, s := range table {
		fmt.Printf("%s\n", s)
	}
}

func part1(data []string) {
	sum := 0

	antennas := map[string][]Point{}

	for y := range data {
		for x := range data[0] {
			char := string(data[x][y])
			if char != "." {
				if _, ok := antennas[char]; !ok {
					antennas[char] = []Point{}
				}
				antennas[char] = append(antennas[char], Point{x, y})
			}
		}
	}

	antinodes := map[string]bool{}

	for _, points := range antennas {
		for i := range len(points) - 1 {
			for j := i + 1; j < len(points); j++ {
				deltaX := points[i].X - points[j].X
				deltaY := points[i].Y - points[j].Y
				a1 := Point{
					X: points[i].X + deltaX,
					Y: points[i].Y + deltaY,
				}
				a2 := Point{
					X: points[j].X - deltaX,
					Y: points[j].Y - deltaY,
				}

				if a1.X >= 0 && a1.X < len(data[0]) && a1.Y >= 0 && a1.Y < len(data) {
					antinodes[a1.ToString()] = true
					// fmt.Printf(" - [%s]\n", a1.ToString())
				}
				if a2.X >= 0 && a2.X < len(data[0]) && a2.Y >= 0 && a2.Y < len(data) {
					antinodes[a2.ToString()] = true
					// fmt.Printf(" - [%s]\n", a2.ToString())
				}
			}
		}
	}
	sum = len(antinodes)

	fmt.Printf("Part 1: %d\n", sum)
}

func part2(data []string) {
	sum := 0

	antennas := map[string][]Point{}

	antinodes := map[string]int{}
	for y := range data {
		for x := range data[0] {
			char := string(data[y][x])
			if char != "." {
				if _, ok := antennas[char]; !ok {
					antennas[char] = []Point{}
				}
				p := Point{x, y}
				antennas[char] = append(antennas[char], p)
				antinodes[p.ToString()] = 1
			}
		}
	}

	antennaPoint := []Point{}

	for _, points := range antennas {
		for i := range len(points) - 1 {
			for j := i + 1; j < len(points); j++ {
				deltaX := points[i].X - points[j].X
				deltaY := points[i].Y - points[j].Y

				// first direction
				a1 := points[i]
				for {
					a1 = Point{
						X: a1.X + deltaX,
						Y: a1.Y + deltaY,
					}
					if a1.X < 0 || a1.X >= len(data[0]) || a1.Y < 0 || a1.Y >= len(data) {
						break
					}
					antinodes[a1.ToString()]++
					antennaPoint = append(antennaPoint, a1)
					// fmt.Printf(" - [%s]\n", a1.ToString())
				}

				a2 := points[j]
				for {
					a2 = Point{
						X: a2.X - deltaX,
						Y: a2.Y - deltaY,
					}

					if a2.X < 0 || a2.X >= len(data[0]) || a2.Y < 0 || a2.Y >= len(data) {
						break
					}
					antinodes[a2.ToString()]++
					antennaPoint = append(antennaPoint, a2)
					// fmt.Printf(" - [%s]\n", a2.ToString())
				}
			}
		}
	}
	sum = len(antinodes)
	// for _, ps := range antennas {
	// 	sum += len(ps)
	// }
	// printTable(antennas, antennaPoint, len(data), len(data[0]))

	fmt.Printf("Part 2: %d\n", sum)
}
