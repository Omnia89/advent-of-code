package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"advent2024/util"
)

func main() {
	//data := util.GetTestByRow("day17")
	data := util.GetDataByRow("day17")

	machine := parseMachine(data)
	machine2 := parseMachine(data)

	part1(machine)
	part2(machine2)
}

type Machine struct {
	registers map[string]int // A, B, C
	program   []int
	pointer   int
}

func (m Machine) toString() string {
	progs := []string{}
	for _, p := range m.program {
		progs = append(progs, fmt.Sprintf("%d", p))
	}
	return fmt.Sprintf("%s - [%d][%d][%d] - index[%d -> %d]", strings.Join(progs, ","), m.registers["A"], m.registers["B"], m.registers["C"], m.pointer, m.program[m.pointer])
}

func parseMachine(data []string) Machine {
	machine := Machine{}

	machine.pointer = 0
	machine.registers = make(map[string]int)

	if len(data) < 5 {
		panic("Not valid input")
	}

	machine.registers["A"] = util.ToInt(strings.Split(data[0], ":")[1])
	machine.registers["B"] = util.ToInt(strings.Split(data[1], ":")[1])
	machine.registers["C"] = util.ToInt(strings.Split(data[2], ":")[1])

	machine.program = util.StringToIntSlice(strings.Split(data[4], ":")[1], ",")

	return machine
}

func comboOp(machine Machine, operand int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return machine.registers["A"]
	case 5:
		return machine.registers["B"]
	case 6:
		return machine.registers["C"]
	}
	panic("Invalid comboOp")
}

func executeOp(machine *Machine) (out int, printOut bool, halt bool) {
	if machine.pointer >= len(machine.program) {
		return 0, false, true
	}
	opcode := machine.program[machine.pointer]
	machine.pointer++

	if machine.pointer >= len(machine.program) {
		return 0, false, true
	}
	operand := machine.program[machine.pointer]
	machine.pointer++

	if opcode == 0 {
		// adv
		numerand := float64(machine.registers["A"])
		denominator := math.Trunc(math.Pow(2, float64(comboOp(*machine, operand))))
		machine.registers["A"] = int(numerand / denominator)
		return 0, false, false
	}

	if opcode == 1 {
		// bxl
		machine.registers["B"] = machine.registers["B"] ^ operand
		return 0, false, false
	}

	if opcode == 2 {
		// bst
		machine.registers["B"] = comboOp(*machine, operand) % 8
		return 0, false, false
	}

	if opcode == 3 {
		// jnz
		if machine.registers["A"] != 0 {
			machine.pointer = operand
		}
		return 0, false, false
	}

	if opcode == 4 {
		// bxc
		machine.registers["B"] = machine.registers["B"] ^ machine.registers["C"]
		// machine.pointer--
		return 0, false, false
	}

	if opcode == 5 {
		// out
		return comboOp(*machine, operand) % 8, true, false
	}

	if opcode == 6 {
		// bdv
		numerand := float64(machine.registers["A"])
		denominator := math.Trunc(math.Pow(2, float64(comboOp(*machine, operand))))
		machine.registers["B"] = int(numerand / denominator)
		return 0, false, false
	}

	if opcode == 7 {
		// cdv
		numerand := float64(machine.registers["A"])
		denominator := math.Trunc(math.Pow(2, float64(comboOp(*machine, operand))))
		machine.registers["C"] = int(numerand / denominator)
		return 0, false, false
	}

	panic("Non valid opcode")
}

func Run(machine Machine) []int {
	halt := false
	output := []int{}
	var val int
	var out bool

	for !halt {
		val, out, halt = executeOp(&machine)
		if out {
			output = append(output, val)
		}
	}

	return output
}

func part1(machine Machine) {
	intOut := Run(machine)

	output := []string{}
	for _, i := range intOut {
		output = append(output, strconv.Itoa(i))
	}

	result := strings.Join(output, ",")

	fmt.Printf("Part 1: %s\n", result)
}

func checkArray(a []int, b []int) bool {
	if len(b) != len(a) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func newMachineFrom(mach Machine, a int) Machine {
	m := Machine{
		registers: make(map[string]int),
		program:   mach.program,
		pointer:   0,
	}

	for k, v := range mach.registers {
		m.registers[k] = v
	}

	m.registers["A"] = a
	return m
}

type item struct {
	a int
	n int
}

func part2(machine Machine) {
	counter := 0

	queue := []item{
		{0, 1},
	}

	var it item
	maxN := 0
	for len(queue) > 0 {
		it, queue = queue[0], queue[1:]

		if maxN < it.n {
			maxN = it.n
		}

		if it.n > len(machine.program) {
			counter = it.a
			break
		}

		// I start checking from the last entry in program, and find which "a" will generate this entry
		// since the output is generate by "a // 8", we can "segment" the registry "a"" to generate every entry,
		// starting from the end
		for v := range 8 {
			newA := (it.a << 3) | v
			newMachine := newMachineFrom(machine, newA)

			programTarget := machine.program[len(machine.program)-it.n:] // get the lat N values

			output := Run(newMachine)
			if checkArray(programTarget, output) {
				queue = append(queue, item{newA, it.n + 1})
			}
		}
	}
	fmt.Printf("maxN [%d]\n", maxN)

	fmt.Printf("Part 2: %d\n", counter)
}
