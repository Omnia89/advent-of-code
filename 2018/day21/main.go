package main

import (
	"fmt"
	"slices"
	"strings"

	"advent2018/util"
)

func main() {
	data := util.GetDataByRow("day21")
	// data := util.GetTestByRow("day21")

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

func findEqrr(instrs []Instruction) (idx, checkRegister int) {
	for i, inst := range instrs {
		if inst.op == "eqrr" {
			checkRegister = inst.values[0]
			if checkRegister == 0 {
				checkRegister = inst.values[1]
			}
			return i, checkRegister
		}
	}
	return -1, -1
}

func part1(pointerRegister int, instrs []Instruction) {
	eqrrIdx, checkRegister := findEqrr(instrs)

	reg := make([]int, 6)
	for {
		ip := reg[pointerRegister]
		if ip == eqrrIdx {
			fmt.Printf("Part 1: %d\n", reg[checkRegister])
			return
		}
		if ip < 0 || ip >= len(instrs) {
			break
		}
		reg = instructionList[instrs[ip].op](reg, instrs[ip])
		reg[pointerRegister]++
	}
}

func part2(pointerRegister int, instrs []Instruction) {
	eqrrIdx, checkRegister := findEqrr(instrs)

	reg := make([]int, 6)
	seen := map[int]bool{}
	last := 0
	for {
		ip := reg[pointerRegister]
		if ip == eqrrIdx {
			val := reg[checkRegister]
			if seen[val] {
				fmt.Printf("Part 2: %d\n", last)
				return
			}
			seen[val] = true
			last = val
		}
		if ip < 0 || ip >= len(instrs) {
			break
		}
		reg = instructionList[instrs[ip].op](reg, instrs[ip])
		reg[pointerRegister]++
	}
}
