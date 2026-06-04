package main

import (
	"fmt"
	"slices"
	"strings"

	"advent2018/util"
)

func main() {
	data := util.GetDataByRow("day12")
	//data := util.GetTestByRow("day12")

	s, rs := parse(data)

	part1(s, rs)
	part2(s, rs)
}

type Rule struct {
	pattern []int
	value   int
}

func (r Rule) match(a []int) (int, bool) {
	for i := range 5 {
		if r.pattern[i] != a[i] {
			return 0, false
		}
	}
	return r.value, true
}

func parse(data []string) ([]int, []Rule) {
	rs := []Rule{}

	parts := strings.Split(data[0], " ")
	initialState := util.StringToIntSlice(strings.ReplaceAll(strings.ReplaceAll(parts[2], "#", "1"), ".", "0"), "")

	for i := 2; i < len(data); i++ {
		rParts := strings.Split(data[i], " ")
		pt := util.StringToIntSlice(strings.ReplaceAll(strings.ReplaceAll(rParts[0], "#", "1"), ".", "0"), "")
		val := util.ToInt(strings.ReplaceAll(strings.ReplaceAll(rParts[2], "#", "1"), ".", "0"))

		rs = append(rs, Rule{pt, val})
	}
	return initialState, rs
}

func getSection(state []int, index int) []int {
	r := make([]int, 0, 5)
	i := index - 2
	for i < 0 {
		r = append(r, 0)
		i++
	}
	for k := i; k <= index+2; k++ {
		if k >= len(state) {
			r = append(r, 0)
		} else {
			r = append(r, state[k])
		}
	}

	return r
}

func printState(i int, state []int) {
	var sb strings.Builder
	fmt.Fprintf(&sb, "[%d]: ", i)

	for _, v := range state {
		c := "."
		if v == 1 {
			c = "#"
		}
		fmt.Fprintf(&sb, "%s", c)
	}
	fmt.Println(sb.String())
}

func part1(state []int, rules []Rule) {
	counter := 0

	steps := 20

	// add to both sides the maximun number of elements that the rules can extent
	state = append(slices.Repeat([]int{0}, steps*2), state...)
	state = append(state, slices.Repeat([]int{0}, steps*2)...)
	deltaIndex := steps * 2

	// printState(-1, state)
	for range steps {
		newState := slices.Clone(state)
		for i := range state {
			section := getSection(state, i)
			matched := false
			for _, r := range rules {
				if v, ok := r.match(section); ok {
					newState[i] = v
					matched = true
					break
				}
			}
			if !matched {
				newState[i] = 0
			}
		}
		// printState(n, newState)
		state = newState
	}

	for i, v := range state {
		idx := i - deltaIndex
		if v == 1 {
			counter += idx
		}
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func trimState(s []int) ([]int, int) {
	start := 0
	for start < len(s) && s[start] == 0 {
		start++
	}
	end := len(s)
	for end > start && s[end-1] == 0 {
		end--
	}
	if start >= end {
		return []int{}, 0
	}
	return s[start:end], start
}

func runStep(state []int, rules []Rule) []int {
	padded := append([]int{0, 0, 0, 0}, state...)
	padded = append(padded, 0, 0, 0, 0)
	newState := slices.Clone(padded)
	for i := range padded {
		section := getSection(padded, i)
		matched := false
		for _, r := range rules {
			if v, ok := r.match(section); ok {
				newState[i] = v
				matched = true
				break
			}
		}
		if !matched {
			newState[i] = 0
		}
	}
	return newState
}

func part2(state []int, rules []Rule) {
	const steps = 50_000_000_000

	trimmed, delta := trimState(state)
	state = slices.Clone(trimmed)
	offset := delta

	for s := 0; s < steps; s++ {
		raw := runStep(state, rules)
		trimmed, delta := trimState(raw)
		newOffset := offset - 4 + delta

		if slices.Equal(trimmed, state) {
			shiftPerStep := newOffset - offset
			remaining := steps - s - 1
			finalOffset := newOffset + shiftPerStep*remaining
			counter := 0
			for i, v := range trimmed {
				if v == 1 {
					counter += i + finalOffset
				}
			}
			fmt.Printf("Part 2: %d\n", counter)
			return
		}

		state = slices.Clone(trimmed)
		offset = newOffset
	}

	// fallback: no cycle found
	counter := 0
	for i, v := range state {
		if v == 1 {
			counter += i + offset
		}
	}
	fmt.Printf("Part 2: %d\n", counter)
}
