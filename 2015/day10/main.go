package main

import (
	"fmt"
	"strconv"
	"strings"

	"advent2015/util"
)

func main() {
	data := util.GetDataByRow("day10")
	// data := util.GetTestByRow("day10")

	part1(data[0])
	part2(data[0])
}

func lookAndTell(s string) string {
	res := []string{}
	arr := strings.Split(s, "")

	v := "+"
	c := 0
	for _, r := range arr {
		if v == "+" {
			v = r
			c++
			continue
		}

		if v == r {
			c++
		} else {
			// res += fmt.Sprintf("%d%s", c, string(v))
			res = append(res, strconv.Itoa(c), v)
			c = 1
			v = r
		}
	}
	// res += fmt.Sprintf("%d%s", c, string(v))
	res = append(res, strconv.Itoa(c), v)
	return strings.Join(res, "")
}

func part1(data string) {
	counter := 0

	s := data
	for range 40 {
		s = lookAndTell(s)
	}

	counter = len(s)

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data string) {
	counter := 0

	s := data
	for range 50 {
		s = lookAndTell(s)
	}

	counter = len(s)
	fmt.Printf("Part 2: %d\n", counter)
}
