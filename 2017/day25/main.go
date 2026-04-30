package main

import (
	"fmt"
	"strings"

	"advent2017/util"
)

func main() {
	data := util.GetDataByRow("day25")
	// data := util.GetTestByRow("day25")

	m := parse(data)

	part1(m)
	part2(m)
}

type Branch struct {
	move      int
	value     bool
	nextState string
}

type State struct {
	name        string
	trueBranch  Branch
	falseBranch Branch
}

type Machine struct {
	position     int
	steps        int
	initialState string
	currentState string
	register     map[int]bool
	states       map[string]State
}

func parse(data []string) Machine {
	m := Machine{}

	m.initialState = strings.Split(data[0][:len(data[0])-1], " ")[3]
	m.currentState = m.initialState
	m.steps = util.ToInt(strings.Split(data[1], " ")[5])
	m.register = make(map[int]bool)
	m.states = make(map[string]State)

	for i := 3; i < len(data); i += 10 {
		s := State{}

		s.name = strings.Split(data[i][:len(data[i])-1], " ")[2]
		if strings.HasSuffix(data[i+2], "0.") {
			s.falseBranch.value = false
		} else {
			s.falseBranch.value = true
		}
		falseMove := 1
		if strings.HasSuffix(data[i+3], "left.") {
			falseMove = -1
		}
		s.falseBranch.move = falseMove
		s.falseBranch.nextState = strings.Split(data[i+4][:len(data[i+4])-1], " ")[8]

		if strings.HasSuffix(data[i+6], "0.") {
			s.trueBranch.value = false
		} else {
			s.trueBranch.value = true
		}
		trueMove := 1
		if strings.HasSuffix(data[i+7], "left.") {
			trueMove = -1
		}
		s.trueBranch.move = trueMove
		s.trueBranch.nextState = strings.Split(data[i+8][:len(data[i+8])-1], " ")[8]

		m.states[s.name] = s
	}

	return m
}

func part1(m Machine) {
	counter := 0

	for range m.steps {
		if m.register[m.position] {
			m.register[m.position] = m.states[m.currentState].trueBranch.value
			m.position += m.states[m.currentState].trueBranch.move
			m.currentState = m.states[m.currentState].trueBranch.nextState
		} else {
			m.register[m.position] = m.states[m.currentState].falseBranch.value
			m.position += m.states[m.currentState].falseBranch.move
			m.currentState = m.states[m.currentState].falseBranch.nextState
		}
	}

	for _, b := range m.register {
		if b {
			counter++
		}
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(m Machine) {
	counter := 0

	fmt.Printf("Part 2: %d\n", counter)
}
