package main

import (
	"fmt"
	"strings"

	"advent2015/util"
)

func main() {
	data := util.GetDataByRow("day14")
	// data := util.GetTestByRow("day14")

	list := parse(data)

	part1(list)
	part2(list)
}

type reindeer struct {
	name     string
	km       int
	duration int
	resting  int
}

func parse(data []string) []reindeer {
	rs := []reindeer{}

	for _, s := range data {
		pieces := strings.Split(s, " ")
		rs = append(rs, reindeer{
			pieces[0],
			util.ToInt(pieces[3]),
			util.ToInt(pieces[6]),
			util.ToInt(pieces[13]),
		})
	}

	return rs
}

func part1(data []reindeer) {
	counter := 0

	seconds := 2503

	maxDistance := 0

	for _, r := range data {
		distance := 0
		completeRounds := seconds / (r.duration + r.resting)

		distance += completeRounds * r.duration * r.km

		remaining := seconds - completeRounds*(r.duration+r.resting)

		remainingSec := util.IntMin(remaining, r.duration)

		distance += remainingSec * r.km

		if distance > maxDistance {
			maxDistance = distance
		}
	}

	counter = maxDistance

	fmt.Printf("Part 1: %d\n", counter)
}

func getMax(values map[string]int) []string {
	maxVal := 0
	byPoints := map[int][]string{}

	for k, v := range values {
		if _, ok := byPoints[v]; !ok {
			byPoints[v] = []string{}
		}
		byPoints[v] = append(byPoints[v], k)
		if v > maxVal {
			maxVal = v
		}
	}

	return byPoints[maxVal]
}

func part2(data []reindeer) {
	counter := 0

	points := map[string]int{}
	kms := map[string]int{}

	for _, r := range data {
		points[r.name] = 0
		kms[r.name] = 0
	}

	seconds := 2503

	for tick := range seconds {
		for _, r := range data {
			relativeTick := tick % (r.duration + r.resting)
			if relativeTick < r.duration {
				kms[r.name] += r.km
			}
		}
		// fmt.Printf(" [%d] [%v]\n", tick, kms)
		for _, k := range getMax(kms) {
			points[k] += 1
		}
		// fmt.Printf("  - [%v]\n", points)
	}

	for _, k := range getMax(points) {
		counter = points[k]
		break
	}

	fmt.Printf("Part 2: %d\n", counter)
}
