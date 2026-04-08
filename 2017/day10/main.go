package main

import (
	"fmt"
	"strings"

	"advent2017/util"
)

func main() {
	data := util.GetDataByRow("day10")
	// data := util.GetTestByRow("day10")

	list := util.StringToIntSlice(data[0], ",")

	part1(list)
	part2(data[0])
}

func buildArray(n int) []int {
	r := make([]int, n)

	for i := range r {
		r[i] = i
	}

	return r
}

func swap(values *[]int, w int, start int, end int) {
	half := w / 2

	for i := range half {
		a := (start + i) % len(*values)
		b := end - i
		if b < 0 {
			b = len(*values) + b
		}
		(*values)[a], (*values)[b] = (*values)[b], (*values)[a]
	}
}

func part1(data []int) {
	counter := 0

	values := buildArray(256)
	// values := buildArray(5)
	position := 0

	for skip, n := range data {
		if n > len(values) {
			continue
		}
		if n > 1 {
			swap(&values, n, position, (position+n-1)%len(values))
		}
		position += skip + n
		position %= len(values)
	}

	counter = values[0] * values[1]

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data string) {
	list := []byte(data)
	list = append(list, []byte{17, 31, 73, 47, 23}...)

	values := buildArray(256)

	position := 0
	skip := 0
	for range 64 {
		for _, b := range list {
			n := int(b)
			if n > 1 {
				swap(&values, n, position, (position+n-1)%len(values))
			}
			position += skip + n
			position %= len(values)
			skip++
		}
	}

	res := []int{}
	t := 0
	for t < 16 {
		i := t * 16
		v := values[i]
		for j := i + 1; j < i+16; j++ {
			v ^= values[j]
		}
		res = append(res, v)
		t++
	}

	var sb strings.Builder
	for _, n := range res {
		fmt.Fprintf(&sb, "%02x", n)
	}

	fmt.Printf("Part 2: %s\n", sb.String())
}
