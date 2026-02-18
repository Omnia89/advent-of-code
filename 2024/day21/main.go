package main

import (
	"fmt"
	"strconv"
	"strings"

	"advent2024/util"
)

func main() {
	// data := util.GetTestByRow("day21")
	data := util.GetDataByRow("day21")

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

func moveOnNumberPad(start Point, destination string) (moves string, end Point) {
	pad := map[string]Point{
		"7": {0, 0},
		"8": {1, 0},
		"9": {2, 0},
		"4": {0, 1},
		"5": {1, 1},
		"6": {2, 1},
		"1": {0, 2},
		"2": {1, 2},
		"3": {2, 2},
		"0": {1, 3},
		"A": {2, 3},
	}

	destinationPoint, ok := pad[destination]
	if !ok {
		return
	}

	dx := destinationPoint.X - start.X
	dy := destinationPoint.Y - start.Y

	moves = ""

	// detect if passing through blank (bottom left) - take the less optimal order
	if start.X == 0 && destinationPoint.Y == 3 || start.Y == 3 && destinationPoint.X == 0 {

		if dy < 0 {
			moves += strings.Repeat("^", -dy)
		}
		if dx > 0 {
			moves += strings.Repeat(">", dx)
		}

		if dy > 0 {
			moves += strings.Repeat("v", dy)
		}
		if dx < 0 {
			moves += strings.Repeat("<", -dx)
		}
	} else {

		// optimal sequence
		if dx < 0 {
			moves += strings.Repeat("<", -dx)
		}
		if dy > 0 {
			moves += strings.Repeat("v", dy)
		}

		if dy < 0 {
			moves += strings.Repeat("^", -dy)
		}
		if dx > 0 {
			moves += strings.Repeat(">", dx)
		}
	}
	moves += "A"
	return moves, destinationPoint
}

// [start][end]sequence
var numPadSequences = map[string]map[string]string{
	"<": {
		"<": "A",
		">": ">>A",
		"^": ">^A",
		"v": ">A",
		"A": ">>^A",
	},
	">": {
		"<": "<<A",
		">": "A",
		"^": "<^A",
		"v": "<A",
		"A": "^A",
	},
	"^": {
		"<": "v<A",
		">": "v>A",
		"^": "A",
		"v": "vA",
		"A": ">A",
	},
	"v": {
		"<": "<A",
		">": ">A",
		"^": "^A",
		"v": "A",
		"A": "^>A",
	},
	"A": {
		"<": "v<<A",
		">": "vA",
		"^": "<A",
		"v": "<vA",
		"A": "A",
	},
}

type cacheKey struct {
	moves  string
	nRobot int
}

func recursiveDPad(moves string, nRobot int, cache map[cacheKey]int) int {
	if c, ok := cache[cacheKey{moves, nRobot}]; ok {
		return c
	}

	cost := 0
	for i := range len(moves) - 1 {
		start := moves[i : i+1]
		end := moves[i+1 : i+2]
		seq := numPadSequences[start][end]

		val := 0
		if nRobot == 1 {
			val = len(seq)
		} else {
			val = recursiveDPad("A"+seq, nRobot-1, cache)
		}

		cost += val
	}
	cache[cacheKey{moves, nRobot}] = cost
	return cost
}

func getCost(moves string, nRobot int) int {
	cache := map[cacheKey]int{}

	cost := recursiveDPad("A"+moves, nRobot, cache)

	return cost
}

func getNumber(val string) int {
	v, _ := strconv.Atoi(strings.ReplaceAll(val, "A", ""))
	return v
}

func part1(data []string) {
	counter := 0

	for _, sequence := range data {
		moves := ""

		point := Point{2, 3}

		var m string
		for _, c := range sequence {
			m, point = moveOnNumberPad(point, string(c))
			moves += m
		}

		cost := getCost(moves, 2)
		number := getNumber(sequence)
		calc := number * cost
		counter += calc
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data []string) {
	counter := 0

	for _, sequence := range data {
		moves := ""

		point := Point{2, 3}

		var m string
		for _, c := range sequence {
			m, point = moveOnNumberPad(point, string(c))
			moves += m
		}

		cost := getCost(moves, 25)
		number := getNumber(sequence)
		calc := number * cost
		counter += calc
	}
	fmt.Printf("Part 2: %d\n", counter)
}
