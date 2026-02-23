package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"advent2015/util"
)

func main() {
	data := util.GetDataByRow("day13")
	// data := util.GetTestByRow("day13")

	happiness := parse(data)

	part1(happiness)
	part2(happiness)
}

func parse(data []string) map[string]map[string]int {
	happiness := map[string]map[string]int{}

	for _, s := range data {
		pieces := strings.Split(s, " ")
		val := util.ToInt(pieces[3])
		if pieces[2] == "lose" {
			val = -val
		}

		if _, ok := happiness[pieces[0]]; !ok {
			happiness[pieces[0]] = map[string]int{}
		}

		happiness[pieces[0]][strings.Trim(pieces[10], ".")] = val
	}

	return happiness
}

func generatePermutation(values []string) <-chan []string {
	ch := make(chan []string)

	go func() {
		defer close(ch)
		if len(values) == 0 {
			return
		}

		// make a copy of values
		arr := make([]string, len(values))
		copy(arr, values)

		permutationStep(arr, 0, ch)
	}()

	return ch
}

func permutationStep(values []string, start int, ch chan<- []string) {
	if start == len(values) {
		// finished permutation, send over channel
		res := make([]string, len(values))
		copy(res, values)
		ch <- res
		return
	}

	for i := start; i < len(values); i++ {
		values[start], values[i] = values[i], values[start]
		permutationStep(values, start+1, ch)
		values[start], values[i] = values[i], values[start]
	}
}

func part1(happiness map[string]map[string]int) {
	counter := 0

	people := slices.Collect(maps.Keys(happiness))

	maxHappiness := 0

	for list := range generatePermutation(people) {
		totH := 0
		for i := range len(list) {
			p1 := list[i]
			p2 := list[(i+1)%len(list)]
			totH += happiness[p1][p2] + happiness[p2][p1]
		}
		if totH > maxHappiness {
			maxHappiness = totH
		}
	}

	counter = maxHappiness

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(happiness map[string]map[string]int) {
	counter := 0

	people := slices.Collect(maps.Keys(happiness))

	happiness["you"] = map[string]int{}
	for _, k := range people {
		happiness["you"][k] = 0
		happiness[k]["you"] = 0
	}

	people = append(people, "you")

	maxHappiness := 0

	for list := range generatePermutation(people) {
		totH := 0
		for i := range len(list) {
			p1 := list[i]
			p2 := list[(i+1)%len(list)]
			totH += happiness[p1][p2] + happiness[p2][p1]
		}
		if totH > maxHappiness {
			maxHappiness = totH
		}
	}

	counter = maxHappiness
	fmt.Printf("Part 2: %d\n", counter)
}
