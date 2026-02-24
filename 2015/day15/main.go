package main

import (
	"fmt"
	"strings"

	"advent2015/util"
)

func main() {
	data := util.GetDataByRow("day15")
	// data := util.GetTestByRow("day15")

	list := parse(data)

	part1(list)
	part2(list)
}

type ingridient struct {
	name   string
	values combination
}

type combination struct {
	cap int
	dur int
	fla int
	tex int
	cal int
}

func parse(data []string) map[string]ingridient {
	ings := map[string]ingridient{}

	for _, s := range data {
		pieces := strings.Split(s, " ")
		name := strings.Trim(pieces[0], ":")
		ings[name] = ingridient{
			name,
			combination{
				util.ToInt(strings.Trim(pieces[2], ",")),
				util.ToInt(strings.Trim(pieces[4], ",")),
				util.ToInt(strings.Trim(pieces[6], ",")),
				util.ToInt(strings.Trim(pieces[8], ",")),
				util.ToInt(strings.Trim(pieces[10], ",")),
			},
		}
	}
	return ings
}

func getCombinations(ingridients map[string]ingridient) []map[string]int {
	ings := make([]string, 0, len(ingridients))
	for _, i := range ingridients {
		ings = append(ings, i.name)
	}
	return combinationR(ings, 100)
}

func combinationR(ings []string, value int) []map[string]int {
	if len(ings) == 1 {
		return []map[string]int{
			{ings[0]: value},
		}
	}
	ret := []map[string]int{}
	for i := range value + 1 {
		subMap := combinationR(ings[1:], value-i)
		for _, s := range subMap {
			s[ings[0]] = i
		}
		ret = append(ret, subMap...)
	}
	return ret
}

func getScore(ingridients map[string]ingridient, spoons map[string]int) (score int, calories int) {
	scores := combination{}

	for ing, spoon := range spoons {
		ingr := ingridients[ing]
		scores.cap += ingr.values.cap * spoon
		scores.dur += ingr.values.dur * spoon
		scores.fla += ingr.values.fla * spoon
		scores.tex += ingr.values.tex * spoon
		scores.cal += ingr.values.cal * spoon
	}

	if scores.cap <= 0 || scores.dur <= 0 || scores.fla <= 0 || scores.tex <= 0 {
		return 0, 0
	}
	score = scores.cap * scores.dur * scores.fla * scores.tex
	calories = scores.cal

	return
}

func part1(ings map[string]ingridient) {
	counter := 0

	combs := getCombinations(ings)
	// fmt.Printf("cc [%v]\n", combs)

	maxScore := 0

	for _, c := range combs {
		score, _ := getScore(ings, c)
		if score > maxScore {
			maxScore = score
		}
	}

	counter = maxScore

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(ings map[string]ingridient) {
	counter := 0

	combs := getCombinations(ings)
	// fmt.Printf("cc [%v]\n", combs)

	maxScore := 0

	for _, c := range combs {
		score, cal := getScore(ings, c)
		if score > maxScore && cal == 500 {
			maxScore = score
		}
	}

	counter = maxScore
	fmt.Printf("Part 2: %d\n", counter)
}
