package main

import (
	"fmt"
	"strings"

	"advent2025/util"
)

func main() {
	data := util.GetDataByRow("day11")
	// data := util.GetTestByRow("day11")

	paths := buildMap(data)

	part1(paths)
	part2(paths)
}

func buildMap(rows []string) map[string][]string {
	pathMap := map[string][]string{}
	for _, r := range rows {
		if strings.HasPrefix(r, "#") {
			continue
		}
		firstSplit := strings.Split(r, ":")
		key := strings.TrimSpace(firstSplit[0])
		outputs := strings.Split(firstSplit[1], " ")
		pathMap[key] = make([]string, 0, len(outputs))

		for _, o := range outputs {
			out := strings.TrimSpace(o)
			if out == "" {
				continue
			}
			pathMap[key] = append(pathMap[key], out)
		}
	}
	return pathMap
}

func numberOfPath(from string, to string, paths map[string][]string) int {
	walkingPaths := []string{from}
	allTo := false

	for !allTo {
		newWalkingPaths := []string{}
		allTo = true

		for _, key := range walkingPaths {
			if key == to {
				newWalkingPaths = append(newWalkingPaths, key)
				continue
			}
			allTo = false
			newWalkingPaths = append(newWalkingPaths, paths[key]...)
		}
		if allTo {
			break
		}

		walkingPaths = newWalkingPaths
	}

	return len(walkingPaths)
}

func countWithCache(start string, end string, paths map[string][]string, cache map[string]int) int {
	if val, ok := cache[start]; ok {
		return val
	}

	count := 0
	for _, children := range paths[start] {
		if children == end {
			cache[start] = 1
			return 1
		}
		count += countWithCache(children, end, paths, cache)
	}
	cache[start] = count
	return count
}

func countPaths(start string, end string, paths map[string][]string) int {
	cache := map[string]int{}
	return countWithCache(start, end, paths, cache)
}

func part1(paths map[string][]string) {
	counter := 0

	counter = numberOfPath("you", "out", paths)

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(paths map[string][]string) {
	counter := 0

	tmp := 0

	fmt.Printf("Start [svr->dac] ")
	tmp = countPaths("svr", "dac", paths)
	firstFlow := tmp
	fmt.Printf("DONE [%d]\n", tmp)

	fmt.Printf("Start [dac->fft] ")
	tmp = countPaths("dac", "fft", paths)
	firstFlow *= tmp
	fmt.Printf("DONE [%d]\n", tmp)

	fmt.Printf("Start [fft->out] ")
	tmp = countPaths("fft", "out", paths)
	firstFlow *= tmp
	fmt.Printf("DONE [%d]\n", tmp)

	fmt.Printf("Start [svr->fft] ")
	tmp = countPaths("svr", "fft", paths)
	secondFlow := tmp
	fmt.Printf("DONE [%d]\n", tmp)

	fmt.Printf("Start [fft->dac] ")
	tmp = countPaths("fft", "dac", paths)
	secondFlow *= tmp
	fmt.Printf("DONE [%d]\n", tmp)

	fmt.Printf("Start [dac->out] ")
	tmp = countPaths("dac", "out", paths)
	secondFlow *= tmp
	fmt.Printf("DONE [%d]\n", tmp)

	counter = firstFlow + secondFlow
	fmt.Printf("Part 2: %d\n", counter)
}
