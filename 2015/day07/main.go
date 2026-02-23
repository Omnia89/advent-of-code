package main

import (
	"fmt"
	"regexp"
	"strings"

	"advent2015/util"
)

func main() {
	data := util.GetDataByRow("day07")
	// data := util.GetTestByRow("day07")

	problem := parse(data)

	part1(problem)

	problem2 := parse(data)

	problem2.values["b"] = problem.values["a"]

	part2(problem2)
}

type instruction struct {
	index int
	op    string
	reg1  string
	reg2  string
	out   string
}

type problem struct {
	instructions []instruction
	values       map[string]int
}

func (i instruction) toString() string {
	return fmt.Sprintf(" [%d][%s] r1[%s] r2[%s] -> [%s]", i.index, i.op, i.reg1, i.reg2, i.out)
}

func parse(data []string) problem {
	inst := []instruction{}
	values := map[string]int{}

	digitRegex, _ := regexp.Compile("^\\d+$")

	counter := 0

	setValue := func(s string) string {
		if digitRegex.MatchString(s) {
			key := fmt.Sprintf("flakey%d", counter)
			values[key] = util.ToInt(s)
			counter++
			return key
		}

		return s
	}

	for index, s := range data {
		i := instruction{}
		left, out, _ := strings.Cut(s, " -> ")
		i.index = index
		i.out = out

		if strings.Contains(left, "AND") {
			a, b, _ := strings.Cut(left, " AND ")
			i.op = "AND"

			i.reg1 = setValue(a)
			i.reg2 = setValue(b)
		} else if strings.Contains(left, "OR") {
			a, b, _ := strings.Cut(left, " OR ")
			i.op = "OR"

			i.reg1 = setValue(a)
			i.reg2 = setValue(b)
		} else if strings.Contains(left, "LSHIFT") {
			a, b, _ := strings.Cut(left, " LSHIFT ")
			i.op = "LSHIFT"

			i.reg1 = setValue(a)
			i.reg2 = setValue(b)
		} else if strings.Contains(left, "RSHIFT") {
			a, b, _ := strings.Cut(left, " RSHIFT ")
			i.op = "RSHIFT"

			i.reg1 = setValue(a)
			i.reg2 = setValue(b)
		} else if strings.Contains(left, "NOT") {
			_, b, _ := strings.Cut(left, " ")
			i.op = "NOT"
			i.reg1 = setValue(b)
		} else {
			i.op = "SET"

			i.reg1 = setValue(left)
		}

		inst = append(inst, i)

	}

	return problem{inst, values}
}

func exec(op string, v1 int, v2 int) int {
	r := 0

	switch op {
	case "SET":
		r = v1
	case "AND":
		r = v1 & v2
	case "OR":
		r = v1 | v2
	case "LSHIFT":
		r = v1 << v2
	case "RSHIFT":
		r = v1 >> v2
	case "NOT":
		r = ^v1
	}

	return r
}

func part1(prob problem) {
	counter := 0

	queue := prob.instructions
	var inst instruction
	for len(queue) > 0 {
		inst, queue = queue[0], queue[1:]

		var c bool
		var v1, v2 int
		ok := true
		v1, c = prob.values[inst.reg1]
		ok = ok && c

		if inst.op != "NOT" && inst.op != "SET" {
			v2, c = prob.values[inst.reg2]
			ok = ok && c
		}

		if !ok {
			queue = append(queue, inst)
			continue
		}
		// fmt.Printf("%s - v[%v]\n", inst.toString(), prob.values)

		prob.values[inst.out] = exec(inst.op, v1, v2)
		if inst.op == "NOT" || inst.op == "SET" {
			delete(prob.values, inst.reg1)
		}
	}

	counter = prob.values["a"]

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(prob problem) {
	counter := 0

	queue := prob.instructions
	var inst instruction
	for len(queue) > 0 {
		inst, queue = queue[0], queue[1:]

		if inst.out == "b" {
			continue
		}

		var c bool
		var v1, v2 int
		ok := true
		v1, c = prob.values[inst.reg1]
		ok = ok && c

		if inst.op != "NOT" && inst.op != "SET" {
			v2, c = prob.values[inst.reg2]
			ok = ok && c
		}

		if !ok {
			queue = append(queue, inst)
			continue
		}
		// fmt.Printf("%s - v[%v]\n", inst.toString(), prob.values)

		prob.values[inst.out] = exec(inst.op, v1, v2)
		if inst.op == "NOT" || inst.op == "SET" {
			delete(prob.values, inst.reg1)
		}
	}

	counter = prob.values["a"]

	fmt.Printf("Part 2: %d\n", counter)
}
