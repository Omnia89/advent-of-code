package main

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"advent2018/util"
)

func main() {
	data := util.GetDataByRow("day04")
	// data := util.GetTestByRow("day04")

	list := parse(data)

	part1(list)
	part2(list)
}

type Guard struct {
	id           int
	date         time.Time
	minutesSleep []bool
}

func parse(data []string) []Guard {
	gs := []Guard{}

	slices.Sort(data)

	var guard *Guard = nil
	lastIndex := 0
	asleep := false

	for _, s := range data {
		parts := strings.Split(s, " ")
		_, minuteStr, _ := strings.Cut(parts[1], ":")
		minute := util.ToInt(minuteStr[:2])
		if strings.Contains(s, "begins") {
			if guard != nil {
				if asleep {
					for i := lastIndex; i < 60; i++ {
						guard.minutesSleep[i] = true
					}
				}

				gs = append(gs, *guard)
			}

			start, _ := time.Parse("2006-01-02", parts[0][1:])
			start = start.Add(time.Hour)
			start = time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, time.UTC)

			guard = &Guard{
				id:           util.ToInt(parts[3][1:]),
				date:         start,
				minutesSleep: make([]bool, 60),
			}
			lastIndex = 0
		} else if strings.Contains(s, "sleep") {
			lastIndex = minute
			asleep = true
		} else if strings.Contains(s, "wakes") {
			for i := lastIndex; i < minute; i++ {
				guard.minutesSleep[i] = true
			}
			asleep = false
			lastIndex = minute
		}

	}
	if asleep {
		for i := lastIndex; i < 60; i++ {
			guard.minutesSleep[i] = true
		}
	}

	gs = append(gs, *guard)

	return gs
}

func part1(data []Guard) {
	counter := 0

	// id - total minutes sleeping
	totSleep := map[int]int{}

	// id - 60 minutes counter
	minutesSleeping := map[int][]int{}

	for _, g := range data {
		if _, ok := minutesSleeping[g.id]; !ok {
			minutesSleeping[g.id] = make([]int, 60)
		}
		for i, m := range g.minutesSleep {
			if m {
				totSleep[g.id] += 1

				minutesSleeping[g.id][i] += 1
			}
		}
	}

	mostAsleep := 0
	mostAsleepC := 0
	for id, v := range totSleep {
		if v > mostAsleepC {
			mostAsleepC = v
			mostAsleep = id
		}
	}

	mostAsleepMinute := 0
	mostAsleepMinuteC := 0

	for i, v := range minutesSleeping[mostAsleep] {
		if v > mostAsleepMinuteC {
			mostAsleepMinuteC = v
			mostAsleepMinute = i
		}
	}

	counter = mostAsleep * mostAsleepMinute

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data []Guard) {
	counter := 0

	// id - 60 minutes counter
	minutesSleeping := map[int][]int{}

	for _, g := range data {
		if _, ok := minutesSleeping[g.id]; !ok {
			minutesSleeping[g.id] = make([]int, 60)
		}
		for i, m := range g.minutesSleep {
			if m {
				minutesSleeping[g.id][i] += 1
			}
		}
	}

	mostAsleep := 0
	mostAsleepMinute := 0
	mostAsleepMinuteC := 0

	for i, r := range minutesSleeping {
		for m, v := range r {
			if v > mostAsleepMinuteC {
				mostAsleepMinuteC = v
				mostAsleepMinute = m
				mostAsleep = i
			}
		}
	}

	counter = mostAsleep * mostAsleepMinute
	fmt.Printf("Part 2: %d\n", counter)
}
