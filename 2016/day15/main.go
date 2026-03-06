package main

import (
	"fmt"
	"strings"

	"advent2016/util"
)

func main() {
	data := util.GetDataByRow("day15")
	// data := util.GetTestByRow("day15")

	list := parse(data)

	part1(list)
	part2(list)
}

type disc struct {
	positions int
	start     int
}

func parse(data []string) []disc {
	ds := []disc{}

	for _, s := range data {
		t := strings.ReplaceAll(s, ".", "")
		parts := strings.Split(t, " ")
		d := disc{
			positions: util.ToInt(parts[3]),
			start:     util.ToInt(parts[11]),
		}
		ds = append(ds, d)
	}
	return ds
}

func part1(data []disc) {
	counter := 0

	init := data[0].positions - data[0].start - 1

timeLoop:
	for time := init; ; time += data[0].positions {
		for n := 1; n < len(data); n++ {
			if (data[n].start+time+n+1)%data[n].positions != 0 {
				continue timeLoop
			}
		}
		counter = time
		break
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data []disc) {
	counter := 0

	data = append(data, disc{11, 0})

	init := data[0].positions - data[0].start - 1

timeLoop:
	for time := init; ; time += data[0].positions {
		for n := 1; n < len(data); n++ {
			if (data[n].start+time+n+1)%data[n].positions != 0 {
				continue timeLoop
			}
		}
		counter = time
		break
	}
	fmt.Printf("Part 2: %d\n", counter)
}
