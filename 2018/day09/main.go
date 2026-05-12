package main

import (
	"container/ring"
	"fmt"
	"strings"

	"advent2018/util"
)

func main() {
	data := util.GetDataByRow("day09")
	//data := util.GetTestByRow("day09")

	p, m := parse(data[0])

	part1(p, m)
	part2(p, m)
}

func parse(data string) (players int, marbles int) {
	parts := strings.Split(data, " ")
	return util.ToInt(parts[0]), util.ToInt(parts[6])
}

func marbliant(players, marbles int) int {
	scores := make([]int, players)

	cur := ring.New(1)
	cur.Value = 0

	player := 0
	for t := 1; t <= marbles; t++ {
		if t%23 == 0 {
			scores[player] += t
			cur = cur.Move(-7)
			scores[player] += cur.Value.(int)
			next := cur.Next()
			cur.Prev().Unlink(1)
			cur = next
		} else {
			cur = cur.Next()
			n := ring.New(1)
			n.Value = t
			cur.Link(n)
			cur = n
		}
		player = (player + 1) % players
	}

	max := 0
	for _, s := range scores {
		if s > max {
			max = s
		}
	}
	return max
}

func part1(players int, marbles int) {
	counter := marbliant(players, marbles)

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(players int, marbles int) {
	counter := marbliant(players, marbles*100)

	fmt.Printf("Part 2: %d\n", counter)
}
