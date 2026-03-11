package main

import (
	"fmt"
	"slices"
	"strings"

	"advent2016/util"
)

func main() {
	data := util.GetDataByRow("day21")
	// data := util.GetTestByRow("day21")

	list := parse(data)

	part1(list)
	part2(list)
}

type operation struct {
	op    string
	val1  int
	val2  int
	char1 string
	char2 string
}

func parse(data []string) []operation {
	ops := []operation{}
	for _, s := range data {
		parts := strings.Split(s, " ")
		if parts[0] == "swap" {
			if parts[1] == "position" {
				ops = append(ops, operation{
					op:   "swap-position",
					val1: util.ToInt(parts[2]),
					val2: util.ToInt(parts[5]),
				})
			} else {
				ops = append(ops, operation{
					op:    "swap-letter",
					char1: parts[2],
					char2: parts[5],
				})
			}
		} else if parts[0] == "rotate" {
			if parts[1] == "based" {
				ops = append(ops, operation{
					op:    "rotate-letter",
					char1: parts[6],
				})
			} else {
				val := util.ToInt(parts[2])
				if parts[1] == "left" {
					val *= -1
				}
				ops = append(ops, operation{
					op:   "rotate-position",
					val1: val,
				})
			}
		} else {
			ops = append(ops, operation{
				op:   parts[0],
				val1: util.ToInt(parts[2]),
				val2: util.ToInt(parts[len(parts)-1]),
			})
		}
	}
	return ops
}

func swapPosition(v string, o operation) string {
	if o.val1 == o.val2 {
		return v
	}
	l := util.IntMin(o.val1, o.val2)
	h := util.IntMax(o.val1, o.val2)
	return v[:l] + v[h:h+1] + v[l+1:h] + v[l:l+1] + v[h+1:]
}

func swapLetter(v string, o operation) string {
	if o.char1 == o.char2 {
		return v
	}
	var sb strings.Builder

	for _, c := range v {
		ch := string(c)
		if ch == o.char1 {
			sb.WriteString(o.char2)
		} else if ch == o.char2 {
			sb.WriteString(o.char1)
		} else {
			sb.WriteString(ch)
		}
	}
	return sb.String()
}

func rotatePosition(v string, o operation) string {
	if o.val1 == 0 {
		return v
	}
	rot := o.val1 % len(v)
	if rot < 0 {
		rot = len(v) + rot
	}

	return v[len(v)-rot:] + v[:len(v)-rot]
}

func rotateLetter(v string, o operation) string {
	idx := strings.Index(v, o.char1)

	if idx >= 4 {
		idx++
	}
	idx += 1
	idx %= len(v)

	return v[len(v)-idx:] + v[:len(v)-idx]
}

func reversePosition(v string, o operation) string {
	if o.val1 == o.val2 {
		return v
	}
	l := util.IntMin(o.val1, o.val2)
	h := util.IntMax(o.val1, o.val2)

	t := strings.Split(v[l:h+1], "")
	slices.Reverse(t)
	reversed := strings.Join(t, "")
	return v[:l] + reversed + v[h+1:]
}

func movePosition(v string, o operation) string {
	if o.val1 == o.val2 {
		return v
	}
	var sb strings.Builder
	l := util.IntMin(o.val1, o.val2)
	h := util.IntMax(o.val1, o.val2)

	sb.WriteString(v[:l])

	if o.val1 < o.val2 {
		// shift-right
		sb.WriteString(v[o.val1+1 : o.val2+1])
		sb.WriteString(v[o.val1 : o.val1+1])
	} else {
		// shift left
		sb.WriteString(v[o.val1 : o.val1+1])
		sb.WriteString(v[o.val2:o.val1])
	}
	sb.WriteString(v[h+1:])

	return sb.String()
}

type opFunc func(string, operation) string

func part1(data []operation) {
	psw := "abcdefgh"
	// psw := "abcde" //  test

	ops := map[string]opFunc{
		"swap-position":   swapPosition,
		"swap-letter":     swapLetter,
		"rotate-position": rotatePosition,
		"rotate-letter":   rotateLetter,
		"reverse":         reversePosition,
		"move":            movePosition,
	}

	for _, o := range data {
		psw = ops[o.op](psw, o)
	}

	fmt.Printf("Part 1: %s\n", psw)
}

func reverseOp(v string, o operation) operation {
	switch o.op {
	case "swap-position", "swap-letter", "reverse":
		return o
	case "rotate-position":
		nO := o
		nO.val1 = -nO.val1
		return nO
	case "rotate-letter":
		nO := o
		nO.op = "rotate-position"
		// BF: try all possibilities, can't invert directly mathematically
		for i := 0; i < len(v); i++ {
			candidate := v[len(v)-i:] + v[:len(v)-i]
			if rotateLetter(candidate, o) == v {
				nO.val1 = i
				return nO
			}
		}
	case "move":
		nO := o
		nO.val1, nO.val2 = nO.val2, nO.val1
		return nO
	}
	return o
}

func part2(data []operation) {
	psw := "fbgdceah"

	ops := map[string]opFunc{
		"swap-position":   swapPosition,
		"swap-letter":     swapLetter,
		"rotate-position": rotatePosition,
		"rotate-letter":   rotateLetter,
		"reverse":         reversePosition,
		"move":            movePosition,
	}

	for i := len(data) - 1; i >= 0; i-- {
		rev := reverseOp(psw, data[i])
		psw = ops[rev.op](psw, rev)
	}
	fmt.Printf("Part 2: %s\n", psw)
}
