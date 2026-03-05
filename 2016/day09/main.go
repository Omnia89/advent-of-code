package main

import (
	"fmt"
	"regexp"
	"strings"

	"advent2016/util"
)

func main() {
	data := util.GetDataByRow("day09")
	// data := util.GetTestByRow("day09")

	part1(data)
	part2(data)
}

func decompress(s string) string {
	var newStringBuilder strings.Builder
	for i := 0; i < len(s); i++ {
		char := string(s[i])
		if char != "(" {
			newStringBuilder.WriteString(char)
			continue
		}
		closingIndex := i + strings.Index(s[i:], ")")
		ps := util.StringToIntSlice(s[i+1:closingIndex], "x")

		toRepeat := s[closingIndex+1 : closingIndex+ps[0]+1]
		newStringBuilder.WriteString(strings.Repeat(toRepeat, ps[1]))

		i = closingIndex + ps[0]
	}
	newString := newStringBuilder.String()
	return newString
}

func part1(data []string) {
	counter := 0

	for _, s := range data {
		newString := decompress(s)
		counter += len(newString)
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func decompressR(s string) int {
	checkRe := regexp.MustCompile(`\(\d+x\d+\)`)

	idx := checkRe.FindStringIndex(s) // s[idx[0]:idx[1]]
	if idx == nil {
		return len(s)
	}

	l := len(s[:idx[0]])
	vals := util.StringToIntSlice(s[idx[0]+1:idx[1]-1], "x")

	return l + vals[1]*decompressR(s[idx[1]:idx[1]+vals[0]]) + decompressR(s[idx[1]+vals[0]:])
}

func part2(data []string) {
	counter := 0

	row := data[0]

	counter = decompressR(row)

	fmt.Printf("Part 2: %d\n", counter)
}
