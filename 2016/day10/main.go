package main

import (
	"fmt"
	"strings"

	"advent2016/util"
)

func main() {
	data := util.GetDataByRow("day10")
	// data := util.GetTestByRow("day10")

	list := parse(data)

	part1(list)
	part2(list)
}

type bot struct {
	idx       int
	values    []int
	lowOutput string // bot 1 -> b1; output 0 -> o0
	highOuput string
}

func parse(data []string) map[string]bot {
	bots := map[string]bot{}

	for _, s := range data {
		pieces := strings.Split(s, " ")
		if pieces[0] == "bot" {
			k := fmt.Sprintf("b%s", pieces[1])
			if _, ok := bots[k]; !ok {
				bots[k] = bot{idx: util.ToInt(pieces[1])}
			}
			b := bots[k]
			lO := pieces[5][:1] + pieces[6]
			hO := pieces[10][:1] + pieces[11]
			b.lowOutput = lO
			b.highOuput = hO
			bots[k] = b
		} else {
			k := fmt.Sprintf("b%s", pieces[5])
			if _, ok := bots[k]; !ok {
				bots[k] = bot{idx: util.ToInt(pieces[5])}
			}
			b := bots[k]
			b.values = append(b.values, util.ToInt(pieces[1]))
			bots[k] = b
		}
	}
	return bots
}

func part1(bots map[string]bot) {
	counter := 0

	twoValuedBots := []string{}

	for k, b := range bots {
		if len(b.values) == 2 {
			twoValuedBots = append(twoValuedBots, k)
		}
	}

	var bK string
	for len(twoValuedBots) > 0 {
		bK, twoValuedBots = twoValuedBots[0], twoValuedBots[1:]

		b := bots[bK]

		low := util.IntMin(b.values[0], b.values[1])
		high := util.IntMax(b.values[0], b.values[1])

		if low == 17 && high == 61 {
			counter = b.idx
			break
		}

		if strings.HasPrefix(b.lowOutput, "b") {
			bl := bots[b.lowOutput]
			bl.values = append(bots[b.lowOutput].values, low)
			bots[b.lowOutput] = bl
			if len(bl.values) == 2 {
				twoValuedBots = append(twoValuedBots, b.lowOutput)
			}
		}

		if strings.HasPrefix(b.highOuput, "b") {
			bh := bots[b.highOuput]
			bh.values = append(bots[b.highOuput].values, high)
			bots[b.highOuput] = bh
			if len(bh.values) == 2 {
				twoValuedBots = append(twoValuedBots, b.highOuput)
			}
		}

		b.values = []int{}

		bots[bK] = b
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(bots map[string]bot) {
	counter := 0

	twoValuedBots := []string{}
	outputs := map[string]int{}

	for k, b := range bots {
		if len(b.values) == 2 {
			twoValuedBots = append(twoValuedBots, k)
		}
	}

	var bK string
	for len(twoValuedBots) > 0 {
		bK, twoValuedBots = twoValuedBots[0], twoValuedBots[1:]

		b := bots[bK]

		low := util.IntMin(b.values[0], b.values[1])
		high := util.IntMax(b.values[0], b.values[1])

		if strings.HasPrefix(b.lowOutput, "b") {
			bl := bots[b.lowOutput]
			bl.values = append(bots[b.lowOutput].values, low)
			bots[b.lowOutput] = bl
			if len(bl.values) == 2 {
				twoValuedBots = append(twoValuedBots, b.lowOutput)
			}
		} else {
			// output
			outputs[b.lowOutput] = low
		}

		if strings.HasPrefix(b.highOuput, "b") {
			bh := bots[b.highOuput]
			bh.values = append(bots[b.highOuput].values, high)
			bots[b.highOuput] = bh
			if len(bh.values) == 2 {
				twoValuedBots = append(twoValuedBots, b.highOuput)
			}
		} else {
			// output
			outputs[b.highOuput] = high
		}

		b.values = []int{}

		bots[bK] = b
	}

	counter = outputs["o0"] * outputs["o1"] * outputs["o2"]
	fmt.Printf("Part 2: %d\n", counter)
}
