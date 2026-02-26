package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"strings"

	"advent2015/util"
)

func main() {
	data := util.GetDataByRow("day19")
	// data := util.GetTestByRow("day19")

	en := parse(data)

	part1(en)

	iEn := parseIndexed(data)

	part2(iEn)
}

func parse(data []string) enigma {
	last := ""
	dict := map[string][]string{}

	for _, s := range data {
		last = s
		a, b, ok := strings.Cut(s, " => ")
		if ok {
			if _, ok := dict[a]; !ok {
				dict[a] = []string{}
			}
			dict[a] = append(dict[a], b)
		}
	}

	return enigma{
		dict:  dict,
		start: last,
	}
}

func parseIndexed(data []string) listedEnigma {
	last := ""
	dict := []entry{}

	for _, s := range data {
		last = s
		a, b, ok := strings.Cut(s, " => ")
		if ok {
			dict = append(dict, entry{
				a,
				b,
			})
		}
	}

	return listedEnigma{
		dict:   dict,
		target: last,
	}
}

type entry struct {
	from string
	to   string
}

type listedEnigma struct {
	target string
	dict   []entry
}

type enigma struct {
	dict  map[string][]string
	start string
}

func part1(data enigma) {
	counter := 0

	molecules := map[string]bool{}

	for k, list := range data.dict {
		for _, v := range list {
			lastIndex := 0
			for {
				t := data.start[lastIndex:]
				i := strings.Index(t, k)
				if i == -1 {
					break
				}
				n := data.start[0:lastIndex+i] + v + data.start[lastIndex+i+len(k):]
				lastIndex += i + 1
				molecules[n] = true
			}
		}
	}

	counter = len(molecules)

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data listedEnigma) {
	counter := 0

	entries := data.dict

	slices.SortFunc(entries, func(a, b entry) int {
		return (len(b.to) + len(b.from)) - (len(a.to) + len(a.from))
	})

	target := data.target

externalLoop:
	for target != "e" {
		for _, e := range entries {
			idx := strings.Index(target, e.to)
			if idx >= 0 {
				target = target[:idx] + e.from + target[idx+len(e.to):]
				counter++
				continue externalLoop
				// break externalLoop
			}
		}
		// No sobstitution done and not "e", shuffle
		fmt.Println(" No SUB, shuffle")
		rand.Shuffle(len(entries), func(i, j int) { entries[i], entries[j] = entries[j], entries[i] })
		counter = 0
	}

	fmt.Printf("Part 2: %d\n", counter)
}
