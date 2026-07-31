package main

import (
	"fmt"
)

type IntCodeChan struct {
	ip          int
	program     []int
	inputs      chan int
	outputs     chan int
	printOutput bool
}

func (i *IntCodeChan) next() (exit bool) {
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

func (i *IntCodeChan) run(output int) int {
	defer close(i.outputs)
	for {
		val := i.next()
		if val {
			break
		}
	}

	return i.program[output]
}

func (i *IntCodeChan) add(modes int) {
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

func (i *IntCodeChan) multi(modes int) {
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

func (i *IntCodeChan) set(modes int) {
	if i.ip+1 >= len(i.program) {
		i.ip = len(i.program)
		return
	}
	input := <-i.inputs

	o := i.program[i.ip+1]

	i.program[o] = input

	i.ip += 2
}

func (i *IntCodeChan) output(modes int) {
	if i.ip+1 >= len(i.program) {
		i.ip = len(i.program)
		return
	}

	val := i.program[i.ip+1]
	if getDigit(modes, 1) == 0 {
		val = i.program[i.program[i.ip+1]]
	}

	if i.printOutput {
		fmt.Printf(" output: [%d]\n", val)
	}
	i.outputs <- val

	i.ip += 2
}

func (i *IntCodeChan) jumpIfTrue(modes int) {
	if i.ip+2 >= len(i.program) {
		i.ip = len(i.program)
		return
	}

	check := i.program[i.ip+1]
	if getDigit(modes, 1) == 0 {
		check = i.program[i.program[i.ip+1]]
	}

	if check != 0 {
		nIP := i.program[i.ip+2]
		if getDigit(modes, 2) == 0 {
			nIP = i.program[i.program[i.ip+2]]
		}
		i.ip = nIP
		return
	}

	i.ip += 3
}

func (i *IntCodeChan) jumpIfFalse(modes int) {
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

func (i *IntCodeChan) lessThan(modes int) {
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

func (i *IntCodeChan) equals(modes int) {
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
