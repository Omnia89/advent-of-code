package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"advent2016/util"
)

func main() {
	data := util.GetDataByRow("day11")
	// data := util.GetTestByRow("day11")

	st := parse(data)

	part1(st)
	part2(st)
}

type state struct {
	steps         int
	elevatorFloor int
	floorObjs     map[int][]string
}

func (s state) key() string {
	positions := make(map[string][2]int)

	for f, objs := range s.floorObjs {
		for _, o := range objs {
			mat, typeObj, _ := strings.Cut(o, "-")
			pos := positions[mat]
			if typeObj == "generator" {
				pos[0] = f
			} else {
				pos[1] = f
			}
			positions[mat] = pos
		}
	}

	pairs := [][]int{}
	for _, p := range positions {
		pairs = append(pairs, []int{p[0], p[1]})
	}

	slices.SortFunc(pairs, func(a, b []int) int {
		if a[0] != b[0] {
			return a[0] - b[0]
		}
		return a[1] - b[1]
	})

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("E:%d", s.elevatorFloor))
	for _, p := range pairs {
		sb.WriteString(fmt.Sprintf("|%d,%d", p[0], p[1]))
	}

	return sb.String()
}

func parse(data []string) state {
	objs := map[int][]string{
		1: {},
		2: {},
		3: {},
		4: {},
	}

	dict := map[string]int{
		"first":  1,
		"second": 2,
		"third":  3,
		"fourth": 4,
	}

	for _, s := range data {
		t := strings.ReplaceAll(strings.ReplaceAll(s, ",", ""), "and ", "")
		parts := strings.Split(t, " ")
		if parts[4] == "nothing" {
			continue
		}
		fl := dict[parts[1]]
		for i := 5; i < len(parts); i += 3 {
			mat, _, ok := strings.Cut(parts[i], "-")

			suffix := "generator"
			if ok {
				suffix = "chip"
			}
			objs[fl] = append(objs[fl], fmt.Sprintf("%s-%s", mat, suffix))
		}
	}
	return state{
		steps:         0,
		elevatorFloor: 1,
		floorObjs:     objs,
	}
}

func getCombo(objs ...string) [][]string {
	res := [][]string{}

	for i := range objs {
		res = append(res, []string{objs[i]})
		for j := i + 1; j < len(objs); j++ {
			res = append(res, []string{objs[i], objs[j]})
		}
	}

	return res
}

func remove(a []string, d []string) []string {
	r := []string{}
	for _, s := range a {
		if !slices.Contains(d, s) {
			r = append(r, s)
		}
	}
	return r
}

func check(el []string, fl []string) bool {
	objs := append(fl, el...)

	chips := []string{}
	generators := []string{}

	for _, o := range objs {
		m, t, _ := strings.Cut(o, "-")
		if t == "chip" {
			chips = append(chips, m)
		} else {
			generators = append(generators, m)
		}
	}

	if len(chips) == 0 || len(generators) == 0 {
		return true
	}

	for _, c := range chips {
		if !slices.Contains(generators, c) {
			return false
		}
	}
	return true
}

func part1(problem state) {
	counter := 0

	q := []state{problem}
	alreadySeen := map[string]struct{}{}

	var s state

	minSteps := math.MaxInt

	for len(q) > 0 {
		s, q = q[0], q[1:]

		// when all objs are in floor 4, drop, and keep track of lower steps
		if len(s.floorObjs[1])+len(s.floorObjs[2])+len(s.floorObjs[3]) == 0 {
			if s.steps < minSteps {
				minSteps = s.steps
				fmt.Println(s.steps)
				continue
			}
		}

		// if steps are greater of lower steps, drop
		if s.steps > minSteps {
			continue
		}

		directions := []int{}
		if s.elevatorFloor < 4 {
			directions = append(directions, 1)
		}
		if s.elevatorFloor > 1 {
			allEmpty := true
			for i := range s.elevatorFloor {
				if len(s.floorObjs[i]) > 0 {
					allEmpty = false
				}
			}
			if !allEmpty {
				directions = append(directions, -1)
			}
		}

		combos := getCombo(s.floorObjs[s.elevatorFloor]...)

		for _, d := range directions {
			moved := 0
			if d == 1 {
				slices.SortFunc(combos, func(a, b []string) int {
					return len(b) - len(a)
				})
			} else {
				moved = 3
				slices.SortFunc(combos, func(a, b []string) int {
					return len(a) - len(b)
				})
			}
			for _, c := range combos {
				if d == 1 && len(c) < moved || d == -1 && len(c) > moved {
					break
				}
				floorRemains := remove(s.floorObjs[s.elevatorFloor], c)
				// check if the remaing obj in floor are compatible, else drop
				if !check(nil, floorRemains) {
					continue
				}
				// check if the elevator obj and the target floor are compatible, else drop
				if !check(c, s.floorObjs[s.elevatorFloor+d]) {
					continue
				}
				// check the elevator objs
				if !check(c, nil) {
					continue
				}
				moved = len(c)
				floors := make(map[int][]string)
				for k, v := range s.floorObjs {
					floors[k] = slices.Clone(v)
				}
				floors[s.elevatorFloor] = floorRemains
				floors[s.elevatorFloor+d] = append(floors[s.elevatorFloor+d], c...)
				nS := state{
					s.steps + 1,
					s.elevatorFloor + d,
					floors,
				}

				if _, ok := alreadySeen[nS.key()]; ok {
					continue
				}
				alreadySeen[nS.key()] = struct{}{}

				q = append(q, nS)
			}
		}
	}
	counter = minSteps

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(problem state) {
	counter := 0

	ffObjs := problem.floorObjs[1]
	ffObjs = append(ffObjs, []string{"elerium-chip", "elerium-generator", "dilithium-chip", "dilithium-generator"}...)
	problem.floorObjs[1] = ffObjs

	q := []state{problem}
	alreadySeen := map[string]struct{}{}

	var s state

	minSteps := math.MaxInt

	for len(q) > 0 {
		s, q = q[0], q[1:]

		// when all objs are in floor 4, drop, and keep track of lower steps
		if len(s.floorObjs[1])+len(s.floorObjs[2])+len(s.floorObjs[3]) == 0 {
			if s.steps < minSteps {
				minSteps = s.steps
				fmt.Println(s.steps)
				continue
			}
		}

		// if steps are greater of lower steps, drop
		if s.steps > minSteps {
			continue
		}

		directions := []int{}
		if s.elevatorFloor < 4 {
			directions = append(directions, 1)
		}
		if s.elevatorFloor > 1 {
			allEmpty := true
			for i := range s.elevatorFloor {
				if len(s.floorObjs[i]) > 0 {
					allEmpty = false
				}
			}
			if !allEmpty {
				directions = append(directions, -1)
			}
		}

		combos := getCombo(s.floorObjs[s.elevatorFloor]...)

		for _, d := range directions {
			moved := 0
			if d == 1 {
				slices.SortFunc(combos, func(a, b []string) int {
					return len(b) - len(a)
				})
			} else {
				moved = 3
				slices.SortFunc(combos, func(a, b []string) int {
					return len(a) - len(b)
				})
			}
			for _, c := range combos {
				if d == 1 && len(c) < moved || d == -1 && len(c) > moved {
					break
				}
				floorRemains := remove(s.floorObjs[s.elevatorFloor], c)
				// check if the remaing obj in floor are compatible, else drop
				if !check(nil, floorRemains) {
					continue
				}
				// check if the elevator obj and the target floor are compatible, else drop
				if !check(c, s.floorObjs[s.elevatorFloor+d]) {
					continue
				}
				// check the elevator objs
				if !check(c, nil) {
					continue
				}
				moved = len(c)
				floors := make(map[int][]string)
				for k, v := range s.floorObjs {
					floors[k] = slices.Clone(v)
				}
				floors[s.elevatorFloor] = floorRemains
				floors[s.elevatorFloor+d] = append(floors[s.elevatorFloor+d], c...)
				nS := state{
					s.steps + 1,
					s.elevatorFloor + d,
					floors,
				}

				if _, ok := alreadySeen[nS.key()]; ok {
					continue
				}
				alreadySeen[nS.key()] = struct{}{}

				q = append(q, nS)
			}
		}
	}
	counter = minSteps

	fmt.Printf("Part 2: %d\n", counter)
}
