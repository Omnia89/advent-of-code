package main

import (
	"fmt"
	"slices"
	"strings"

	"advent2024/util"
)

func main() {
	// data := util.GetTestByRow("day23")
	data := util.GetDataByRow("day23")

	joins := parse(data)

	part1(joins)
	part2(joins)
}

// alphabetical ordered
type join struct {
	first  string
	second string
}

func parse(data []string) []join {
	joins := []join{}

	for _, d := range data {
		nodes := strings.Split(d, "-")
		slices.Sort(nodes)
		joins = append(joins, join{nodes[0], nodes[1]})
	}
	return joins
}

func getLinkMap(joins []join) map[string][]string {
	links := map[string][]string{}

	for _, j := range joins {
		if _, ok := links[j.first]; !ok {
			links[j.first] = []string{}
		}
		if _, ok := links[j.second]; !ok {
			links[j.second] = []string{}
		}

		links[j.first] = append(links[j.first], j.second)
		links[j.second] = append(links[j.second], j.first)
	}
	return links
}

func part1(joins []join) {
	counter := 0

	links := getLinkMap(joins)

	unique := map[string]bool{}

	addUnique := func(a, b, c string) {
		temp := []string{a, b, c}
		slices.Sort(temp)
		unique[fmt.Sprintf("%s-%s-%s", temp[0], temp[1], temp[2])] = true
	}

	for k, l := range links {
		if !strings.HasPrefix(k, "t") {
			continue
		}

		for _, middle := range l {
			for _, third := range links[middle] {
				if third == k {
					continue
				}
				if slices.Contains(l, third) {
					addUnique(k, middle, third)
				}
			}
		}
	}
	counter = len(unique)

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(joins []join) {
	counter := 0

	fmt.Printf("Part 2: %d\n", counter)
}
