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

func moveOnDPad(start Point, destination string) (moves string, end Point) {
	pad := map[string]Point{
		"^": {1, 0},
		"A": {2, 0},
		"<": {0, 1},
		"v": {1, 1},
		">": {2, 1},
	}
	inversePad := map[Point]string{
		{1, 0}: "^",
		{2, 0}: "A",
		{0, 1}: "<",
		{1, 1}: "v",
		{2, 1}: ">",
	}

	startSymbol, ok := inversePad[start]
	if !ok {
		return
	}

	destinationPoint, ok := pad[destination]
	if !ok {
		return
	}

	return numPadSequences[startSymbol][destination], destinationPoint
}

func getNumber(val string) int {
	v, _ := strconv.Atoi(strings.Replace(val, "A", "", -1))
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

		for range 2 {
			tempMoves := ""
			tempPoint := Point{2, 0}
			var m string

			for _, c := range moves {
				m, tempPoint = moveOnDPad(tempPoint, string(c))
				tempMoves += m
			}
			moves = tempMoves
		}
		calc := getNumber(sequence) * len(moves)
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

		for range 25 {
			tempMoves := ""
			tempPoint := Point{2, 0}
			var m string

			for _, c := range moves {
				m, tempPoint = moveOnDPad(tempPoint, string(c))
				tempMoves += m
			}
			moves = tempMoves
		}
		calc := getNumber(sequence) * len(moves)
		counter += calc
	}
	fmt.Printf("Part 2: %d\n", counter)
}
