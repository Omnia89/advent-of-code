package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"maps"
	"slices"

	"advent2016/util"
)

func main() {
	data := util.GetDataByRow("day14")
	// data := util.GetTestByRow("day14")

	part1(data[0])
	part2(data[0])
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func findTriplets(s string) string {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+1] && s[i] == s[i+2] {
			return s[i : i+1]
		}
	}

	return ""
}

func findQuintuplets(s string) []string {
	t := map[string]bool{}

	for i := 0; i < len(s)-4; i++ {
		if s[i] == s[i+1] && s[i] == s[i+2] && s[i] == s[i+3] && s[i] == s[i+4] {
			t[s[i:i+1]] = true
			// skip a little
			i += 4
		}
	}

	r := slices.Collect(maps.Keys(t))
	return r
}

func solve(salt string, encodings int) int {
	triplets := map[string][]int{}

	psws := []int{}

	n := 0
	lastIdx := -1
	for len(psws) <= 64 || n <= lastIdx+1000 {
		encoded := fmt.Sprintf("%s%d", salt, n)
		for range encodings {
			encoded = getMD5Hash(encoded)
		}

		trips := findTriplets(encoded)
		quints := findQuintuplets(encoded)

		for _, q := range quints {
			for _, idx := range triplets[q] {
				delta := n - idx
				if delta > 1000 {
					continue
				}
				if lastIdx < idx {
					lastIdx = idx
				}
				psws = append(psws, idx)
			}
			triplets[q] = []int{}
		}

		// after, so we do not catch the same row
		triplets[trips] = append(triplets[trips], n)
		n++
	}

	slices.Sort(psws)
	return psws[63]
}

func part1(salt string) {
	counter := 0

	counter = solve(salt, 1)

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(salt string) {
	counter := 0

	counter = solve(salt, 2017)

	fmt.Printf("Part 2: %d\n", counter)
}
