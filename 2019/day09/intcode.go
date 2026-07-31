package main

import (
	"fmt"
	"math"
)

type IntCode struct {
	ip           int
	relativeBase int
	program      map[int]int // Permit to go out of boundaries
	useChannels  bool
	inputChan    chan int
	outputChan   chan int
	inputArray   []int
	outputArray  []int
	printOutput  bool
}

func NewIntcodeChannel(program []int, inputChannel chan int, outputChannel chan int, printOutput bool) *IntCode {
	programMap := map[int]int{}

	for i, n := range program {
		programMap[i] = n
	}

	return &IntCode{
		ip:           0,
		relativeBase: 0,
		program:      programMap,
		useChannels:  true,
		inputChan:    inputChannel,
		outputChan:   outputChannel,
		printOutput:  printOutput,
	}
}

func NewIntcodeArray(program []int, inputArray []int, printOutput bool) *IntCode {
	programMap := map[int]int{}

	for i, n := range program {
		programMap[i] = n
	}

	return &IntCode{
		ip:           0,
		relativeBase: 0,
		program:      programMap,
		useChannels:  false,
		inputArray:   inputArray,
		printOutput:  printOutput,
	}
}

func getDigit(n int, position int) int {
	return (n / int(math.Pow10(position-1))) % 10
}

func (i *IntCode) getValue(modes int, position int) int {
	var val int
	switch getDigit(modes, position) {
	case 0:
		val = i.program[i.program[i.ip+position]]
	case 1:
		val = i.program[i.ip+position]
	case 2:
		val = i.program[i.relativeBase+i.program[i.ip+position]]
	}

	return val
}

func (i *IntCode) getAddress(modes int, position int) int {
	if getDigit(modes, position) == 2 {
		return i.relativeBase + i.program[i.ip+position]
	}
	return i.program[i.ip+position]
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
	case 9:
		i.setRelativeBase(modes)
	}

	return false
}

func (i *IntCode) run(output int) int {
	if i.useChannels {
		defer close(i.outputChan)
	}
	for {
		val := i.next()
		if val {
			break
		}
	}

	return i.program[output]
}

func (i *IntCode) add(modes int) {
	a1 := i.getValue(modes, 1)
	a2 := i.getValue(modes, 2)

	o := i.getAddress(modes, 3)

	i.program[o] = a1 + a2

	i.ip += 4
}

func (i *IntCode) multi(modes int) {
	a1 := i.getValue(modes, 1)
	a2 := i.getValue(modes, 2)

	o := i.getAddress(modes, 3)

	i.program[o] = a1 * a2

	i.ip += 4
}

func (i *IntCode) set(modes int) {
	var input int
	if i.useChannels {
		input = <-i.inputChan
	} else {
		in, inputs := i.inputArray[0], i.inputArray[1:]
		input = in
		i.inputArray = inputs
	}

	o := i.getAddress(modes, 1)

	i.program[o] = input

	i.ip += 2
}

func (i *IntCode) output(modes int) {
	val := i.getValue(modes, 1)

	if i.printOutput {
		fmt.Printf(" output: [%d]\n", val)
	}
	if i.useChannels {
		i.outputChan <- val
	} else {
		i.outputArray = append(i.outputArray, val)
	}

	i.ip += 2
}

func (i *IntCode) jumpIfTrue(modes int) {
	check := i.getValue(modes, 1)

	if check != 0 {
		nIP := i.getValue(modes, 2)
		i.ip = nIP
		return
	}

	i.ip += 3
}

func (i *IntCode) jumpIfFalse(modes int) {
	check := i.getValue(modes, 1)

	if check == 0 {
		nIP := i.getValue(modes, 2)
		i.ip = nIP
		return
	}

	i.ip += 3
}

func (i *IntCode) lessThan(modes int) {
	a1 := i.getValue(modes, 1)
	a2 := i.getValue(modes, 2)

	o := i.getAddress(modes, 3)

	if a1 < a2 {
		i.program[o] = 1
	} else {
		i.program[o] = 0
	}

	i.ip += 4
}

func (i *IntCode) equals(modes int) {
	a1 := i.getValue(modes, 1)
	a2 := i.getValue(modes, 2)

	o := i.getAddress(modes, 3)

	if a1 == a2 {
		i.program[o] = 1
	} else {
		i.program[o] = 0
	}

	i.ip += 4
}

func (i *IntCode) setRelativeBase(modes int) {
	val := i.getValue(modes, 1)

	i.relativeBase += val

	i.ip += 2
}
