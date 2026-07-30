package main

import (
	"fmt"
	"maps"
	"strings"

	"advent2019/util"
)

func main() {
	data := util.GetDataByRow("day06")
	// data := util.GetTestByRow("day06")

	orbits := parse(data)

	part1(orbits)
	part2(orbits)
}

func parse(data []string) map[string][]string {
	orbits := map[string][]string{}

	for _, s := range data {
		center, satellite, _ := strings.Cut(s, ")")

		orbits[center] = append(orbits[center], satellite)
	}

	return orbits
}

func part1(orbits map[string][]string) {
	counter := 0

	orbitsValue := map[string]int{
		"COM": 0,
	}

	fronts := []string{"COM"}

	for len(fronts) > 0 {
		newFronts := []string{}

		for _, f := range fronts {
			os := orbits[f]
			newFronts = append(newFronts, os...)

			for _, s := range os {
				orbitsValue[s] = orbitsValue[f] + 1
			}
		}

		fronts = newFronts
	}

	for _, n := range orbitsValue {
		counter += n
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(orbits map[string][]string) {
	counter := 0

	start := ""

	paths := maps.Clone(orbits)
	for c, ss := range orbits {
		for _, s := range ss {
			paths[s] = append(paths[s], c)
			if s == "YOU" {
				start = c
			}
		}
	}

	target := "SAN"

	visited := map[string]bool{
		"YOU": true,
	}
	distances := map[string]int{
		"YOU": 0,
		start: 0,
	}

	queue := []string{start}
	var s string

	for len(queue) > 0 {
		s, queue = queue[0], queue[1:]

		currentDistance := distances[s]
		for _, p := range paths[s] {
			if visited[p] {
				continue
			}
			if p == target {
				counter = currentDistance
				break
			}
			visited[p] = true
			distances[p] = currentDistance + 1
			queue = append(queue, p)
		}
	}

	fmt.Printf("Part 2: %d\n", counter)
}
