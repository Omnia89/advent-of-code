package main

import (
	"advent2025/util"
	"fmt"
	"strings"
)

func main() {

	rawData := util.GetRawData("day02")
	// rawData := util.GetRawTest("day02")

	ranges := strings.Split(rawData, ",")

	part1(ranges)
	part2(ranges)
}

func part1(data []string) {
	counter := 0

	for _, r := range data {
		if r == "" {
			continue
		}
		
		ranges := strings.Split(r, "-")
		lower := util.ToInt(ranges[0])
		upper := util.ToInt(ranges[1])

		for i := lower; i <= upper; i++ {
			sVal := fmt.Sprintf("%d", i)
			if len(sVal) % 2 == 1 {
				continue
			}
			if sVal[:len(sVal)/2] == sVal[len(sVal)/2:] {
				counter += i
			}
		}
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data []string) {
	counter := 0

	for _, r := range data {
		if r == "" {
			continue
		}
		
		ranges := strings.Split(r, "-")
		lower := util.ToInt(ranges[0])
		upper := util.ToInt(ranges[1])

		for i := lower; i <= upper; i++ {
			sVal := fmt.Sprintf("%d", i)
			lenVal := len(sVal)
			for j := 1; j <= lenVal / 2; j++ {
				if lenVal % j != 0 {
					continue
				}
				parts := lenVal / j;
				
				ok := true
				for k := 1; k < parts; k++ {
					if sVal[0:j] != sVal[k*j:j*(k+1)] {
						ok = false
						break;
					}
				}
				if ok {
					counter += i
					break
				}
			}
		}
	}

	fmt.Printf("Part 2: %d\n", counter)
}
