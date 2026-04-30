package main

import (
	"fmt"
	"slices"
	"strings"

	"advent2017/util"
)

func main() {
	data := util.GetDataByRow("day24")
	// data := util.GetTestByRow("day24")

	list := parse(data)

	part1(list)
	part2(list)
}

type Connector struct {
	sideA int
	sideB int
}

func parse(data []string) []Connector {
	cs := []Connector{}

	for _, s := range data {
		a, b, _ := strings.Cut(s, "/")
		cs = append(cs, Connector{
			sideA: util.ToInt(a),
			sideB: util.ToInt(b),
		})
	}

	return cs
}

func getStrength(cs []Connector) int {
	s := 0
	for _, c := range cs {
		s += c.sideA + c.sideB
	}
	return s
}

func part1(data []Connector) {
	counter := 0

	connectorMap := map[int][]Connector{}
	for _, c := range data {
		connectorMap[c.sideA] = append(connectorMap[c.sideA], c)
		connectorMap[c.sideB] = append(connectorMap[c.sideB], c)
	}

	type state struct {
		chain []Connector
		next  int
	}

	maxStrength := 0

	queue := []state{}
	for _, c := range connectorMap[0] {

		s := state{chain: []Connector{c}}
		if c.sideA == 0 {
			s.next = c.sideB
		} else {
			s.next = c.sideA
		}

		st := getStrength(s.chain)
		if maxStrength < st {
			maxStrength = st
		}

		queue = append(queue, s)
	}
	var s state

	for len(queue) > 0 {
		s, queue = queue[0], queue[1:]

		for _, c := range connectorMap[s.next] {
			if slices.Contains(s.chain, c) {
				continue
			}
			var next int
			if c.sideA == s.next {
				next = c.sideB
			} else {
				next = c.sideA
			}

			cs := slices.Clone(s.chain)

			s := state{
				chain: append(cs, c),
				next:  next,
			}

			strength := getStrength(s.chain)
			if strength > maxStrength {
				maxStrength = strength
			}

			queue = append(queue, s)
		}
	}
	counter = maxStrength

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data []Connector) {
	counter := 0

	connectorMap := map[int][]Connector{}
	for _, c := range data {
		connectorMap[c.sideA] = append(connectorMap[c.sideA], c)
		connectorMap[c.sideB] = append(connectorMap[c.sideB], c)
	}

	type state struct {
		chain []Connector
		next  int
	}

	maxStrength := 0
	maxLength := 1

	queue := []state{}
	for _, c := range connectorMap[0] {

		s := state{chain: []Connector{c}}
		if c.sideA == 0 {
			s.next = c.sideB
		} else {
			s.next = c.sideA
		}

		st := getStrength(s.chain)
		if maxStrength < st {
			maxStrength = st
		}

		queue = append(queue, s)
	}
	var s state

	for len(queue) > 0 {
		s, queue = queue[0], queue[1:]

		for _, c := range connectorMap[s.next] {
			if slices.Contains(s.chain, c) {
				continue
			}
			var next int
			if c.sideA == s.next {
				next = c.sideB
			} else {
				next = c.sideA
			}

			cs := slices.Clone(s.chain)

			s := state{
				chain: append(cs, c),
				next:  next,
			}

			strength := getStrength(s.chain)
			length := len(s.chain)
			if maxLength < length || maxLength == length && maxStrength < strength {
				maxStrength = strength
				maxLength = length
			}

			queue = append(queue, s)
		}
	}
	counter = maxStrength
	fmt.Printf("Part 2: %d\n", counter)
}
