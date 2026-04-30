package main

import (
	"fmt"
	"strings"

	"advent2018/util"
)

func main() {
	data := util.GetDataByRow("day02")
	// data := util.GetTestByRow("day02")

	part1(data)
	part2(data)
}

func count(s string) (pair bool, trio bool) {
	m := map[rune]int{}

	for _, c := range s {
		m[c]++
	}

	for _, n := range m {
		if n == 2 {
			pair = true
		} else if n == 3 {
			trio = true
		}
	}

	return
}

func part1(data []string) {
	counter := 0

	pairs := 0
	tris := 0

	for _, s := range data {
		p, t := count(s)
		if p {
			pairs++
		}
		if t {
			tris++
		}
	}

	counter = pairs * tris

	fmt.Printf("Part 1: %d\n", counter)
}

func compare(a string, b string) (bool, string) {
	var sb strings.Builder

	firstDiff := false

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			if !firstDiff {
				firstDiff = true
				continue
			} else {
				return false, ""
			}
		}
		sb.WriteByte(a[i])
	}
	return true, sb.String()
}

func part2(data []string) {
	res := ""

free:
	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			if ok, s := compare(data[i], data[j]); ok {
				res = s
				break free
			}
		}
	}

	fmt.Printf("Part 2: %s\n", res)
}
