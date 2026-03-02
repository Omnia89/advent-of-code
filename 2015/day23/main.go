package main

import (
	"fmt"
	"strings"

	"advent2015/util"
)

func main() {
	data := util.GetDataByRow("day23")
	// data := util.GetTestByRow("day23")

	list := parse(data)

	part1(list)
	part2(list)
}

type instruction struct {
	op       string
	register string
	offset   int
}

func parse(data []string) []instruction {
	inst := []instruction{}

	for _, s := range data {
		t := strings.ReplaceAll(s, ",", "")
		t = strings.ReplaceAll(t, "+", "")
		pieces := strings.Split(t, " ")

		i := instruction{}
		i.op = pieces[0]

		if pieces[0] == "jmp" {
			i.offset = util.ToInt(pieces[1])
		} else if pieces[0] == "jie" || pieces[0] == "jio" {
			i.register = pieces[1]
			i.offset = util.ToInt(pieces[2])
		} else {
			i.register = pieces[1]
		}
		inst = append(inst, i)
	}

	return inst
}

// 20895 too high
func part1(data []instruction) {
	counter := 0

	pointer := 0
	registry := map[string]int{
		"a": 0,
		"b": 0,
	}

	for pointer >= 0 && pointer < len(data) {
		inst := data[pointer]
		switch inst.op {
		case "hlf":
			registry[inst.register] /= 2
			pointer++
		case "tpl":
			registry[inst.register] *= 3
			pointer++
		case "inc":
			registry[inst.register] += 1
			pointer++
		case "jmp":
			pointer += inst.offset
		case "jie":
			if registry[inst.register]%2 == 0 {
				pointer += inst.offset
			} else {
				pointer++
			}
		case "jio":
			if registry[inst.register] == 1 {
				pointer += inst.offset
			} else {
				pointer++
			}
		}
	}

	counter = registry["b"]
	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data []instruction) {
	counter := 0
	pointer := 0
	registry := map[string]int{
		"a": 1,
		"b": 0,
	}

	for pointer >= 0 && pointer < len(data) {
		inst := data[pointer]
		switch inst.op {
		case "hlf":
			registry[inst.register] /= 2
			pointer++
		case "tpl":
			registry[inst.register] *= 3
			pointer++
		case "inc":
			registry[inst.register] += 1
			pointer++
		case "jmp":
			pointer += inst.offset
		case "jie":
			if registry[inst.register]%2 == 0 {
				pointer += inst.offset
			} else {
				pointer++
			}
		case "jio":
			if registry[inst.register] == 1 {
				pointer += inst.offset
			} else {
				pointer++
			}
		}
	}

	counter = registry["b"]
	fmt.Printf("Part 2: %d\n", counter)
}
