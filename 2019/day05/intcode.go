package main

import (
	"fmt"
	"math"
)

type IntCode struct {
	ip      int
	program []int
	inputs  []int
}

func (i *IntCode) next() (exit bool) {
	if i.ip >= len(i.program) {
		return true
	}

	p := i.program[i.ip] % 100
	modes := i.program[i.ip] / 100

	switch p {
	case 99:
		return true
	case 1:
		i.add(modes)
	case 2:
		i.multi(modes)
	case 3:
		i.set(modes)
	case 4:
		i.output(modes)
	case 5:
		i.jumpIfTrue(modes)
	case 6:
		i.jumpIfFalse(modes)
	case 7:
		i.lessThan(modes)
	case 8:
		i.equals(modes)
	}

	return false
}

func (i *IntCode) run(output int) int {
	for {
		val := i.next()
		if val {
			break
		}
	}

	return i.program[output]
}

func getDigit(n int, position int) int {
	return (n / int(math.Pow10(position-1))) % 10
}

func (i *IntCode) add(modes int) {
	if i.ip+3 >= len(i.program) {
		i.ip = len(i.program)
		return
	}

	a1 := i.program[i.ip+1]
	if getDigit(modes, 1) == 0 {
		a1 = i.program[i.program[i.ip+1]]
	}

	a2 := i.program[i.ip+2]
	if getDigit(modes, 2) == 0 {
		a2 = i.program[i.program[i.ip+2]]
	}
	o := i.program[i.ip+3]

	i.program[o] = a1 + a2

	i.ip += 4
}

func (i *IntCode) multi(modes int) {
	if i.ip+3 >= len(i.program) {
		i.ip = len(i.program)
		return
	}

	a1 := i.program[i.ip+1]
	if getDigit(modes, 1) == 0 {
		a1 = i.program[i.program[i.ip+1]]
	}

	a2 := i.program[i.ip+2]
	if getDigit(modes, 2) == 0 {
		a2 = i.program[i.program[i.ip+2]]
	}
	o := i.program[i.ip+3]

	i.program[o] = a1 * a2

	i.ip += 4
}

func (i *IntCode) set(modes int) {
	if i.ip+1 >= len(i.program) {
		i.ip = len(i.program)
		return
	}

	input, inputs := i.inputs[0], i.inputs[1:]
	i.inputs = inputs

	o := i.program[i.ip+1]

	i.program[o] = input

	i.ip += 2
}

func (i *IntCode) output(modes int) {
	if i.ip+1 >= len(i.program) {
		i.ip = len(i.program)
		return
	}

	val := i.program[i.ip+1]
	if getDigit(modes, 1) == 0 {
		val = i.program[i.program[i.ip+1]]
	}

	fmt.Printf(" output: [%d]\n", val)

	i.ip += 2
}

func (i *IntCode) jumpIfTrue(modes int) {
	if i.ip+2 >= len(i.program) {
		i.ip = len(i.program)
		return
	}

	check := i.program[i.ip+1]
	if getDigit(modes, 1) == 0 {
		check = i.program[i.program[i.ip+1]]
	}

	if check != 0 {
		nIp := i.program[i.ip+2]
		if getDigit(modes, 2) == 0 {
			nIp = i.program[i.program[i.ip+2]]
		}
		i.ip = nIp
		return
	}

	i.ip += 3
}

func (i *IntCode) jumpIfFalse(modes int) {
	if i.ip+2 >= len(i.program) {
		i.ip = len(i.program)
		return
	}

	check := i.program[i.ip+1]
	if getDigit(modes, 1) == 0 {
		check = i.program[i.program[i.ip+1]]
	}

	if check == 0 {
		nIp := i.program[i.ip+2]
		if getDigit(modes, 2) == 0 {
			nIp = i.program[i.program[i.ip+2]]
		}
		i.ip = nIp
		return
	}

	i.ip += 3
}

func (i *IntCode) lessThan(modes int) {
	if i.ip+3 >= len(i.program) {
		i.ip = len(i.program)
		return
	}

	a1 := i.program[i.ip+1]
	if getDigit(modes, 1) == 0 {
		a1 = i.program[i.program[i.ip+1]]
	}

	a2 := i.program[i.ip+2]
	if getDigit(modes, 2) == 0 {
		a2 = i.program[i.program[i.ip+2]]
	}
	o := i.program[i.ip+3]

	if a1 < a2 {
		i.program[o] = 1
	} else {
		i.program[o] = 0
	}

	i.ip += 4
}

func (i *IntCode) equals(modes int) {
	if i.ip+3 >= len(i.program) {
		i.ip = len(i.program)
		return
	}

	a1 := i.program[i.ip+1]
	if getDigit(modes, 1) == 0 {
		a1 = i.program[i.program[i.ip+1]]
	}

	a2 := i.program[i.ip+2]
	if getDigit(modes, 2) == 0 {
		a2 = i.program[i.program[i.ip+2]]
	}
	o := i.program[i.ip+3]

	if a1 == a2 {
		i.program[o] = 1
	} else {
		i.program[o] = 0
	}

	i.ip += 4
}
