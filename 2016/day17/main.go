package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math"
	"slices"
	"strings"

	"advent2016/util"
)

func main() {
	data := util.GetDataByRow("day17")
	// data := util.GetTestByRow("day17")

	part1(data[0])
	part2(data[0])
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// Up -  Down - Left - Right
func getDoorStatus(s string) []bool {
	hash := getMD5Hash(s)
	openDoors := "bcdef"

	doors := make([]bool, 4)

	for i := range 4 {
		doors[i] = strings.Contains(openDoors, string(hash[i]))
	}
	return doors
}

type Point struct {
	x int
	y int
}

type state struct {
	p         Point
	direction []string
}

func part1(data string) {
	target := Point{3, 3}

	minLen := math.MaxInt
	path := ""
	q := []state{{Point{0, 0}, nil}}
	var s state
	for len(q) > 0 {
		s, q = q[0], q[1:]

		if s.p == target {
			if len(s.direction) < minLen {
				minLen = len(s.direction)
				path = strings.Join(s.direction, "")
			}
			continue
		}
		if len(s.direction) == minLen {
			continue
		}

		doors := getDoorStatus(fmt.Sprintf("%s%s", data, strings.Join(s.direction, "")))

		// Up
		if doors[0] && s.p.y > 0 {
			dir := slices.Clone(s.direction)
			dir = append(dir, "U")
			nS := state{Point{s.p.x, s.p.y - 1}, dir}
			q = append(q, nS)
		}
		// Down
		if doors[1] && s.p.y < 3 {
			dir := slices.Clone(s.direction)
			dir = append(dir, "D")
			nS := state{Point{s.p.x, s.p.y + 1}, dir}
			q = append(q, nS)
		}
		// Left
		if doors[2] && s.p.x > 0 {
			dir := slices.Clone(s.direction)
			dir = append(dir, "L")
			nS := state{Point{s.p.x - 1, s.p.y}, dir}
			q = append(q, nS)
		}
		// Right
		if doors[3] && s.p.x < 3 {
			dir := slices.Clone(s.direction)
			dir = append(dir, "R")
			nS := state{Point{s.p.x + 1, s.p.y}, dir}
			q = append(q, nS)
		}
	}

	fmt.Printf("Part 1: %s\n", path)
}

func part2(data string) {
	counter := 0

	target := Point{3, 3}

	maxLen := 0
	q := []state{{Point{0, 0}, nil}}
	var s state
	for len(q) > 0 {
		s, q = q[0], q[1:]

		if s.p == target {
			if len(s.direction) > maxLen {
				maxLen = len(s.direction)
			}
			continue
		}

		doors := getDoorStatus(fmt.Sprintf("%s%s", data, strings.Join(s.direction, "")))

		// Up
		if doors[0] && s.p.y > 0 {
			dir := slices.Clone(s.direction)
			dir = append(dir, "U")
			nS := state{Point{s.p.x, s.p.y - 1}, dir}
			q = append(q, nS)
		}
		// Down
		if doors[1] && s.p.y < 3 {
			dir := slices.Clone(s.direction)
			dir = append(dir, "D")
			nS := state{Point{s.p.x, s.p.y + 1}, dir}
			q = append(q, nS)
		}
		// Left
		if doors[2] && s.p.x > 0 {
			dir := slices.Clone(s.direction)
			dir = append(dir, "L")
			nS := state{Point{s.p.x - 1, s.p.y}, dir}
			q = append(q, nS)
		}
		// Right
		if doors[3] && s.p.x < 3 {
			dir := slices.Clone(s.direction)
			dir = append(dir, "R")
			nS := state{Point{s.p.x + 1, s.p.y}, dir}
			q = append(q, nS)
		}
	}
	counter = maxLen
	fmt.Printf("Part 2: %d\n", counter)
}
