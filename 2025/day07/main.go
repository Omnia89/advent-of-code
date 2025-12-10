package main

import (
	"fmt"
	"slices"
	"strings"

	"advent2025/util"
)

func main() {
	data := util.GetDataByRow("day07") // La soluzione è 1613
	// data := util.GetTestByRow("day07") // La soluzione è 21

	part1(data)
	part2(data)
}

func part1(rows []string) {
	counter := 0

	beamC := []int{strings.Index(rows[0], "S")}

	for beamR := 1; beamR < len(rows); beamR++ {
		newBeamC := []int{}
		for i := 0; i < len(beamC); i++ {
			if string(rows[beamR][beamC[i]]) == "^" {
				counter++
				newBeamC = append(newBeamC, beamC[i]-1)
				newBeamC = append(newBeamC, beamC[i]+1)
			} else {
				newBeamC = append(newBeamC, beamC[i])
			}
		}
		beamC = slices.Compact(newBeamC)
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(rows []string) {
	beamC := []int{strings.Index(rows[0], "S")}
	counts := make([]int, len(rows[0]))
	counts[beamC[0]] = 1

	for beamR := 1; beamR < len(rows); beamR++ {
		newBeamC := []int{}
		for i := 0; i < len(beamC); i++ {
			if string(rows[beamR][beamC[i]]) == "^" {
				newBeamC = append(newBeamC, beamC[i]-1)
				newBeamC = append(newBeamC, beamC[i]+1)

				counts[beamC[i]-1] += counts[beamC[i]]
				counts[beamC[i]+1] += counts[beamC[i]]
				counts[beamC[i]] = 0
			} else {
				newBeamC = append(newBeamC, beamC[i])
			}
		}
		beamC = slices.Compact(newBeamC)
	}

	counter := 0
	for _, c := range counts {
		counter += c
	}

	fmt.Printf("Part 2: %d\n", counter)
}
