package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"advent2017/util"
)

func main() {
	data := util.GetDataByRow("day12")
	// data := util.GetTestByRow("day12")

	paths := parse(data)

	part1(paths)
	part2(paths)
}

func parse(data []string) map[string][]string {
	paths := map[string][]string{}

	for _, str := range data {
		s := strings.ReplaceAll(str, " ", "")
		idx, others, _ := strings.Cut(s, "<->")
		if _, ok := paths[idx]; !ok {
			paths[idx] = []string{}
		}

		for o := range strings.SplitSeq(others, ",") {
			if o != idx {
				paths[idx] = append(paths[idx], o)
			}
		}
	}
	return paths
}

func part1(data map[string][]string) {
	counter := 0

	done := []string{"0"}

	var s string
	q := []string{"0"}

	for len(q) > 0 {
		s, q = q[0], q[1:]

		for _, n := range data[s] {
			if !slices.Contains(done, n) {
				done = append(done, n)
				q = append(q, n)
			}
		}
	}

	counter = len(done)

	fmt.Printf("Part 1: %d\n", counter)
}

func removeFromArray(v []string, el string) []string {
	return slices.DeleteFunc(v, func(e string) bool {
		return e == el
	})
}

func part2(data map[string][]string) {
	counter := 0

	toDo := slices.Collect(maps.Keys(data))

	var current string
	for len(toDo) > 0 {
		current, toDo = toDo[0], toDo[1:]
		counter++

		var s string
		q := []string{current}
		for len(q) > 0 {
			s, q = q[0], q[1:]

			for _, n := range data[s] {
				if slices.Contains(toDo, n) {
					toDo = removeFromArray(toDo, n)
					q = append(q, n)
				}
			}
		}
	}

	fmt.Printf("Part 2: %d\n", counter)
}
