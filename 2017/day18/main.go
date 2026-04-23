package main

import (
	"fmt"
	"regexp"
	"strings"

	"advent2017/util"
)

func main() {
	data := util.GetDataByRow("day18")
	// data := util.GetTestByRow("day18")

	list := parse(data)

	part1(list)
	part2(list)
}

type instruction struct {
	op       string
	register string
	value    string
}

func parse(data []string) []instruction {
	r := []instruction{}

	for _, s := range data {
		parts := strings.Split(s, " ")
		o := instruction{
			op:       parts[0],
			register: parts[1],
		}
		if len(parts) == 3 {
			o.value = parts[2]
		}
		r = append(r, o)
	}

	return r
}

var alpha = regexp.MustCompile(`[a-z]`)

func getValue(op instruction, registers map[string]int) int {
	if alpha.MatchString(op.value) {
		return registers[op.value]
	}
	return util.ToInt(op.value)
}

func execute(op instruction, pos int, registers map[string]int) (int, int, bool) {
	value := getValue(op, registers)
	switch op.op {
	case "snd":
		return pos + 1, registers[op.register], true
	case "set":
		registers[op.register] = value
		return pos + 1, 0, true
	case "add":
		registers[op.register] += value
		return pos + 1, 0, true
	case "mul":
		registers[op.register] *= value
		return pos + 1, 0, true
	case "mod":
		registers[op.register] %= value
		return pos + 1, 0, true
	case "rcv":
		return pos + 1, 0, registers[op.register] != 0
	case "jgz":
		p := pos + 1
		regVal := 0
		if alpha.MatchString(op.register) {
			regVal = registers[op.register]
		} else {
			regVal = util.ToInt(op.register)
		}
		if regVal > 0 {
			p = pos + value
		}
		return p, 0, true
	}
	panic("Invalid op")
}

func part1(ops []instruction) {
	counter := 0

	regs := map[string]int{}
	pos := 0
	lastSound := 0

	for pos < len(ops) {
		o := ops[pos]
		p, sound, ok := execute(o, pos, regs)
		pos = p
		if o.op == "snd" {
			lastSound = sound
		}
		if o.op == "rcv" && ok {
			counter = lastSound
			break
		}
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(ops []instruction) {
	counter := 0

	regsA := map[string]int{"p": 0}
	regsB := map[string]int{"p": 1}

	posA := 0
	posB := 0

	queueA := []int{}
	queueB := []int{}

	for {
		skipA := false
		skipB := false
		var oA instruction
		var oB instruction

		if posA >= len(ops) {
			skipA = true
		} else {
			oA = ops[posA]
		}
		if posB >= len(ops) {
			skipB = true
		} else {
			oB = ops[posB]
		}

		if (skipA && skipB) || (skipA || (oA.op == "rcv" && len(queueA) == 0)) && (skipB || (oB.op == "rcv" && len(queueB) == 0)) {
			break
		}

		if !skipA && oA.op == "rcv" {
			if len(queueA) > 0 {
				var v int
				v, queueA = queueA[0], queueA[1:]
				regsA[oA.register] = v
				posA++
			}
			skipA = true
		}

		if !skipB && oB.op == "rcv" {
			if len(queueB) > 0 {
				var v int
				v, queueB = queueB[0], queueB[1:]
				regsB[oB.register] = v
				posB++
			}
			skipB = true
		}
		if !skipA {
			pA, sendA, _ := execute(oA, posA, regsA)
			posA = pA
			if oA.op == "snd" {
				queueB = append(queueB, sendA)
			}
		}

		if !skipB {
			pB, sendB, _ := execute(oB, posB, regsB)
			posB = pB
			if oB.op == "snd" {
				counter++
				queueA = append(queueA, sendB)
			}
		}
		//fmt.Printf("pA[%d][%s][%d] pB[%d][%s][%d]\n", posA, oA.op, regsA["i"], posB, oB.op, regsB["i"])
	}

	fmt.Printf("Part 2: %d\n", counter)
}
