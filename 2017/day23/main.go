package main

import (
	"fmt"
	"regexp"
	"strings"

	"advent2017/util"
)

func main() {
	data := util.GetDataByRow("day23")
	// data := util.GetTestByRow("day23")

	list := parse(data)

	part1(list)
	// part2(list)
	part2Optimized(list)
}

type instruction struct {
	op   string
	arg1 string
	arg2 string
}

func parse(data []string) []instruction {
	is := []instruction{}

	for _, s := range data {
		parts := strings.Split(s, " ")
		is = append(is, instruction{
			parts[0],
			parts[1],
			parts[2],
		})
	}

	return is
}

var alpha = regexp.MustCompile(`[a-z]`)

func getValue(val string, registers map[string]int) int {
	if alpha.MatchString(val) {
		return registers[val]
	}
	return util.ToInt(val)
}

func execute(i instruction, registers map[string]int) int {
	delta := 1
	switch i.op {
	case "set":
		registers[i.arg1] = getValue(i.arg2, registers)
	case "sub":
		registers[i.arg1] -= getValue(i.arg2, registers)
	case "mul":
		registers[i.arg1] *= getValue(i.arg2, registers)
	case "jnz":
		check := getValue(i.arg1, registers)
		if check != 0 {
			delta = getValue(i.arg2, registers)
		}
	}
	return delta
}

func part1(data []instruction) {
	counter := 0

	pointer := 0
	registers := map[string]int{}

	for pointer < len(data) {
		i := data[pointer]
		d := execute(i, registers)
		pointer += d

		if i.op == "mul" {
			counter++
		}
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data []instruction) {
	counter := 0

	pointer := 0
	registers := map[string]int{"a": 1}

	for pointer < len(data) {
		i := data[pointer]
		d := execute(i, registers)
		pointer += d
	}

	counter = registers["h"]
	fmt.Printf("Part 2: %d\n", counter)
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func findStep(data []instruction) int {
	// the outer loop increments b via "sub b -N" (i.e. b += N)
	// it's the last "sub b <negative>" in the instruction list
	for i := len(data) - 1; i >= 0; i-- {
		ins := data[i]
		if ins.op == "sub" && ins.arg1 == "b" {
			n := util.ToInt(ins.arg2)
			if n < 0 {
				return -n
			}
		}
	}
	return 1
}

func part2Optimized(data []instruction) {
	registers := map[string]int{"a": 1}
	pointer := 0
	for pointer < len(data) {
		i := data[pointer]
		if i.op == "set" && i.arg1 == "f" && i.arg2 == "1" {
			break
		}
		pointer += execute(i, registers)
	}

	b := registers["b"]
	c := registers["c"]
	step := findStep(data)

	h := 0
	for n := b; n <= c; n += step {
		if !isPrime(n) {
			h++
		}
	}

	fmt.Printf("Part 2: %d\n", h)
}
