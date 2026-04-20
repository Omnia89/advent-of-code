package main

import (
	"fmt"
	"slices"
	"strings"

	"advent2017/util"
)

func main() {
	data := util.GetRawData("day16")
	// data := util.GetRawTest("day16")

	list := strings.Split(strings.TrimSpace(data), ",")

	part1(list)
	part2(list)
}

func execute(step string, programs *[]string) {
	switch step[0] {
	case 's':
		n := len(*programs) - util.ToInt(step[1:])
		*programs = append((*programs)[n:], (*programs)[:n]...)
	case 'x':
		idx := util.StringToIntSlice(step[1:], "/")
		a := idx[0]
		b := idx[1]
		(*programs)[a], (*programs)[b] = (*programs)[b], (*programs)[a]
	case 'p':
		aS, bS, _ := strings.Cut(step[1:], "/")
		a := slices.Index(*programs, aS)
		b := slices.Index(*programs, bS)
		// fmt.Printf("s[%s] aS[%s] bS[%s] a[%d], b[%d] %v\n", step, aS, bS, a, b, *programs)
		(*programs)[a], (*programs)[b] = (*programs)[b], (*programs)[a]
	}
}

func part1(data []string) {
	result := ""

	programs := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"}
	// programs := []string{"a", "b", "c", "d", "e"}

	for _, s := range data {
		execute(s, &programs)
		// fmt.Printf("s[%s] %v\n", s, programs)
	}

	result = strings.Join(programs, "")

	fmt.Printf("Part 1: %s\n", result)
}

func part2(data []string) {
	result := ""

	programs := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"}

	prev := strings.Join(programs, "")
	cache := map[string][]string{}
	for range 1_000_000_000 {
		if r, ok := cache[prev]; ok {
			prev = strings.Join(r, "")
			programs = r
			continue
		}

		for _, s := range data {
			execute(s, &programs)
		}
		cache[prev] = programs
		prev = strings.Join(programs, "")
	}
	result = strings.Join(programs, "")
	fmt.Printf("Part 2: %s\n", result)
}
