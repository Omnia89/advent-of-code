package main

import (
	"fmt"
	"strings"

	"advent2024/util"
)

func main() {
	// data := util.GetTestByRow("day19")
	data := util.GetDataByRow("day19")

	problem := parse(data)

	part1(problem)
	part2(problem)
}

type Problem struct {
	Towels   []string
	Patterns []string
}

func parse(data []string) Problem {
	towels := strings.Split(strings.ReplaceAll(data[0], " ", ""), ",")

	patterns := []string{}
	for i := 2; i < len(data); i++ {
		patterns = append(patterns, data[i])
	}

	return Problem{
		towels,
		patterns,
	}
}

func checkTowel(pattern string, towel string, index int) bool {
	if index+len(towel) > len(pattern) {
		return false
	}

	for i := index; i < index+len(towel); i++ {
		if pattern[i] != towel[i-index] {
			return false
		}
	}
	return true
}

func solve(pattern string, towelMap map[string][]string) bool {
	ts := towelMap[string(pattern[0])]
	for _, t := range ts {
		if checkTowel(pattern, t, 0) {
			if len(pattern)-len(t) == 0 {
				return true
			}
			check := solve(pattern[len(t):], towelMap)
			if check {
				return true
			}
		}
	}
	return false
}

func solveNum(pattern string, towelMap map[string][]string, cache map[string]int) int {
	if v, ok := cache[pattern]; ok {
		return v
	}
	ts := towelMap[string(pattern[0])]
	num := 0
	for _, t := range ts {
		if checkTowel(pattern, t, 0) {
			if len(pattern)-len(t) == 0 {
				num++
			} else {
				num += solveNum(pattern[len(t):], towelMap, cache)
			}
		}
	}
	cache[pattern] = num
	return num
}

func part1(problem Problem) {
	counter := 0

	towelMap := map[string][]string{}

	for _, t := range problem.Towels {
		towelMap[string(t[0])] = append(towelMap[string(t[0])], t)
	}

	for _, p := range problem.Patterns {
		solved := solve(p, towelMap)
		if solved {
			counter++
		}
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(problem Problem) {
	counter := 0

	towelMap := map[string][]string{}

	for _, t := range problem.Towels {
		towelMap[string(t[0])] = append(towelMap[string(t[0])], t)
	}

	cache := map[string]int{}
	for _, p := range problem.Patterns {
		num := solveNum(p, towelMap, cache)
		counter += num
	}
	fmt.Printf("Part 2: %d\n", counter)
}
