package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"advent2017/util"
)

func main() {
	data := util.GetDataByRow("day08")
	// data := util.GetTestByRow("day08")

	list := parse(data)

	part1(list)
	part2(list)
}

type operation struct {
	register string
	value    int
	left     string
	compare  string
	right    int
}

func parse(data []string) []operation {
	ops := []operation{}

	for _, s := range data {
		parts := strings.Split(s, " ")

		o := operation{}
		o.register = parts[0]
		val := util.ToInt(parts[2])
		if parts[1] == "dec" {
			val *= -1
		}
		o.value = val
		o.left = parts[4]
		o.compare = parts[5]
		o.right = util.ToInt(parts[6])
		ops = append(ops, o)
	}

	return ops
}

func execute(op operation, registers map[string]int) int {
	left := registers[op.left]

	result := false

	switch op.compare {
	case "<":
		result = left < op.right
	case "<=":
		result = left <= op.right
	case "==":
		result = left == op.right
	case ">=":
		result = left >= op.right
	case ">":
		result = left > op.right
	case "!=":
		result = left != op.right
	}

	if result {
		registers[op.register] += op.value
	}
	return registers[op.register]
}

func part1(data []operation) {
	counter := 0

	registers := map[string]int{}

	for _, o := range data {
		execute(o, registers)
	}

	counter = slices.Max(slices.Collect(maps.Values(registers)))
	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data []operation) {
	counter := 0

	registers := map[string]int{}

	for _, o := range data {
		v := execute(o, registers)
		if v > counter {
			counter = v
		}
	}

	fmt.Printf("Part 2: %d\n", counter)
}
