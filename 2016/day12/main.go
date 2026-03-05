package main

import (
	"fmt"
	"regexp"
	"strings"

	"advent2016/util"
)

func main() {
	data := util.GetDataByRow("day12")
	// data := util.GetTestByRow("day12")

	list := parse(data)

	part1(list)
	part2(list)
}

type operation struct {
	op     string
	input  string
	value  int
	output string
}

func parse(data []string) []operation {
	ops := []operation{}

	numRe := regexp.MustCompile(`\d+`)

	for _, s := range data {
		parts := strings.Split(s, " ")
		o := operation{}
		o.op = parts[0]
		if o.op == "inc" || o.op == "dec" {
			o.output = parts[1]
		}
		if o.op == "jnz" {
			if numRe.MatchString(parts[1]) {
				o.value = util.ToInt(parts[1])
			} else {
				o.input = parts[1]
			}
			o.value = util.ToInt(parts[2])
		}
		if o.op == "cpy" {
			o.output = parts[2]
			if numRe.MatchString(parts[1]) {
				o.value = util.ToInt(parts[1])
			} else {
				o.input = parts[1]
			}
		}
		ops = append(ops, o)
	}
	return ops
}

// 9227771 too high
func part1(ops []operation) {
	counter := 0

	registers := map[string]int{
		"a": 0,
		"b": 0,
		"c": 0,
		"d": 0,
	}

	i := 0
	for i < len(ops) {
		op := ops[i]

		switch op.op {
		case "inc":
			registers[op.output]++
			i++
		case "dec":
			registers[op.output]--
			i++
		case "jnz":
			chk := op.value
			if op.input != "" {
				chk = registers[op.input]
			}
			if chk != 0 {
				i += op.value
			} else {
				i++
			}
		case "cpy":
			val := op.value
			if op.input != "" {
				val = registers[op.input]
			}
			registers[op.output] = val
			i++
		}
	}

	counter = registers["a"]

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(ops []operation) {
	counter := 0
	registers := map[string]int{
		"a": 0,
		"b": 0,
		"c": 1,
		"d": 0,
	}

	i := 0
	for i < len(ops) {
		op := ops[i]

		switch op.op {
		case "inc":
			registers[op.output]++
			i++
		case "dec":
			registers[op.output]--
			i++
		case "jnz":
			chk := op.value
			if op.input != "" {
				chk = registers[op.input]
			}
			if chk != 0 {
				i += op.value
			} else {
				i++
			}
		case "cpy":
			val := op.value
			if op.input != "" {
				val = registers[op.input]
			}
			registers[op.output] = val
			i++
		}
	}

	counter = registers["a"]

	fmt.Printf("Part 2: %d\n", counter)
}
