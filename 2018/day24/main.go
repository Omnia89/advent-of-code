package main

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"advent2018/util"
)

func main() {
	data := util.GetDataByRow("day24")
	// data := util.GetTestByRow("day24")

	groups := parse(data)

	part1(groups)
	part2(groups)
}

type Group struct {
	groupType string
	id        int

	units      int
	hitPoint   int
	damage     int
	damageType string
	weaknesses []string
	immunities []string
	initiative int

	power int
}

func (g Group) String() string {
	return fmt.Sprintf("[%s:%d] [u:%d|d:%d|pow:%d]", g.groupType, g.id, g.units, g.damage, g.power)
}

func parse(data []string) (groups []Group) {
	group := "immune"
	id := 1

	r1 := regexp.MustCompile(`^(\d+)\sunits.*?(\d+)\shit\spoints.*?does\s(\d+)\s([a-z]+).*?(\d+)$`)
	r2 := regexp.MustCompile(`(weak|immune)\sto\s(.*?)(?:;|\))`)

	for i, s := range data {
		if i == 0 || s == "" {
			continue
		}
		if s == "Infection:" {
			group = "infection"
			id = 1
			continue
		}
		g := Group{
			groupType: group,
			id:        id,
		}
		g1 := r1.FindStringSubmatch(s)
		g2 := r2.FindAllStringSubmatch(s, -1)

		g.units = util.ToInt(g1[1])
		g.hitPoint = util.ToInt(g1[2])
		g.damage = util.ToInt(g1[3])
		g.damageType = g1[4]
		g.initiative = util.ToInt(g1[5])

		g.immunities = []string{}
		g.weaknesses = []string{}
		for _, gg := range g2 {
			list := strings.Split(gg[2], ", ")
			if gg[1] == "weak" {
				g.weaknesses = append(g.weaknesses, list...)
			} else {
				g.immunities = append(g.immunities, list...)
			}
		}
		g.power = g.units * g.damage
		groups = append(groups, g)
		id++
	}

	return
}

func compareGroupInitiative(a, b Group) int {
	return b.initiative - a.initiative
}

func selectTarget(attacker Group, groups []Group, selected map[int]bool) int {
	index := -1
	damage := map[int]int{}

	for i, g := range groups {
		if attacker.groupType == g.groupType || selected[i] {
			continue
		}
		if slices.Contains(g.immunities, attacker.damageType) {
			continue
		}
		d := attacker.power
		if slices.Contains(g.weaknesses, attacker.damageType) {
			d *= 2
		}
		damage[i] = d
	}

	greater := 0
	for i, d := range damage {
		if greater < d {
			index = i
			greater = d
		} else if greater == d {
			if groups[i].power == groups[index].power {
				if groups[i].initiative > groups[index].initiative {
					index = i
				}
			} else if groups[i].power > groups[index].power {
				index = i
			}
		}
	}

	return index
}

func cleanAndCheck(groups []Group) ([]Group, bool) {
	gs := make([]Group, 0, len(groups))
	gg := map[string]bool{}
	for _, g := range groups {
		if g.units > 0 {
			g.power = g.units * g.damage
			gs = append(gs, g)
			gg[g.groupType] = true
		}
	}
	return gs, len(gg) == 1
}

func getDamage(a Group, d Group) int {
	if slices.Contains(d.immunities, a.damageType) {
		return 0
	}

	damage := a.power
	if slices.Contains(d.weaknesses, a.damageType) {
		damage *= 2
	}

	return damage
}

func getSelectTargetOrder(groups []Group) []int {
	order := make([]int, 0, len(groups))
	for i := range groups {
		order = append(order, i)
	}

	slices.SortFunc(order, func(a, b int) int {
		if groups[a].power == groups[b].power {
			return groups[b].initiative - groups[a].initiative
		}

		return groups[b].power - groups[a].power
	})

	return order
}

func fight(groups []Group) (int, string) {
	var end bool

	slices.SortFunc(groups, compareGroupInitiative)

	for !end {
		selectOrder := getSelectTargetOrder(groups)
		targets := map[int]int{}    // [attacker]defender
		defenders := map[int]bool{} // [defender]used

		for _, g := range selectOrder {
			if groups[g].units <= 0 {
				continue
			}
			t := selectTarget(groups[g], groups, defenders)
			if t != -1 {
				targets[g] = t
				defenders[t] = true
			}
		}

		// deal damage
		draw := true
		for a, g := range groups {
			if g.units <= 0 {
				continue
			}
			if d, ok := targets[a]; ok {
				damage := getDamage(groups[a], groups[d])
				units := damage / groups[d].hitPoint
				groups[d].units -= units
				groups[d].power = groups[d].units * groups[d].damage
				if units > 0 {
					draw = false
				}
			}
		}
		// infection win on draw
		if draw {
			return 1, "infection"
		}

		groups, end = cleanAndCheck(groups)
	}

	survived := 0
	faction := ""
	for _, g := range groups {
		survived += g.units
		if faction == "" {
			faction = g.groupType
		}
	}

	return survived, faction
}

func part1(groups []Group) {
	counter := 0

	nG := slices.Clone(groups)

	counter, _ = fight(nG)

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(groups []Group) {
	counter := 0

	boostResult := map[int]int{} // map[boost]survived, negative if infection won

	lowerBound := 0
	upperBound := 1_000_000_000_000 // math.MaxInt
	boost := upperBound / 2

	for {
		nG := slices.Clone(groups)

		for i, g := range nG {
			if g.groupType == "immune" {
				nG[i].damage += boost
				nG[i].power = g.units * nG[i].damage
			}
		}

		//fmt.Printf("boost [%d]\t", boost)
		survived, faction := fight(nG)
		//fmt.Printf("faction [%s]\tsurvived [%d]\n", faction, survived)

		if faction == "infection" {
			boostResult[boost] = -survived

			if n, ok := boostResult[boost+1]; ok && n > 0 {
				counter = boostResult[boost+1]
				break
			}

			lowerBound = boost
		} else {
			boostResult[boost] = survived

			if n, ok := boostResult[boost-1]; ok && n <= 0 {
				counter = survived
				break
			}

			upperBound = boost
		}

		boost = (upperBound-lowerBound)/2 + lowerBound
	}

	fmt.Printf("Part 2: %d\n", counter)
}
