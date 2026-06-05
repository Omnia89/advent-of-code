package main

import (
	"fmt"
	"strings"

	"advent2018/util"
)

func main() {
	data := util.GetRawData("day14")
	//data := util.GetRawTest("day14")

	part1(data)
	part2(data)
}

func part1(in string) {
	table := []int{3, 7}

	num := util.ToInt(in)

	idx1 := 0
	idx2 := 1

	for len(table) < num+10 {
		nn := table[idx1] + table[idx2]
		// fmt.Printf("idx1[%d] idx2[%d] nn[%d] table: %v\n", idx1, idx2, nn, table)
		if nn > 9 {
			table = append(table, nn/10)
		}
		table = append(table, nn%10)

		idx1 = (idx1 + table[idx1] + 1) % len(table)
		idx2 = (idx2 + table[idx2] + 1) % len(table)
	}

	var sb strings.Builder
	for i := range 10 {
		fmt.Fprintf(&sb, "%d", table[num+i])
	}
	res := sb.String()

	fmt.Printf("Part 1: %s\n", res)
}

func part2(num string) {
	table := []int{3, 7}

	idx1 := 0
	idx2 := 1
	lastDigit := util.ToInt(num[len(num)-1:])
	delta := 0
	for {
		nn := table[idx1] + table[idx2]
		// fmt.Printf("idx1[%d] idx2[%d] nn[%d] table: %v\n", idx1, idx2, nn, table)
		if nn > 9 {
			table = append(table, nn/10)
		}
		table = append(table, nn%10)

		idx1 = (idx1 + table[idx1] + 1) % len(table)
		idx2 = (idx2 + table[idx2] + 1) % len(table)
		if nn%10 == lastDigit && len(table) >= len(num) {
			check := util.IntJoin(table[len(table)-len(num):], "")
			if check == num {
				break
			}
		}

		if nn > 9 && nn/10 == lastDigit && len(table) >= len(num)+1 {
			check := util.IntJoin(table[len(table)-len(num)-1:len(table)-1], "")
			if check == num {
				delta = 1
				break
			}
		}
	}
	counter := len(table) - len(num) - delta
	fmt.Printf("Part 2: %d\n", counter)
}
