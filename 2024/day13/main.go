package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"advent2024/util"
)

func main() {
	// data := util.GetTestByRow("day13")
	data := util.GetDataByRow("day13")

	list := parseMachines(data)

	part1(list)
	part2(list)
}

type Point struct {
	X int
	Y int
}

type Machine struct {
	ButtonA Point
	ButtonB Point
	Prize   Point
}

func (m Machine) ToString() string {
	return fmt.Sprintf("A[%d, %d] B[%d, %d] = Prize[%d, %d]", m.ButtonA.X, m.ButtonA.Y, m.ButtonB.X, m.ButtonB.Y, m.Prize.X, m.Prize.Y)
}

func parseMachines(data []string) []Machine {
	machines := []Machine{}

	m := Machine{}
	btnARegex, _ := regexp.Compile(`Button\sA:\s*X\+(\d+),\s*Y\+(\d+)`)
	btnBRegex, _ := regexp.Compile(`Button\sB:\s*X\+(\d+),\s*Y\+(\d+)`)
	prizeRegex, _ := regexp.Compile(`Prize:\s*X\=(\d+),\s*Y\=(\d+)`)

	for _, s := range data {
		if s == "" {
			machines = append(machines, m)
			m = Machine{}
			continue
		}

		matchA := btnARegex.FindStringSubmatch(s)
		if len(matchA) > 0 {
			x, _ := strconv.Atoi(matchA[1])
			y, _ := strconv.Atoi(matchA[2])
			m.ButtonA = Point{x, y}
			continue
		}

		matchB := btnBRegex.FindStringSubmatch(s)
		if len(matchB) > 0 {
			x, _ := strconv.Atoi(matchB[1])
			y, _ := strconv.Atoi(matchB[2])
			m.ButtonB = Point{x, y}
			continue
		}

		matchPrize := prizeRegex.FindStringSubmatch(s)
		if len(matchPrize) > 0 {
			x, _ := strconv.Atoi(matchPrize[1])
			y, _ := strconv.Atoi(matchPrize[2])
			m.Prize = Point{x, y}
			continue
		}
	}
	machines = append(machines, m)

	return machines
}

func checkInteger(n float64) bool {
	return math.Mod(n, 1) == 0
}

func part1(machines []Machine) {
	counter := 0

	for _, m := range machines {

		aPresses := float64(m.Prize.Y*m.ButtonB.X-m.Prize.X*m.ButtonB.Y) / float64(m.ButtonA.Y*m.ButtonB.X-m.ButtonA.X*m.ButtonB.Y)
		bPresses := (float64(m.Prize.X) - float64(m.ButtonA.X)*aPresses) / float64(m.ButtonB.X)

		if !checkInteger(aPresses) || !checkInteger(bPresses) {
			continue
		}

		sum := aPresses*3 + bPresses

		intSum := int(math.Round(sum))

		counter += intSum
	}
	fmt.Printf("Part 1: %d\n", counter)
}

func part2(machines []Machine) {
	counter := 0

	for _, m := range machines {

		yPrize := m.Prize.Y + 10000000000000
		xPrize := m.Prize.X + 10000000000000

		aPresses := float64(yPrize*m.ButtonB.X-xPrize*m.ButtonB.Y) / float64(m.ButtonA.Y*m.ButtonB.X-m.ButtonA.X*m.ButtonB.Y)
		bPresses := (float64(xPrize) - float64(m.ButtonA.X)*aPresses) / float64(m.ButtonB.X)

		if !checkInteger(aPresses) || !checkInteger(bPresses) {
			continue
		}

		sum := aPresses*3 + bPresses

		intSum := int(math.Round(sum))

		counter += intSum
	}
	fmt.Printf("Part 2: %d\n", counter)
}
