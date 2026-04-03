package main

import (
	"fmt"

	"advent2017/util"
)

func main() {
	data := util.GetRawData("day09")
	// data := util.GetRawTest("day09")

	part1(data)
	part2(data)
}

func part1(data string) {
	counter := 0

	inGarbage := false
	escaping := false
	actualScore := 1

	for _, c := range data {
		if escaping {
			escaping = false
			continue
		}
		if c == '!' {
			escaping = true
			continue
		}
		if c == '<' {
			inGarbage = true
			continue
		}
		if inGarbage {
			if c == '>' {
				inGarbage = false
			}
		} else {
			if c == '{' {
				counter += actualScore
				actualScore++
			}
			if c == '}' {
				actualScore--
			}
		}
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data string) {
	counter := 0
	inGarbage := false
	escaping := false

	for _, c := range data {
		if escaping {
			escaping = false
			continue
		}
		if c == '!' {
			escaping = true
			continue
		}
		if c == '<' && !inGarbage {
			inGarbage = true
			continue
		}
		if inGarbage {
			if c == '>' {
				inGarbage = false
			} else {
				counter++
			}
		}
	}

	fmt.Printf("Part 2: %d\n", counter)
}
