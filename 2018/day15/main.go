package main

import (
	"fmt"
	"slices"
	"strings"

	"advent2018/util"
)

func main() {
	data := util.GetDataByRow("day15")
	//data := util.GetTestByRow("day15")

	area, e, g := parse(data)

	e2 := make([]Creature, len(e))
	g2 := make([]Creature, len(g))

	copy(e2, e)
	copy(g2, g)

	part1(area, e, g)
	part2(area, e2, g2)
}

type Vec struct {
	x int
	y int
}

func (v Vec) getNear(maxX int, maxY int) []Vec {
	n := []Vec{}
	// Up
	if v.y > 0 {
		n = append(n, Vec{v.x, v.y - 1})
	}
	// Left
	if v.x > 0 {
		n = append(n, Vec{v.x - 1, v.y})
	}
	// Right
	if v.x < maxX {
		n = append(n, Vec{v.x + 1, v.y})
	}
	// Down
	if v.y < maxY {
		n = append(n, Vec{v.x, v.y + 1})
	}
	return n
}

type Point = Vec

type Creature struct {
	t   string
	pos Point
	hp  int
	atk int
}

func (c *Creature) attack(enemies []Creature, allCreatures []Creature) {
	enemiesPos := map[Point]Creature{}
	for _, e := range enemies {
		if e.hp > 0 {
			enemiesPos[e.pos] = e
		}
	}
	lowestHp := 300
	var e *Creature
	ns := c.pos.getNear(10000, 10000)
	for _, n := range ns {
		if cc, ok := enemiesPos[n]; ok && cc.hp < lowestHp {
			e = &cc
			lowestHp = cc.hp
		}
	}
	if e != nil {
		for i := range enemies {
			if enemies[i].pos == e.pos {
				enemies[i].hp -= c.atk
			}
		}
		for i := range allCreatures {
			if allCreatures[i].pos == e.pos {
				allCreatures[i].hp -= c.atk
			}
		}
	}
}

func (c *Creature) move(enemies []Creature, allies []Creature, area []string) {
	enemiesPos := map[Point]bool{}
	for _, e := range enemies {
		enemiesPos[e.pos] = true
	}
	maxX := len(area[0]) - 1
	maxY := len(area) - 1

	nAllies := make([]Creature, 0, len(allies)-1)
	for _, a := range allies {
		if a.pos != c.pos {
			nAllies = append(nAllies, a)
		}
	}
	newArea := redrawArea(area, nAllies)

	type P struct {
		Point
		origin Point
		dist   int
	}

	q := []P{}
	var p P
	visited := map[Point]bool{
		c.pos: true,
	}

	for _, n := range c.pos.getNear(maxX, maxY) {
		if enemiesPos[n] {
			return
		}
		if newArea[n.y][n.x] == '.' {
			q = append(q, P{n, n, 1})
			visited[n] = true
		}
	}

free:
	for len(q) > 0 {
		p, q = q[0], q[1:]

		for _, n := range p.getNear(maxX, maxY) {
			if visited[n] {
				continue
			}

			if enemiesPos[n] {
				c.pos = p.origin
				break free
			}
			if newArea[n.y][n.x] == '.' {
				q = append(q, P{n, p.origin, p.dist + 1})
				visited[n] = true
			}
		}
		slices.SortFunc(q, func(a P, b P) int {
			if a.dist != b.dist {
				return a.dist - b.dist
			}
			if a.y != b.y {
				return a.y - b.y
			}
			return a.x - b.x
		})
	}
}

func replaceString(s string, x int, n string) string {
	return fmt.Sprintf("%s%s%s", s[:x], n, s[x+1:])
}

func printArea(round int, area []string, cs []Creature) {
	temp := make([]string, len(area))
	copy(temp, area)

	for _, c := range cs {
		if c.hp <= 0 {
			continue
		}
		t := "G"
		if c.t == "elf" {
			t = "E"
		}
		temp[c.pos.y] = replaceString(temp[c.pos.y], c.pos.x, t)
	}
	fmt.Println(strings.Repeat("-", 25))
	fmt.Printf("Round: %d\n", round)
	for _, s := range temp {
		fmt.Println(s)
	}
	fmt.Println(strings.Repeat("-", 25))
}

func redrawArea(area []string, allies []Creature) []string {
	tempArea := make([]string, len(area))
	copy(tempArea, area)

	for _, a := range allies {
		if a.hp <= 0 {
			continue
		}
		t := tempArea[a.pos.y]
		t = replaceString(t, a.pos.x, "#")
		tempArea[a.pos.y] = t
	}
	return tempArea
}

func indexAt(s, sub string, i int) int {
	idx := strings.Index(s[i:], sub)
	if idx != -1 {
		return idx + i
	}
	return -1
}

func parse(data []string) (area []string, elfs []Creature, goblins []Creature) {
	for y, s := range data {
		i := indexAt(s, "E", 0)
		for i >= 0 {
			elfs = append(elfs, Creature{
				t:   "elf",
				pos: Point{i, y},
				hp:  200,
				atk: 3,
			})
			i = indexAt(s, "E", i+1)
		}
		i = indexAt(s, "G", 0)
		for i >= 0 {
			goblins = append(goblins, Creature{
				t:   "goblin",
				pos: Point{i, y},
				hp:  200,
				atk: 3,
			})
			i = indexAt(s, "G", i+1)
		}
		t := strings.ReplaceAll(s, "E", ".")
		t = strings.ReplaceAll(t, "G", ".")
		area = append(area, t)
	}
	return
}

func sortCreature(a, b Creature) int {
	if a.pos.y != b.pos.y {
		return a.pos.y - b.pos.y
	}
	return a.pos.x - b.pos.x
}

func removeCreature(cs []Creature, pos Point) []Creature {
	cc := []Creature{}
	for _, c := range cs {
		if c.pos != pos {
			cc = append(cc, c)
		}
	}
	return cc
}

func fight(area []string, elfs []Creature, goblins []Creature) (round int) {
	counter := 0

	findCreature := func(e Creature) func(c Creature) bool {
		return func(c Creature) bool {
			return c.pos == e.pos
		}
	}

	//printArea(0, area, append(elfs, goblins...))

	for {
		allCreatures := append(elfs, goblins...)
		slices.SortFunc(allCreatures, sortCreature)

		completeRound := true
		for i := range allCreatures {
			c := allCreatures[i]

			// Skip if dead
			if c.hp <= 0 {
				continue
			}
			aliveElfs := []Creature{}
			aliveGoblin := []Creature{}
			for _, cr := range elfs {
				if cr.hp > 0 {
					aliveElfs = append(aliveElfs, cr)
				}
			}
			for _, cr := range goblins {
				if cr.hp > 0 {
					aliveGoblin = append(aliveGoblin, cr)
				}
			}
			if len(aliveGoblin) == 0 || len(aliveElfs) == 0 {
				completeRound = false
				break
			}
			if c.t == "elf" {
				// move
				j := slices.IndexFunc(elfs, findCreature(c))
				c.move(aliveGoblin, elfs, area)
				elfs[j].pos = c.pos

				// attack
				c.attack(goblins, allCreatures)
			} else {
				// move
				j := slices.IndexFunc(goblins, findCreature(c))
				c.move(aliveElfs, goblins, area)
				goblins[j].pos = c.pos

				// attack
				c.attack(elfs, allCreatures)
			}

		}
		//printArea(counter+1, area, append(elfs, goblins...))
		if !completeRound {
			break
		}
		counter++

		aElfs := 0
		aGoblin := 0
		for _, cr := range allCreatures {
			if cr.hp <= 0 {
				continue
			}
			if cr.t == "elf" {
				aElfs++
			} else {
				aGoblin++
			}
		}
		if aElfs == 0 || aGoblin == 0 {
			break
		}
	}
	return counter
}

func part1(area []string, elfs []Creature, goblins []Creature) {
	counter := fight(area, elfs, goblins)

	hps := 0
	for _, c := range append(elfs, goblins...) {
		if c.hp <= 0 {
			continue
		}
		hps += c.hp
	}
	counter *= hps

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(area []string, originalElfs []Creature, originalGoblins []Creature) {
	counter := 0
	atk := 4

	for {
		goblins := make([]Creature, len(originalGoblins))
		copy(goblins, originalGoblins)
		elfs := make([]Creature, len(originalElfs))
		for i, elf := range originalElfs {
			elfs[i] = elf
			elfs[i].atk = atk
		}
		rounds := fight(area, elfs, goblins)
		aliveElfs := 0
		for _, e := range elfs {
			if e.hp > 0 {
				aliveElfs++
			}
		}
		if aliveElfs == len(originalElfs) {
			hps := 0
			for _, c := range append(elfs, goblins...) {
				if c.hp <= 0 {
					continue
				}
				hps += c.hp
			}
			counter = rounds * hps
			break
		}
		atk++
	}

	fmt.Printf("Part 2: %d\n", counter)
}
