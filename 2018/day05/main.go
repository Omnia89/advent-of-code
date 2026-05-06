package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"

	"advent2018/util"
)

func main() {
	data := util.GetRawData("day05")
	// data := util.GetRawTest("day05")

	part1(data)
	part2(data)
}

func isUpper(r rune) bool {
	return unicode.IsUpper(r) && unicode.IsLetter(r)
}

func collapse(poly string) string {
	var sb strings.Builder

	lastLen := len(poly)

	for {
		for i := 0; i < len(poly); i++ {
			if i == len(poly)-1 {
				sb.WriteRune(rune(poly[i]))
				continue
			}
			toSkip := false
			if isUpper(rune(poly[i])) && !isUpper(rune(poly[i+1])) {
				if poly[i:i+1] == strings.ToUpper(poly[i+1:i+2]) {
					toSkip = true
				}
			} else if !isUpper(rune(poly[i])) && isUpper(rune(poly[i+1])) {
				if poly[i:i+1] == strings.ToLower(poly[i+1:i+2]) {
					toSkip = true
				}
			}
			if toSkip {
				// skip the next
				i++
				continue
			}

			sb.WriteRune(rune(poly[i]))
		}

		poly = sb.String()
		sb.Reset()
		if lastLen == len(poly) {
			break
		}
		lastLen = len(poly)
	}
	return poly
}

func part1(data string) {
	counter := 0

	counter = len(collapse(data))

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data string) {
	counter := math.MaxInt

	alpha := "abcdefghijklmnopqrstuvwxyz"

	for _, a := range alpha {
		c := string(a)

		t := strings.ReplaceAll(strings.ReplaceAll(data, c, ""), strings.ToUpper(c), "")

		collapsedLen := len(collapse(t))

		if collapsedLen < counter {
			counter = collapsedLen
		}
	}

	fmt.Printf("Part 2: %d\n", counter)
}
