package main

import (
	"fmt"
	"slices"
	"strings"

	"advent2018/util"
)

func main() {
	data := util.GetDataByRow("day07")
	// data := util.GetTestByRow("day07")

	deps := parse(data)
	deps2 := parse(data)

	part1(deps)
	part2(deps2)
}

func parse(data []string) map[string][]string {
	deps := map[string][]string{}

	for _, s := range data {
		parts := strings.Split(s, " ")

		if _, ok := deps[parts[1]]; !ok {
			deps[parts[1]] = []string{}
		}

		deps[parts[7]] = append(deps[parts[7]], parts[1])
	}

	return deps
}

func remove(arr []string, k string) []string {
	n := make([]string, 0, len(arr)-1)
	for _, s := range arr {
		if s != k {
			n = append(n, s)
		}
	}
	return n
}

func part1(data map[string][]string) {
	var sb strings.Builder

	queue := []string{}
	for k, s := range data {
		if len(s) == 0 {
			queue = append(queue, k)
		}
	}

	slices.Sort(queue)
	var c string

	for len(queue) > 0 {
		c, queue = queue[0], queue[1:]
		sb.WriteString(c)

		for k, r := range data {
			if slices.Contains(r, c) {
				if len(r) == 1 {
					queue = append(queue, k)
				}
				data[k] = remove(r, c)
			}
		}
		slices.Sort(queue)
	}

	fmt.Printf("Part 1: %s\n", sb.String())
}

var timeString string = "-ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func getTime(s string) int {
	GAP := 60
	return GAP + strings.Index(timeString, s)
}

func printW(t int, nodes []string, times []int) {
	var sb strings.Builder

	fmt.Fprintf(&sb, "- [%d]", t)

	for i := range nodes {
		fmt.Fprintf(&sb, "\t[%s][%d]", nodes[i], times[i])
	}
	fmt.Println(sb.String())
}

func part2(data map[string][]string) {
	counter := 0

	WORKERS := 5

	workerNode := make([]string, WORKERS)
	workerTime := make([]int, WORKERS)

	queue := []string{}
	for k, s := range data {
		if len(s) == 0 {
			queue = append(queue, k)
		}
	}

	slices.Sort(queue)
	var c string

	wip := true

	check := func() bool {
		if len(queue) > 0 {
			return false
		}

		for _, n := range workerNode {
			if n != "" {
				return false
			}
		}
		return true
	}

	for wip {

		for i := range len(workerNode) {
			if workerNode[i] == "" {
				continue
			}
			workerTime[i] -= 1
			if workerTime[i] == 0 {
				for k, r := range data {
					if slices.Contains(r, workerNode[i]) {
						if len(r) == 1 {
							queue = append(queue, k)
						}
						data[k] = remove(r, workerNode[i])
					}
				}
				workerNode[i] = ""
				slices.Sort(queue)
			}
		}

		for i := range len(workerNode) {
			if len(queue) == 0 {
				break
			}
			if workerNode[i] != "" {
				continue
			}
			c, queue = queue[0], queue[1:]
			workerNode[i] = c
			workerTime[i] = getTime(c)
		}
		// printW(counter, workerNode, workerTime)
		if check() {
			break
		}
		counter++
	}
	fmt.Printf("Part 2: %d\n", counter)
}
