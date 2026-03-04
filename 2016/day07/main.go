package main

import (
	"fmt"
	"regexp"
	"strings"

	"advent2016/util"
)

func main() {
	data := util.GetDataByRow("day07")
	// data := util.GetTestByRow("day07")

	part1(data)
	part2(data)
}

func part1(data []string) {
	counter := 0

	for _, s := range data {
		inside := false
		found := false
		for i := 0; i < len(s)-3; i++ {
			if s[i] == ']' {
				inside = false
				continue
			}
			if s[i] == '[' {
				inside = true
				continue
			}

			if s[i] != s[i+1] && s[i] == s[i+3] && s[i+1] == s[i+2] {
				if !inside {
					found = true
				} else {
					found = false
					break
				}
			}
		}
		if found {
			counter++
		}
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data []string) {
	counter := 0

	for _, s := range data {
		triplets := []string{}

		inside := false
		for i := 0; i < len(s)-2; i++ {
			if s[i] == ']' {
				inside = false
				continue
			}
			if s[i] == '[' {
				inside = true
			}
			if inside {
				continue
			}

			if s[i] != s[i+1] && s[i] == s[i+2] {
				a := rune(s[i])
				b := rune(s[i+1])
				triplets = append(triplets, fmt.Sprintf("%c%c%c", b, a, b))
			}
		}

		hyperRe := regexp.MustCompile(`\[(.*?)\]`)
		matches := hyperRe.FindAllStringSubmatch(s, -1)
		parts := []string{}
		for _, m := range matches {
			parts = append(parts, m[1])
		}
	externalLoop:
		for _, p := range parts {
			for _, t := range triplets {
				if strings.Contains(p, t) {
					counter++
					break externalLoop
				}
			}
		}
	}

	fmt.Printf("Part 2: %d\n", counter)
}
