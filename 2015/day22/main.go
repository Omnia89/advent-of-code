package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"advent2015/util"
)

func main() {
	data := util.GetDataByRow("day22")
	// data := util.GetTestByRow("day22")

	g := parse(data)

	// Run on real input
	part1(g)
	part2(g)
}

type game struct {
	turn          int
	playerTurn    bool
	heroHp        int
	bossHp        int
	bossDamage    int
	mana          int
	spentMana     int
	shieldConter  int
	poisonCounter int
	manaCounter   int
	// debugging: keep history of spells cast
	history []string
}

func (g game) toString() string {
	turn := "boss"
	if g.playerTurn {
		turn = "play"
	}
	return fmt.Sprintf("[%s][%03d] h[%03d] b[%02d] m[%03d - %03d] c[%d-%d-%d]", turn, g.turn, g.heroHp, g.bossHp, g.mana, g.spentMana, g.shieldConter, g.poisonCounter, g.manaCounter)
}

type spell struct {
	name          string
	cost          int
	damage        int
	heal          int
	shieldCounter int
	poisonCounter int
	manaCounter   int
}

func parse(data []string) game {
	_, bossHp, _ := strings.Cut(data[0], ": ")
	_, damage, _ := strings.Cut(data[1], ": ")

	return game{
		playerTurn: true,
		heroHp:     50,
		bossHp:     util.ToInt(bossHp),
		bossDamage: util.ToInt(damage),
		mana:       500,
	}
}

func part1(data game) {
	counter := 0

	spells := []spell{
		{"MagicMissile", 53, 4, 0, 0, 0, 0},
		{"Drain", 73, 2, 2, 0, 0, 0},
		{"Shield", 113, 0, 0, 6, 0, 0},
		{"Poison", 173, 0, 0, 0, 6, 0},
		{"Recharge", 229, 0, 0, 0, 0, 5},
	}

	queue := []game{data}
	var g game

externalLoop:
	for len(queue) > 0 {
		g, queue = queue[0], queue[1:]
		g.turn++

		// fmt.Println(g.toString())

		damage := g.bossDamage
		if g.shieldConter > 0 {
			g.shieldConter--
			damage -= 7
			damage = util.IntMax(damage, 1)
		}

		if g.manaCounter > 0 {
			g.manaCounter--
			g.mana += 101
		}

		if g.poisonCounter > 0 {
			g.poisonCounter--
			g.bossHp -= 3
		}

		if g.bossHp <= 0 {
			counter = g.spentMana
			fmt.Println("Winning sequence:", g.history, "spent:", g.spentMana)
			break externalLoop
		}

		if !g.playerTurn {
			g.heroHp -= damage

			// If hero dies, drop the branch
			if g.heroHp <= 0 {
				continue
			}
			g.playerTurn = true
			queue = append(queue, g)
		} else {
			for _, s := range spells {
				if g.mana-s.cost < 0 {
					continue
				}
				// skip if already present effects
				if s.manaCounter > 0 && g.manaCounter > 0 {
					continue
				}
				if s.poisonCounter > 0 && g.poisonCounter > 0 {
					continue
				}
				if s.shieldCounter > 0 && g.shieldConter > 0 {
					continue
				}

				ng := g

				ng.spentMana += s.cost
				ng.mana -= s.cost

				ng.bossHp -= s.damage
				ng.heroHp += s.heal

				ng.shieldConter += s.shieldCounter
				ng.poisonCounter += s.poisonCounter
				ng.manaCounter += s.manaCounter

				// append spell name to history
				ng.history = append(append([]string{}, g.history...), s.name)

				ng.playerTurn = false

				if ng.bossHp <= 0 {
					counter = ng.spentMana
					fmt.Println("Winning sequence:", ng.history, "spent:", ng.spentMana)
					break externalLoop
				}
				if ng.mana < 53 && ng.manaCounter == 0 {
					continue
				}

				queue = append(queue, ng)
			}
		}
		slices.SortFunc(queue, func(a, b game) int {
			if a.spentMana == b.spentMana {
				if a.bossHp == b.bossHp {
					return b.heroHp - a.heroHp
				}
				return a.bossHp - b.bossHp
			}
			return a.spentMana - b.spentMana
		})
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data game) {
	counter := 0

	spells := []spell{
		{"MagicMissile", 53, 4, 0, 0, 0, 0},
		{"Drain", 73, 2, 2, 0, 0, 0},
		{"Shield", 113, 0, 0, 6, 0, 0},
		{"Poison", 173, 0, 0, 0, 6, 0},
		{"Recharge", 229, 0, 0, 0, 0, 5},
	}

	queue := []game{data}
	var g game

	minMana := math.MaxInt
	bestGame := game{}

externalLoop:
	for len(queue) > 0 {
		g, queue = queue[0], queue[1:]
		g.turn++

		if g.playerTurn {
			g.heroHp--

			// If hero dies, drop the branch
			if g.heroHp <= 0 {
				continue
			}
		}

		damage := g.bossDamage
		if g.shieldConter > 0 {
			g.shieldConter--
			damage -= 7
			damage = util.IntMax(damage, 1)
		}

		if g.manaCounter > 0 {
			g.manaCounter--
			g.mana += 101
		}

		if g.poisonCounter > 0 {
			g.poisonCounter--
			g.bossHp -= 3
		}

		if g.bossHp <= 0 {
			if g.spentMana < minMana {
				minMana = g.spentMana
				bestGame = g
			}
			continue
		}

		if g.spentMana > minMana {
			continue
		}

		if !g.playerTurn {
			g.heroHp -= damage

			// If hero dies, drop the branch
			if g.heroHp <= 0 {
				continue
			}
			g.playerTurn = true
			queue = append(queue, g)
		} else {
			for _, s := range spells {
				if g.mana-s.cost < 0 {
					continue
				}
				// skip if already present effects
				if s.manaCounter > 0 && g.manaCounter > 0 {
					continue
				}
				if s.poisonCounter > 0 && g.poisonCounter > 0 {
					continue
				}
				if s.shieldCounter > 0 && g.shieldConter > 0 {
					continue
				}

				ng := g

				ng.spentMana += s.cost
				ng.mana -= s.cost

				ng.bossHp -= s.damage
				ng.heroHp += s.heal

				ng.shieldConter += s.shieldCounter
				ng.poisonCounter += s.poisonCounter
				ng.manaCounter += s.manaCounter

				ng.playerTurn = false

				// append spell name to history
				ng.history = append(append([]string{}, g.history...), s.name)

				if ng.bossHp <= 0 {
					if ng.spentMana < minMana {
						minMana = ng.spentMana
						bestGame = ng
					}
					continue externalLoop
				}

				queue = append(queue, ng)
			}
		}
	}
	counter = minMana
	fmt.Println("Winning sequence:", bestGame.history, "spent:", bestGame.spentMana)

	fmt.Printf("Part 2: %d\n", counter)
}
