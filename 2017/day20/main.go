package main

import (
	"fmt"
	"math"
	"regexp"
	"slices"

	"advent2017/util"
)

func main() {
	data := util.GetDataByRow("day20")
	// data := util.GetTestByRow("day20")

	list := parse(data)
	list2 := slices.Clone(list)

	part1(list)
	part2(list2)
}

type Vect struct {
	x int
	y int
	z int
}

func (v *Vect) add(a Vect) {
	v.x += a.x
	v.y += a.y
	v.z += a.z
}

type Particle struct {
	pos Vect
	vel Vect
	acc Vect
}

func (p Particle) String() string {
	return fmt.Sprintf("p[%d, %d, %d] v[%d, %d, %d] a[%d, %d, %d]", p.pos.x, p.pos.y, p.pos.z, p.vel.x, p.vel.y, p.vel.z, p.acc.x, p.acc.y, p.acc.z)
}

func (p *Particle) tick() {
	p.vel.add(p.acc)
	p.pos.add(p.vel)
}

func (p Particle) distance() int {
	return util.IntAbs(p.pos.x) + util.IntAbs(p.pos.y) + util.IntAbs(p.pos.z)
}

func parse(data []string) []Particle {
	ps := []Particle{}

	extractRe := regexp.MustCompile(`<(-?\d+),(-?\d+),(-?\d+)>`)

	for _, s := range data {
		parts := extractRe.FindAllStringSubmatch(s, 3)
		ps = append(ps, Particle{
			pos: Vect{
				util.ToInt(parts[0][1]),
				util.ToInt(parts[0][2]),
				util.ToInt(parts[0][3]),
			},
			vel: Vect{
				util.ToInt(parts[1][1]),
				util.ToInt(parts[1][2]),
				util.ToInt(parts[1][3]),
			},
			acc: Vect{
				util.ToInt(parts[2][1]),
				util.ToInt(parts[2][2]),
				util.ToInt(parts[2][3]),
			},
		})

	}
	return ps
}

func part1(data []Particle) {
	counter := 0

	lastNearest := -1

	for counter < 1000 {
		nearest := -1
		nearestVal := math.MaxInt

		for i := range data {
			data[i].tick()
			d := data[i].distance()
			if d < nearestVal {
				nearestVal = d
				nearest = i
			}
		}

		if nearest == lastNearest {
			counter++
		} else {
			counter = 0
			lastNearest = nearest
		}
	}

	fmt.Printf("Part 1: %d\n", lastNearest)
}

func removeIdx(data []Particle, idxs []int) []Particle {
	ps := []Particle{}
	for i, p := range data {
		if !slices.Contains(idxs, i) {
			ps = append(ps, p)
		}
	}
	return ps
}

func part2(data []Particle) {
	counter := 0

	for counter < 1000 {
		positions := map[Vect][]int{}

		for i := range data {
			data[i].tick()
			positions[data[i].pos] = append(positions[data[i].pos], i)
		}

		idxs := []int{}
		for _, i := range positions {
			if len(i) > 1 {
				idxs = append(idxs, i...)
			}
		}

		if len(idxs) > 0 {
			data = removeIdx(data, idxs)
		}
		counter++
	}
	counter = len(data)
	fmt.Printf("Part 2: %d\n", counter)
}
