package main

import (
	"fmt"
	"math"
	"strings"

	"advent2015/util"
)

func main() {
	data := util.GetDataByRow("day21")
	// data := util.GetTestByRow("day21")

	shop := map[string][]equip{
		"weapons": {
			{8, 4, 0},
			{10, 5, 0},
			{25, 6, 0},
			{40, 7, 0},
			{74, 8, 0},
		},
		"armors": {
			{13, 0, 1},
			{31, 0, 2},
			{53, 0, 3},
			{75, 0, 4},
			{102, 0, 5},
		},
		"rings": {
			{25, 1, 0},
			{50, 2, 0},
			{100, 3, 0},
			{20, 0, 1},
			{40, 0, 2},
			{80, 0, 3},
		},
	}

	boss := parseBoss(data)

	part1(shop, boss)
	part2(shop, boss)
}

func parseBoss(data []string) character {
	_, hp, _ := strings.Cut(data[0], ": ")
	_, dmg, _ := strings.Cut(data[1], ": ")
	_, arm, _ := strings.Cut(data[2], ": ")

	return character{
		util.ToInt(hp),
		util.ToInt(dmg),
		util.ToInt(arm),
	}
}

type character struct {
	hitpoint int
	damage   int
	armor    int
}

type equip struct {
	cost   int
	damage int
	armor  int
}

func combinationList() [][]int {
	r := [][]int{}
	for w := range 5 {
		for a := range 6 {
			for r1 := range 7 {
				for r2 := range 7 {
					// permit both -1: no buy
					if r1 == r2 && r1 > 0 {
						continue
					}
					r = append(r, []int{w, a - 1, r1 - 1, r2 - 1})
				}
			}
		}
	}
	return r
}

func buildHero(equips []equip) (character, int) {
	c := character{hitpoint: 100}
	cost := 0

	for _, e := range equips {
		c.armor += e.armor
		c.damage += e.damage
		cost += e.cost
	}
	return c, cost
}

func simulateBattle(hero, boss character) (int, bool) {
	playerTurn := true
	step := 0
	for hero.hitpoint > 0 && boss.hitpoint > 0 {
		step++
		if playerTurn {
			boss.hitpoint -= util.IntMax(hero.damage-boss.armor, 1)
		} else {
			hero.hitpoint -= util.IntMax(boss.damage-hero.armor, 1)
		}
		playerTurn = !playerTurn
	}

	won := hero.hitpoint > 0

	return step, won
}

func part1(shop map[string][]equip, boss character) {
	counter := 0

	minCost := math.MaxInt
	for _, combo := range combinationList() {
		equips := []equip{}
		if v := combo[0]; v >= 0 {
			equips = append(equips, shop["weapons"][v])
		}
		if v := combo[1]; v >= 0 {
			equips = append(equips, shop["armors"][v])
		}
		if v := combo[2]; v >= 0 {
			equips = append(equips, shop["rings"][v])
		}
		if v := combo[3]; v >= 0 {
			equips = append(equips, shop["rings"][v])
		}
		hero, cost := buildHero(equips)

		_, won := simulateBattle(hero, boss)
		if won && cost < minCost {
			minCost = cost
		}
	}

	counter = minCost

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(shop map[string][]equip, boss character) {
	counter := 0

	maxCost := 0
	for _, combo := range combinationList() {

		equips := []equip{}
		if v := combo[0]; v >= 0 {
			equips = append(equips, shop["weapons"][v])
		}
		if v := combo[1]; v >= 0 {
			equips = append(equips, shop["armors"][v])
		}
		if v := combo[2]; v >= 0 {
			equips = append(equips, shop["rings"][v])
		}
		if v := combo[3]; v >= 0 {
			equips = append(equips, shop["rings"][v])
		}
		hero, cost := buildHero(equips)

		_, won := simulateBattle(hero, boss)
		if !won && cost > maxCost {
			maxCost = cost
		}
	}

	counter = maxCost
	fmt.Printf("Part 2: %d\n", counter)
}
