package main

import (
	"fmt"
	"maps"
	"math"
	"slices"
	"strings"

	"advent2015/util"
)

func main() {
	data := util.GetDataByRow("day09")
	// data := util.GetTestByRow("day09")

	distances := parse(data)

	part1(distances)
	part2(distances)
}

func parse(data []string) map[string]map[string]int {
	distances := map[string]map[string]int{}

	for _, s := range data {
		cities, distance, _ := strings.Cut(s, " = ")
		a, b, _ := strings.Cut(cities, " to ")

		if _, ok := distances[a]; !ok {
			distances[a] = make(map[string]int)
		}

		distances[a][b] = util.ToInt(distance)

		if _, ok := distances[b]; !ok {
			distances[b] = make(map[string]int)
		}

		distances[b][a] = util.ToInt(distance)
	}

	return distances
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

func part1(distances map[string]map[string]int) {
	counter := 0

	cities := slices.Collect(maps.Keys(distances))

	minDist := math.MaxInt

	for list := range generatePermutation(cities) {
		totDist := 0
		for i := range len(list) - 1 {
			totDist += distances[list[i]][list[i+1]]
		}
		if totDist < minDist {
			minDist = totDist
		}
	}

	counter = minDist

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(distances map[string]map[string]int) {
	counter := 0

	cities := slices.Collect(maps.Keys(distances))

	maxDist := 0

	for list := range generatePermutation(cities) {
		totDist := 0
		for i := range len(list) - 1 {
			totDist += distances[list[i]][list[i+1]]
		}
		if totDist > maxDist {
			maxDist = totDist
		}
	}

	counter = maxDist
	fmt.Printf("Part 2: %d\n", counter)
}
