package main

import (
	"fmt"
	"maps"
	"regexp"
	"slices"
	"strings"

	"advent2018/util"
)

func main() {
	data := util.GetDataByRow("day16")
	// data := util.GetTestByRow("day16")

	tests, insts := parse(data)

	part1(tests)
	part2(tests, insts)
}

type Instruction [4]int

type TestCase struct {
	before []int
	after  []int
	instr  Instruction
}

func (t TestCase) String() string {
	return fmt.Sprintf("Before: %v\n%v\nAfter: %v", t.before, t.instr, t.after)
}

func parse(data []string) (tests []TestCase, insts []Instruction) {
	i := 0

	numberRegex := regexp.MustCompile(`\[?(\d+),?\s(\d+),?\s(\d+),?\s(\d+)\]?`)

	var t TestCase
	toSave := false
	emptyCount := 0
	for i < len(data) {
		s := data[i]
		i++
		if s == "" {
			if toSave {
				tests = append(tests, t)
			}
			toSave = false
			emptyCount++
			if emptyCount == 3 {
				break
			}
			continue
		}
		toSave = true
		emptyCount = 0
		strNumbers := numberRegex.FindStringSubmatch(s)
		numbers := []int{
			util.ToInt(strNumbers[1]),
			util.ToInt(strNumbers[2]),
			util.ToInt(strNumbers[3]),
			util.ToInt(strNumbers[4]),
		}

		if strings.HasPrefix(s, "Before") {
			t.before = numbers
		} else if strings.HasPrefix(s, "After") {
			t.after = numbers
		} else {
			t.instr = Instruction{numbers[0], numbers[1], numbers[2], numbers[3]}
		}
	}
	if toSave {
		tests = append(tests, t)
	}
	for i < len(data) {
		s := data[i]
		strNumbers := numberRegex.FindStringSubmatch(s)
		insts = append(insts, Instruction{
			util.ToInt(strNumbers[1]),
			util.ToInt(strNumbers[2]),
			util.ToInt(strNumbers[3]),
			util.ToInt(strNumbers[4]),
		})
		i++
	}
	return
}

func addr(reg []int, instr Instruction) []int {
	r := slices.Clone(reg)
	r[instr[3]] = reg[instr[1]] + reg[instr[2]]
	return r
}

func addi(reg []int, instr Instruction) []int {
	r := slices.Clone(reg)
	r[instr[3]] = reg[instr[1]] + instr[2]
	return r
}

func mulr(reg []int, instr Instruction) []int {
	r := slices.Clone(reg)
	r[instr[3]] = reg[instr[1]] * reg[instr[2]]
	return r
}

func muli(reg []int, instr Instruction) []int {
	r := slices.Clone(reg)
	r[instr[3]] = reg[instr[1]] * instr[2]
	return r
}

func banr(reg []int, instr Instruction) []int {
	r := slices.Clone(reg)
	r[instr[3]] = reg[instr[1]] & reg[instr[2]]
	return r
}

func bani(reg []int, instr Instruction) []int {
	r := slices.Clone(reg)
	r[instr[3]] = reg[instr[1]] & instr[2]
	return r
}

func borr(reg []int, instr Instruction) []int {
	r := slices.Clone(reg)
	r[instr[3]] = reg[instr[1]] | reg[instr[2]]
	return r
}

func bori(reg []int, instr Instruction) []int {
	r := slices.Clone(reg)
	r[instr[3]] = reg[instr[1]] | instr[2]
	return r
}

func setr(reg []int, instr Instruction) []int {
	r := slices.Clone(reg)
	r[instr[3]] = reg[instr[1]]
	return r
}

func seti(reg []int, instr Instruction) []int {
	r := slices.Clone(reg)
	r[instr[3]] = instr[1]
	return r
}

func gtir(reg []int, instr Instruction) []int {
	r := slices.Clone(reg)
	if instr[1] > reg[instr[2]] {
		r[instr[3]] = 1
	} else {
		r[instr[3]] = 0
	}
	return r
}

func gtri(reg []int, instr Instruction) []int {
	r := slices.Clone(reg)
	if reg[instr[1]] > instr[2] {
		r[instr[3]] = 1
	} else {
		r[instr[3]] = 0
	}
	return r
}

func gtrr(reg []int, instr Instruction) []int {
	r := slices.Clone(reg)
	if reg[instr[1]] > reg[instr[2]] {
		r[instr[3]] = 1
	} else {
		r[instr[3]] = 0
	}
	return r
}

func eqir(reg []int, instr Instruction) []int {
	r := slices.Clone(reg)
	if instr[1] == reg[instr[2]] {
		r[instr[3]] = 1
	} else {
		r[instr[3]] = 0
	}
	return r
}

func eqri(reg []int, instr Instruction) []int {
	r := slices.Clone(reg)
	if reg[instr[1]] == instr[2] {
		r[instr[3]] = 1
	} else {
		r[instr[3]] = 0
	}
	return r
}

func eqrr(reg []int, instr Instruction) []int {
	r := slices.Clone(reg)
	if reg[instr[1]] == reg[instr[2]] {
		r[instr[3]] = 1
	} else {
		r[instr[3]] = 0
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

func part1(tests []TestCase) {
	counter := 0

	testCounter := 0
loopTest:
	for _, t := range tests {
		// fmt.Println(t.String())
		testCounter = 0
		for _, inst := range instructionList {
			if testCounter >= 3 {
				counter++
				continue loopTest
			}
			res := inst(t.before, t.instr)
			// fmt.Printf("%s\nafter: %v\nresul: %v\n\n", k, t.after, res)
			if slices.Compare(res, t.after) == 0 {
				testCounter++
			}
		}
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(tests []TestCase, insts []Instruction) {
	counter := 0

	foundCodes := map[int]InstrFunc{}
	stillInstr := maps.Clone(instructionList)

	cache := map[int]map[string]bool{}

	for len(stillInstr) > 0 {
		removed := 0
		for i, t := range tests {
			if _, ok := foundCodes[t.instr[0]]; ok {
				continue
			}
			count := 0
			if _, ok := cache[i]; !ok {
				for k, instr := range stillInstr {
					res := instr(t.before, t.instr)
					if slices.Compare(res, t.after) == 0 {
						if _, ok := cache[i]; !ok {
							cache[i] = map[string]bool{}
						}
						cache[i][k] = true
						count++
					}
				}
			} else {
				count = len(cache[i])
			}
			if count == 1 {
				toRemove := ""
				for s := range cache[i] {
					foundCodes[t.instr[0]] = stillInstr[s]
					toRemove = s
				}
				for ic := range cache {
					if ic == i {
						continue
					}
					delete(cache[ic], toRemove)
					delete(stillInstr, toRemove)
				}
				removed++
			}
		}
		if removed == 0 && len(stillInstr) > 0 {
			panic("infinite loop, check closely")
		}
	}

	registers := []int{0, 0, 0, 0}

	for _, ins := range insts {
		op := foundCodes[ins[0]]
		registers = op(registers, ins)
	}
	counter = registers[0]

	fmt.Printf("Part 2: %d\n", counter)
}
