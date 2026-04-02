package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"advent2017/util"
)

func main() {
	data := util.GetDataByRow("day07")
	// data := util.GetTestByRow("day07")

	list := parse(data)

	part1(list)
	part2(list)
}

type hierarchy struct {
	name   string
	leaves []string
	weight int
}

func parse(data []string) []hierarchy {
	hs := []hierarchy{}

	for _, s := range data {
		t := strings.ReplaceAll(s, ", ", "-")
		parts := strings.Split(t, " ")
		leaves := []string{}
		if len(parts) > 2 {
			leaves = strings.Split(parts[3], "-")
		}
		hs = append(hs, hierarchy{
			name:   parts[0],
			leaves: leaves,
			weight: util.ToInt(parts[1][1 : len(parts[1])-1]),
		})
	}
	return hs
}

func part1(data []hierarchy) {
	program := ""

	leaves := []string{}

	for _, h := range data {
		leaves = append(leaves, h.leaves...)
	}

	for _, h := range data {
		if !slices.Contains(leaves, h.name) {
			program = h.name
			break
		}
	}

	fmt.Printf("Part 1: %s\n", program)
}

func buildWeightMap(treeMap map[string]hierarchy) map[string]int {
	ws := map[string]int{}

	q := slices.Collect(maps.Keys(treeMap))
	var s string

outerLoop:
	for len(q) > 0 {
		s, q = q[0], q[1:]
		h := treeMap[s]
		if len(h.leaves) == 0 {
			ws[h.name] = h.weight
			continue
		}
		c := 0
		for _, l := range h.leaves {
			v, ok := ws[l]
			if !ok {
				q = append(q, s)
				continue outerLoop
			}
			c += v
		}
		ws[s] = h.weight + c
	}
	return ws
}

func findDifferent(parentProgram string, treeMap map[string]hierarchy, weightMap map[string]int) (string, bool) {
	doubles := map[int][]string{}

	for _, s := range treeMap[parentProgram].leaves {
		doubles[weightMap[s]] = append(doubles[weightMap[s]], treeMap[s].name)
	}

	for _, ls := range doubles {
		if len(ls) == 1 {
			if len(treeMap[parentProgram].leaves) == 2 {
				panic("to implement")
			}
			return ls[0], true
		}
	}
	return "", false
}

func printHierarchy(program string, treeMap map[string]hierarchy, weightMap map[string]int) {
	var sb strings.Builder

	fmt.Fprintf(&sb, "%s [%d] -> ", program, weightMap[program])

	for _, s := range treeMap[program].leaves {
		fmt.Fprintf(&sb, " %s[%d] ", s, weightMap[s])
	}

	fmt.Println(sb.String())
}

func part2(data []hierarchy) {
	counter := 0

	leaves := []string{}
	treeMap := map[string]hierarchy{}
	current := ""

	for _, h := range data {
		leaves = append(leaves, h.leaves...)
		treeMap[h.name] = h
	}

	for _, h := range data {
		if !slices.Contains(leaves, h.name) {
			current = h.name
			break
		}
	}

	weightMap := buildWeightMap(treeMap)

	previous := ""
	var double bool
	for {
		// printHierarchy(current, treeMap, weightMap)
		tmp := current
		current, double = findDifferent(current, treeMap, weightMap)
		if !double {
			current = tmp
			break
		}
		previous = tmp
	}

	sibling := ""
	for _, k := range treeMap[previous].leaves {
		if k != current {
			sibling = k
			break
		}
	}

	fmt.Printf("c[%s] p[%s]\n", current, sibling)

	delta := weightMap[sibling] - weightMap[current]

	counter = treeMap[current].weight + delta

	// TODO: trovo la differenza di peso nel previous, e lo sottraggo al valore di peso del singolo nodo previous

	fmt.Printf("Part 2: %d\n", counter)
}
