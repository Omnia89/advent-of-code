package main

import (
	"fmt"
	"strings"

	"advent2016/util"
)

func main() {
	data := util.GetDataByRow("day16")
	// data := util.GetTestByRow("day16")

	part1(data[0])
	part2(data[0])
}

func dragon(s string) string {
	var sb strings.Builder

	for i := range s {
		o := len(s) - 1 - i
		if s[o] == '1' {
			sb.WriteString("0")
		} else {
			sb.WriteString("1")
		}
	}

	return s + "0" + sb.String()
}

func checksum(s string) string {
	first := true
	check := s
	var sb strings.Builder
	for first || len(check)%2 == 0 {
		first = false
		for i := 0; i < len(check); i += 2 {
			if check[i] == check[i+1] {
				sb.WriteString("1")
			} else {
				sb.WriteString("0")
			}
		}
		check = sb.String()
		sb.Reset()
	}
	return check
}

func part1(data string) {
	target := 272

	s := data
	for len(s) <= target {
		s = dragon(s)
	}

	check := checksum(s[:target])

	fmt.Printf("Part 1: %s\n", check)
}

func part2(data string) {
	target := 35651584

	s := data
	for len(s) <= target {
		s = dragon(s)
	}

	check := checksum(s[:target])

	fmt.Printf("Part 2: %s\n", check)
}
