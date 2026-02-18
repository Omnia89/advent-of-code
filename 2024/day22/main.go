package main

import (
	"fmt"

	"advent2024/util"
)

func main() {
	// data := util.GetRawTest("day22")
	data := util.GetRawData("day22")

	ints := util.StringToIntSlice(data, "\n")

	part1(ints)
	part2(ints)
}

func mixAndPrune(a int, secret int) int {
	pruneConst := 16777216
	temp := a ^ secret
	return temp % pruneConst
}

func getNext(secret int) int {
	next := secret
	// first step
	temp := secret * 64
	next = mixAndPrune(temp, next)

	// second step
	temp = next / 32
	next = mixAndPrune(temp, next)

	// third step
	temp = next * 2048
	next = mixAndPrune(temp, next)

	return next
}

func part1(data []int) {
	counter := 0

	for _, i := range data {
		// fmt.Printf("init [%d]\n", i)
		next := i
		for range 2000 {
			next = getNext(next)
		}
		// fmt.Printf("   [%d]\n", next)
		counter += next
	}

	fmt.Printf("Part 1: %d\n", counter)
}

type sequence struct {
	a int
	b int
	c int
	d int
}

func (s sequence) toString() string {
	return fmt.Sprintf("(%d,%d,%d,%d)", s.a, s.b, s.c, s.d)
}

func part2(data []int) {
	counter := 0

	limit := 2000

	deltaArrays := make([][]int, len(data))
	valueArrays := make([][]int, len(data))
	for i, v := range data {
		deltaArrays[i] = make([]int, limit)
		valueArrays[i] = make([]int, limit)
		next := v
		for j := range limit {
			temp := getNext(next)
			deltaArrays[i][j] = temp%10 - next%10
			valueArrays[i][j] = temp % 10
			next = temp
		}
	}

	allSequenceMap := map[sequence]bool{}
	values := []map[sequence]int{}

	for index, deltas := range deltaArrays {
		values = append(values, make(map[sequence]int))
		values[index] = map[sequence]int{}
		for i := range len(deltas) - 4 {
			seq := sequence{deltas[i], deltas[i+1], deltas[i+2], deltas[i+3]}
			if _, ok := values[index][seq]; ok {
				continue
			}
			allSequenceMap[seq] = true
			values[index][seq] = valueArrays[index][i+3]
		}
	}

	bestValue := 0
	// var bestSeq sequence

	for s := range allSequenceMap {
		sum := 0
		for _, v := range values {
			sum += v[s]
		}
		if sum > bestValue {
			bestValue = sum
			// bestSeq = s
		}
	}

	// fmt.Printf(" best sequence [%s] value[%d]\n", bestSeq.toString(), bestValue)

	counter = bestValue

	fmt.Printf("Part 2: %d\n", counter)
}
