package main

import (
	"fmt"

	"advent2016/util"
)

func main() {
	data := util.GetDataByRow("day03")
	// data := util.GetTestByRow("day03")

	part1(data)
	part2(data)
}

func part1(data []string) {
	counter := 0

triangleLoop:
	for _, s := range data {
		sides := util.StringToIntSlice(s, " ")

		for i := range 3 {
			j := (i + 1) % 3
			k := (i + 2) % 3

			if sides[i] >= sides[j]+sides[k] {
				continue triangleLoop
			}
		}
		counter++
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2Parse(data []string) [][3]int {
	a := [3]int{}
	b := [3]int{}
	c := [3]int{}
	values := [][3]int{}
	for r := 0; r < len(data); r += 3 {
		s1 := util.StringToIntSlice(data[r], " ")
		s2 := util.StringToIntSlice(data[r+1], " ")
		s3 := util.StringToIntSlice(data[r+2], " ")
		a = [3]int{s1[0], s2[0], s3[0]}
		b = [3]int{s1[1], s2[1], s3[1]}
		c = [3]int{s1[2], s2[2], s3[2]}

		values = append(values, a, b, c)
	}
	return values
}

func part2(data []string) {
	counter := 0

	values := part2Parse(data)

triangleLoop:
	for _, sides := range values {

		for i := range 3 {
			j := (i + 1) % 3
			k := (i + 2) % 3

			if sides[i] >= sides[j]+sides[k] {
				continue triangleLoop
			}
		}
		counter++
	}
	fmt.Printf("Part 2: %d\n", counter)
}
