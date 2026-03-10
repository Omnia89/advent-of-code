package main

import (
	"fmt"
	"math"
	"slices"

	"advent2016/util"
)

func main() {
	data := util.GetDataByRow("day20")
	// data := util.GetTestByRow("day20")

	list := parse(data)

	part1(list)
	part2(list)
}

type Range struct {
	start int
	end   int
}

func (r Range) overlap(rr Range) bool {
	return r.start <= rr.start && rr.start <= r.end || r.start <= rr.end && rr.end <= r.end || r.start == rr.end+1 || r.end+1 == rr.start
}

func (r Range) merge(rr Range) Range {
	return Range{
		start: util.IntMin(r.start, rr.start),
		end:   util.IntMax(r.end, rr.end),
	}
}

func parse(data []string) []Range {
	r := []Range{}

	for _, s := range data {
		pieces := util.StringToIntSlice(s, "-")
		s := pieces[0]
		e := pieces[1]
		if e < s {
			s, e = e, s
		}
		r = append(r, Range{s, e})
	}

	return r
}

func mergeRanges(rs []Range) []Range {
	newR := []Range{}

	skipIds := []int{}
	for i := range len(rs) - 1 {
		if slices.Contains(skipIds, i) {
			continue
		}
		for j := i + 1; j < len(rs); j++ {
			if slices.Contains(skipIds, j) {
				continue
			}
			if rs[i].overlap(rs[j]) {
				rs[i] = rs[i].merge(rs[j])
				skipIds = append(skipIds, j)
			}
		}
		newR = append(newR, rs[i])
	}
	return newR
}

func part1(data []Range) {
	counter := 0

	slices.SortFunc(data, func(a, b Range) int {
		return a.start - b.start
	})

	newRanges := []Range{}
	oldLen := len(data)
	for {
		newRanges = mergeRanges(data)
		if oldLen == len(newRanges) {
			break
		}
		oldLen = len(newRanges)
	}

	lower := math.MaxInt

	for _, r := range newRanges {
		if r.start > 0 && r.start-1 < lower {
			lower = r.start - 1
		}
		if r.end < 4294967295 && r.end+1 < lower {
			lower = r.end + 1
		}
	}
	counter = lower

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data []Range) {
	counter := 0

	slices.SortFunc(data, func(a, b Range) int {
		return a.start - b.start
	})

	newRanges := []Range{}
	oldLen := len(data)
	for {
		newRanges = mergeRanges(data)
		if oldLen == len(newRanges) {
			break
		}
		oldLen = len(newRanges)
	}

	lastEnd := 0
	for i := range len(newRanges) - 1 {
		counter += newRanges[i+1].start - newRanges[i].end - 1
		lastEnd = newRanges[i+1].end
	}
	counter += 4294967295 - lastEnd

	fmt.Printf("Part 2: %d\n", counter)
}
