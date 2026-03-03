package main

import (
	"fmt"
	"maps"
	"regexp"
	"slices"
	"strings"

	"advent2016/util"
)

func main() {
	data := util.GetDataByRow("day04")
	// data := util.GetTestByRow("day04")

	list := parse(data)

	part1(list)
	part2(list)
}

type room struct {
	original string
	names    []string
	sector   int
	checksum string
}

func parse(data []string) []room {
	rooms := []room{}

	extractor := regexp.MustCompile(`(.*?)\-(\d+)\[(.*?)\]`)

	for _, s := range data {
		parts := extractor.FindStringSubmatch(s)
		rooms = append(rooms, room{
			s,
			strings.Split(parts[1], "-"),
			util.ToInt(parts[2]),
			parts[3],
		})
	}
	return rooms
}

func part1(data []room) {
	counter := 0

	for _, r := range data {
		letters := map[string]int{}

		for _, s := range strings.Join(r.names, "") {
			letters[string(s)]++
		}

		keys := slices.Collect(maps.Keys(letters))
		slices.SortFunc(keys, func(a, b string) int {
			vA := letters[a]
			vB := letters[b]
			if vA == vB {
				if a > b {
					return 1
				}
				return -1
			}
			return vB - vA
		})

		check := strings.Join(keys[:5], "")
		if check == r.checksum {
			counter += r.sector
		}
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func getChar(char string, r int) string {
	alpha := "abcdefghijklmnopqrstuvwxyz"
	index := (strings.Index(alpha, char) + r) % 26

	return alpha[index : index+1]
}

func part2(data []room) {
	counter := 0

	for _, r := range data {
		newString := ""
		for _, s := range strings.Join(r.names, " ") {
			t := string(s)
			if t == " " {
				newString = fmt.Sprintf("%s%s", newString, t)
				continue
			}
			newString = fmt.Sprintf("%s%s", newString, getChar(t, r.sector))
		}
		if strings.Contains(newString, "northpole") {
			counter = r.sector
			break
		}
	}

	fmt.Printf("Part 2: %d\n", counter)
}
