package main

import (
	"fmt"
	"regexp"
	"strings"

	"advent2015/util"
)

func main() {
	data := util.GetDataByRow("day08")
	// data := util.GetTestByRow("day08")

	part1(data)
	part2(data)
}

func getRealString(s string) string {
	t := strings.Trim(s, `"`)

	hexReg := regexp.MustCompile(`\\x[0-9a-f]{2}`)
	slashReg := regexp.MustCompile(`\\\\`)
	quoteReg := regexp.MustCompile(`\\"`)

	t = hexReg.ReplaceAllString(t, "?")
	t = slashReg.ReplaceAllString(t, `\`)
	t = quoteReg.ReplaceAllString(t, `"`)

	return t
}

func encodeString(s string) string {
	t := s

	slashReg := regexp.MustCompile(`\\`)
	quoteReg := regexp.MustCompile(`"`)

	t = slashReg.ReplaceAllString(t, `\\`)
	t = quoteReg.ReplaceAllString(t, `\"`)

	return fmt.Sprintf(`"%s"`, t)
}

func part1(data []string) {
	counter := 0

	for _, s := range data {
		counter += len(s) - len(getRealString(s))
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data []string) {
	counter := 0

	for _, s := range data {
		counter += len(encodeString(s)) - len(s)
	}

	fmt.Printf("Part 2: %d\n", counter)
}
