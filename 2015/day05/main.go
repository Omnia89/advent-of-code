package main

import (
	"fmt"
	"strings"

	"advent2015/util"
)

func main() {
	data := util.GetDataByRow("day05")
	// data := util.GetTestByRow("day05")

	part1(data)
	part2(data)
}

func removeVowel(s string) string {
	n := s
	vow := []string{"a", "e", "i", "o", "u"}
	for _, v := range vow {
		n = strings.ReplaceAll(n, v, "")
	}
	return n
}

func part1(data []string) {
	counter := 0

	badSequence := []string{"ab", "cd", "pq", "xy"}

dataLoop:
	for _, s := range data {
		for _, b := range badSequence {
			if strings.Contains(s, b) {
				continue dataLoop
			}
		}

		nVowels := len(s) - len(removeVowel(s))
		if nVowels < 3 {
			continue
		}

		foundDouble := false

		last := s[0]
		for i := 1; i < len(s); i++ {
			if s[i] == last {
				foundDouble = true
				break
			}
			last = s[i]
		}

		if foundDouble {
			counter++
		}
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data []string) {
	counter := 0

	for _, s := range data {
		foundDouble := false

		for i := 0; i < len(s)-1; i++ {
			check := s[i : i+2]
			if strings.Contains(s[i+2:], check) {
				foundDouble = true
				break
			}
		}

		if !foundDouble {
			continue
		}

		foundLeap := false

		for i := range len(s) - 2 {
			if s[i] == s[i+2] {
				foundLeap = true
				continue
			}
		}
		if foundLeap {
			counter++
		}
	}

	fmt.Printf("Part 2: %d\n", counter)
}
