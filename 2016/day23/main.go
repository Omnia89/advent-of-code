package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"advent2016/util"
)

func main() {
	data := util.GetDataByRow("day23")
	// data := util.GetTestByRow("day23")

	list := parse(data)
	list2 := parse(data)

	part1(list)
	part2(list2)
}

type operation struct {
	op   string
	arg1 string
	arg2 string
}

func parse(data []string) []operation {
	ops := []operation{}

	for _, s := range data {
		parts := strings.Split(s, " ")
		o := operation{}
		o.op = parts[0]
		o.arg1 = parts[1]
		if len(parts) > 2 {
			o.arg2 = parts[2]
		}
		ops = append(ops, o)
	}
	return ops
}

func toggle(op string) string {
	values := map[string]string{
		"inc": "dec",
		"dec": "inc",
		"tgl": "inc",
		"jnz": "cpy",
		"cpy": "jnz",
	}

	return values[op]
}

func isNumber(v string) (int, bool) {
	numRe := regexp.MustCompile(`-?\d+`)
	if numRe.MatchString(v) {
		num, _ := strconv.Atoi(numRe.FindString(v))
		return num, true
	}
	return 0, false
}

func resolveVal(v string, registers map[string]int) int {
	if n, ok := isNumber(v); ok {
		return n
	}
	return registers[v]
}

func isMultPattern(i int, ops []operation, registers map[string]int) (bool, int) {
	if i+5 >= len(ops) {
		return false, i
	}

	o0 := ops[i]   // cpy B tmp
	o1 := ops[i+1] // inc A
	o2 := ops[i+2] // dec tmp
	o3 := ops[i+3] // jnz tmp -2
	o4 := ops[i+4] // dec C
	o5 := ops[i+5] // jnz C -5

	if o0.op != "cpy" {
		return false, i
	}
	b := o0.arg1
	tmp := o0.arg2

	if o1.op != "inc" {
		return false, i
	}
	a := o1.arg1

	if o2.op != "dec" || o2.arg1 != tmp {
		return false, i
	}

	if o3.op != "jnz" || o3.arg1 != tmp || o3.arg2 != "-2" {
		return false, i
	}

	if o4.op != "dec" {
		return false, i
	}
	c := o4.arg1

	if o5.op != "jnz" || o5.arg1 != c || o5.arg2 != "-5" {
		return false, i
	}

	registers[a] += resolveVal(b, registers) * resolveVal(c, registers)

	return true, i + 6
}

func exec(registers map[string]int, ops []operation) {
	i := 0
	for i < len(ops) {

		isMult, newI := isMultPattern(i, ops, registers)
		if isMult {
			i = newI
			continue
		}

		op := ops[i]

		switch op.op {
		case "inc":
			registers[op.arg1]++
			i++
		case "dec":
			registers[op.arg1]--
			i++
		case "jnz":
			check := resolveVal(op.arg1, registers)
			val := resolveVal(op.arg2, registers)

			if check != 0 {
				i += val
			} else {
				i++
			}
		case "cpy":
			_, okO := isNumber(op.arg2)
			if !okO {
				registers[op.arg2] = resolveVal(op.arg1, registers)
			}
			i++
		case "tgl":
			val := resolveVal(op.arg1, registers)
			if i+val < len(ops) && i+val >= 0 {
				ops[i+val].op = toggle(ops[i+val].op)
			}
			i++
		}
	}
}

func part1(ops []operation) {
	counter := 0

	registers := map[string]int{
		"a": 7,
		"b": 0,
		"c": 0,
		"d": 0,
	}

	exec(registers, ops)

	counter = registers["a"]

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(ops []operation) {
	counter := 0

	registers := map[string]int{
		"a": 12,
		"b": 0,
		"c": 0,
		"d": 0,
	}

	exec(registers, ops)

	counter = registers["a"]
	fmt.Printf("Part 2: %d\n", counter)
}
