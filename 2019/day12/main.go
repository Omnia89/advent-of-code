package main

import (
	"fmt"
	"regexp"
	"slices"

	"advent2019/util"
)

func main() {
	data := util.GetDataByRow("day12")
	// data := util.GetTestByRow("day12")

	list := parse(data)
	list2 := slices.Clone(list)

	part1(list)

	part2(list2)
}

type Point struct {
	x int
	y int
	z int
}

func (p Point) toString() string {
	return fmt.Sprintf("<%d,%d,%d>", p.x, p.y, p.z)
}

type Moon struct {
	position Point
	velocity Point
}

func (m Moon) toString() string {
	return fmt.Sprintf("p:%s|v:%s", m.position.toString(), m.velocity.toString())
}

func parse(data []string) []Moon {
	ms := make([]Moon, 0, len(data))

	reg := regexp.MustCompile(`x=(\-?\d+),\sy=(\-?\d+),\sz=(\-?\d+)`)

	for _, s := range data {
		matches := reg.FindStringSubmatch(s)
		m := Moon{
			position: Point{
				x: util.ToInt(matches[1]),
				y: util.ToInt(matches[2]),
				z: util.ToInt(matches[3]),
			},
			velocity: Point{},
		}
		ms = append(ms, m)
	}

	return ms
}

func updateVelocity(moons []Moon) {
	for i := range len(moons) - 1 {
		for j := i + 1; j < len(moons); j++ {
			if moons[i].position.x < moons[j].position.x {
				moons[i].velocity.x += 1
				moons[j].velocity.x -= 1
			} else if moons[i].position.x > moons[j].position.x {
				moons[i].velocity.x -= 1
				moons[j].velocity.x += 1
			}

			if moons[i].position.y < moons[j].position.y {
				moons[i].velocity.y += 1
				moons[j].velocity.y -= 1
			} else if moons[i].position.y > moons[j].position.y {
				moons[i].velocity.y -= 1
				moons[j].velocity.y += 1
			}

			if moons[i].position.z < moons[j].position.z {
				moons[i].velocity.z += 1
				moons[j].velocity.z -= 1
			} else if moons[i].position.z > moons[j].position.z {
				moons[i].velocity.z -= 1
				moons[j].velocity.z += 1
			}
		}
	}
}

func step(moons []Moon) {
	for i := range moons {
		moons[i].position.x += moons[i].velocity.x
		moons[i].position.y += moons[i].velocity.y
		moons[i].position.z += moons[i].velocity.z
	}
}

func part1(moons []Moon) {
	counter := 0

	time := 0

	for time < 1000 {

		// update velocity
		updateVelocity(moons)
		// apply velocity
		step(moons)

		time++
	}

	for _, m := range moons {
		pot := util.IntAbs(m.position.x) + util.IntAbs(m.position.y) + util.IntAbs(m.position.z)
		kin := util.IntAbs(m.velocity.x) + util.IntAbs(m.velocity.y) + util.IntAbs(m.velocity.z)
		counter += pot * kin
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func axisState(moons []Moon, axis string) []int {
	s := make([]int, 0, len(moons)*2)
	for _, m := range moons {
		switch axis {
		case "x":
			s = append(s, m.position.x, m.velocity.x)
		case "y":
			s = append(s, m.position.y, m.velocity.y)
		case "z":
			s = append(s, m.position.z, m.velocity.z)
		}
	}
	return s
}

func axisCycle(moons []Moon, axis string) int {
	start := axisState(moons, axis)

	sim := slices.Clone(moons)
	steps := 0
	for {
		updateVelocity(sim)
		step(sim)
		steps++
		if slices.Equal(axisState(sim, axis), start) {
			return steps
		}
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func part2(moons []Moon) {
	cx := axisCycle(moons, "x")
	cy := axisCycle(moons, "y")
	cz := axisCycle(moons, "z")

	counter := lcm(lcm(cx, cy), cz)

	fmt.Printf("Part 2: %d\n", counter)
}
