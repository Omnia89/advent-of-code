package main

import (
	"fmt"
	"strconv"
	"strings"

	"advent2024/util"
)

func main() {
	// data := util.GetTestByRow("day25")
	data := util.GetDataByRow("day25")

	locks, keys := parse(data)

	part1(locks, keys)
	part2(locks, keys)
}

type Problem struct {
	Towels   []string
	Patterns []string
}

func parse(data []string) (locks [][]int, keys [][]int) {
	locks = [][]int{}
	keys = [][]int{}

	for i := 0; i < len(data); i += 8 {
		isKey := data[i] == "....."
		buffer := make([]int, 5)

		for j := range 5 {
			temp := util.StringToIntSlice(strings.ReplaceAll(strings.ReplaceAll(data[i+j+1], ".", "0"), "#", "1"), "")
			for k, v := range temp {
				buffer[k] += v
			}
		}
		if isKey {
			keys = append(keys, buffer)
		} else {
			locks = append(locks, buffer)
		}

	}
	return
}

func pairId(lock []int, key []int) string {
	var sb strings.Builder
	for _, l := range lock {
		sb.WriteString(strconv.Itoa(l))
	}
	sb.WriteString("-")
	for _, k := range key {
		sb.WriteString(strconv.Itoa(k))
	}

	return sb.String()
}

func sum(a, b []int) []int {
	c := make([]int, 5)
	for i := range 5 {
		c[i] = a[i] + b[i]
	}
	return c
}

func check(a []int, limit int) bool {
	for _, i := range a {
		if i > limit {
			return false
		}
	}
	return true
}

func part1(locks [][]int, keys [][]int) {
	counter := 0

	// fmt.Printf("locks %v\n", locks)
	// fmt.Printf("keys %v\n", keys)

	uniquePairs := map[string]bool{}
	for _, k := range keys {
		for _, l := range locks {
			s := sum(k, l)
			if !check(s, 5) {
				continue
			}
			// fmt.Printf(" k[%v] l[%v] s[%v]\n", k, l, s)
			uniquePairs[pairId(l, k)] = true
		}
	}

	counter = len(uniquePairs)

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(locks [][]int, keys [][]int) {
	counter := 0

	fmt.Printf("Part 2: %d\n", counter)
}
