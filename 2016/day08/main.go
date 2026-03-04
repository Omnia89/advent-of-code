package main

import (
	"fmt"
	"strings"

	"advent2016/util"
)

func main() {
	data := util.GetDataByRow("day08")
	// data := util.GetTestByRow("day08")

	list := parse(data)

	part1(list)
	part2(list)
}

type operation struct {
	op   string
	val1 int
	val2 int
}

func parse(data []string) []operation {
	ops := []operation{}

	for _, s := range data {
		t := strings.ReplaceAll(s, "=", " ")
		parts := strings.Split(t, " ")

		o := operation{}

		if parts[0] == "rect" {
			o.op = parts[0]
			v1, v2, _ := strings.Cut(parts[1], "x")
			o.val1 = util.ToInt(v1)
			o.val2 = util.ToInt(v2)
		} else {
			o.op = parts[1]
			o.val1 = util.ToInt(parts[3])
			o.val2 = util.ToInt(parts[5])
		}
		ops = append(ops, o)
	}

	return ops
}

func makeScreen(w, h int) [][]int {
	screen := make([][]int, 0, h)
	for range h {
		screen = append(screen, make([]int, w))
	}
	return screen
}

func exec(screen [][]int, op operation) {
	if op.op == "rect" {
		for x := range op.val1 {
			for y := range op.val2 {
				screen[y][x] = 1
			}
		}
		return
	}

	if op.op == "row" {
		newRow := make([]int, len(screen[op.val1]))
		for x, val := range screen[op.val1] {
			newX := (x + op.val2) % len(screen[op.val1])
			newRow[newX] = val
		}
		screen[op.val1] = newRow
		return
	}

	// column
	newColumn := make([]int, len(screen))
	for y := range screen {
		val := screen[y][op.val1]
		newY := (y + op.val2) % len(screen)
		newColumn[newY] = val
	}
	for y, val := range newColumn {
		screen[y][op.val1] = val
	}
}

func printScreen(s [][]int) {
	fmt.Println("------------------")
	for _, r := range s {
		for _, v := range r {
			c := " "
			if v == 1 {
				c = "#"
			}
			fmt.Print(c)
		}
		fmt.Println()
	}
	fmt.Println("------------------")
}

func part1(operations []operation) {
	counter := 0

	screen := makeScreen(50, 6)

	for _, o := range operations {
		exec(screen, o)
	}
	// printScreen(screen)

	for _, r := range screen {
		for _, v := range r {
			counter += v
		}
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(operations []operation) {
	counter := 0

	screen := makeScreen(50, 6)

	for _, o := range operations {
		exec(screen, o)
	}
	printScreen(screen)

	fmt.Printf("Part 2: %d\n", counter)
}
