package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"advent2016/util"
)

func main() {
	data := util.GetDataByRow("day25")
	// data := util.GetTestByRow("day25")

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

func execSingle(registers map[string]int, ops []operation, i int, output *[]int) int {

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
	case "out":
		val := resolveVal(op.arg1, registers)
		*output = append(*output, val)
		i++
	}
	return i
}

func part1(ops []operation) {
	counter := 0

	aInit := 0

outerLoop:
	for {
		registers := map[string]int{
			"a": aInit,
			"b": 0,
			"c": 0,
			"d": 0,
		}

		output := []int{}
		i := 0
		oldLen := 0
		for {
			i = execSingle(registers, ops, i, &output)
			if len(output) != oldLen {
				l := len(output)
				if l > 0 && output[l-1] != (l-1)%2 {
					aInit++
					continue outerLoop
				}
				oldLen = len(output)
			}
			if oldLen > 30 {
				break outerLoop
			}
		}
	}

	counter = aInit

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(ops []operation) {
	counter := 0

	fmt.Printf("Part 2: %d\n", counter)
}
