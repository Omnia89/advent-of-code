package main

import (
	"fmt"
	"slices"
	"strings"

	"advent2018/util"
)

func main() {
	data := util.GetDataByRow("day19")
	// data := util.GetTestByRow("day19")

	pointer, insts := parse(data)

	part1(pointer, insts)
	part2(pointer, insts)
}

type Instruction struct {
	op     string
	values [3]int
}

func parse(data []string) (pointerRegister int, insts []Instruction) {
	for i, s := range data {
		if i == 0 {
			_, pointer, _ := strings.Cut(s, " ")
			pointerRegister = util.ToInt(pointer)
			continue
		}
		op, numbers, _ := strings.Cut(s, " ")
		insts = append(insts, Instruction{
			op,
			[3]int(util.StringToIntSlice(numbers, " ")),
		})
	}

	return
}

func addr(reg []int, instr Instruction) []int {
	r := slices.Clone(reg)
	r[instr.values[2]] = reg[instr.values[0]] + reg[instr.values[1]]
	return r
}

func addi(reg []int, instr Instruction) []int {
	r := slices.Clone(reg)
	r[instr.values[2]] = reg[instr.values[0]] + instr.values[1]
	return r
}

func mulr(reg []int, instr Instruction) []int {
	r := slices.Clone(reg)
	r[instr.values[2]] = reg[instr.values[0]] * reg[instr.values[1]]
	return r
}

func muli(reg []int, instr Instruction) []int {
	r := slices.Clone(reg)
	r[instr.values[2]] = reg[instr.values[0]] * instr.values[1]
	return r
}

func banr(reg []int, instr Instruction) []int {
	r := slices.Clone(reg)
	r[instr.values[2]] = reg[instr.values[0]] & reg[instr.values[1]]
	return r
}

func bani(reg []int, instr Instruction) []int {
	r := slices.Clone(reg)
	r[instr.values[2]] = reg[instr.values[0]] & instr.values[1]
	return r
}

func borr(reg []int, instr Instruction) []int {
	r := slices.Clone(reg)
	r[instr.values[2]] = reg[instr.values[0]] | reg[instr.values[1]]
	return r
}

func bori(reg []int, instr Instruction) []int {
	r := slices.Clone(reg)
	r[instr.values[2]] = reg[instr.values[0]] | instr.values[1]
	return r
}

func setr(reg []int, instr Instruction) []int {
	r := slices.Clone(reg)
	r[instr.values[2]] = reg[instr.values[0]]
	return r
}

func seti(reg []int, instr Instruction) []int {
	r := slices.Clone(reg)
	r[instr.values[2]] = instr.values[0]
	return r
}

func gtir(reg []int, instr Instruction) []int {
	r := slices.Clone(reg)
	if instr.values[0] > reg[instr.values[1]] {
		r[instr.values[2]] = 1
	} else {
		r[instr.values[2]] = 0
	}
	return r
}

func gtri(reg []int, instr Instruction) []int {
	r := slices.Clone(reg)
	if reg[instr.values[0]] > instr.values[1] {
		r[instr.values[2]] = 1
	} else {
		r[instr.values[2]] = 0
	}
	return r
}

func gtrr(reg []int, instr Instruction) []int {
	r := slices.Clone(reg)
	if reg[instr.values[0]] > reg[instr.values[1]] {
		r[instr.values[2]] = 1
	} else {
		r[instr.values[2]] = 0
	}
	return r
}

func eqir(reg []int, instr Instruction) []int {
	r := slices.Clone(reg)
	if instr.values[0] == reg[instr.values[1]] {
		r[instr.values[2]] = 1
	} else {
		r[instr.values[2]] = 0
	}
	return r
}

func eqri(reg []int, instr Instruction) []int {
	r := slices.Clone(reg)
	if reg[instr.values[0]] == instr.values[1] {
		r[instr.values[2]] = 1
	} else {
		r[instr.values[2]] = 0
	}
	return r
}

func eqrr(reg []int, instr Instruction) []int {
	r := slices.Clone(reg)
	if reg[instr.values[0]] == reg[instr.values[1]] {
		r[instr.values[2]] = 1
	} else {
		r[instr.values[2]] = 0
	}
	return r
}

type InstrFunc = func([]int, Instruction) []int

var instructionList = map[string]InstrFunc{
	"addr": addr,
	"addi": addi,
	"mulr": mulr,
	"muli": muli,
	"banr": banr,
	"bani": bani,
	"borr": borr,
	"bori": bori,
	"setr": setr,
	"seti": seti,
	"gtir": gtir,
	"gtri": gtri,
	"gtrr": gtrr,
	"eqir": eqir,
	"eqri": eqri,
	"eqrr": eqrr,
}

func part1(pointerRegister int, instrs []Instruction) {
	counter := 0

	registers := slices.Repeat([]int{0}, 6)
	pointer := registers[pointerRegister]

	for pointer >= 0 && pointer < len(instrs) {
		inst := instrs[pointer]
		registers = instructionList[inst.op](registers, inst)
		registers[pointerRegister]++
		pointer = registers[pointerRegister]
	}

	counter = registers[0]

	fmt.Printf("Part 1: %d\n", counter)
}

func sumOfDivisors(n int) int {
	sum := 0
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			sum += i
			if i != n/i {
				sum += n / i
			}
		}
	}
	return sum
}

func part2(pointerRegister int, instrs []Instruction) {
	registers := slices.Repeat([]int{0}, 6)
	registers[0] = 1
	pointer := registers[pointerRegister]

	// Run until first back-edge: init done, main loop registers are set
	prevPointer := pointer
	for pointer >= 0 && pointer < len(instrs) {
		inst := instrs[pointer]
		registers = instructionList[inst.op](registers, inst)
		registers[pointerRegister]++
		pointer = registers[pointerRegister]
		if pointer < prevPointer {
			break
		}
		prevPointer = pointer
	}

	// reg[1] is input specific, found looking at the input by hand
	fmt.Printf("Part 2: %d\n", sumOfDivisors(registers[1]))
}
