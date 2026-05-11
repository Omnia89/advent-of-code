package main

import (
	"fmt"

	"advent2018/util"
)

func main() {
	data := util.GetDataByRow("day08")
	//data := util.GetTestByRow("day08")

	node := parse(data[0])

	part1(node)
	part2(node)
}

type Node struct {
	childNum int
	metaNum  int
	childs   []*Node
	meta     []int
}

func parse(data string) Node {
	entries := util.StringToIntSlice(data, " ")

	list := []*Node{}
	var n *Node

	initiating := true

	for i := 0; i < len(entries); {
		if initiating {
			nn := Node{
				childNum: entries[i],
				metaNum:  entries[i+1],
			}
			n = &nn
			i += 2
			initiating = false
			continue
		}
		if n.childNum > len(n.childs) {
			// stack last node and initialize new child
			initiating = true
			list = append(list, n)
		} else {
			// save meta, get parent and push child
			for range n.metaNum {
				n.meta = append(n.meta, entries[i])
				i++
			}
			if len(list) > 0 {
				parent := list[len(list)-1]
				list = list[:len(list)-1]

				parent.childs = append(parent.childs, n)
				n = parent
			}
		}
	}

	return *n
}

func recSum(n Node) int {
	s := 0

	for _, m := range n.meta {
		s += m
	}

	for _, c := range n.childs {
		s += recSum(*c)
	}
	return s
}

func part1(n Node) {
	counter := 0

	counter = recSum(n)

	fmt.Printf("Part 1: %d\n", counter)
}

func valueRec(n Node) int {
	if len(n.childs) == 0 {
		s := 0
		for _, m := range n.meta {
			s += m
		}
		return s
	}
	s := 0
	for _, m := range n.meta {
		if m > n.childNum {
			continue
		}
		s += valueRec(*n.childs[m-1])
	}
	return s
}

func part2(n Node) {
	counter := 0

	counter = valueRec(n)

	fmt.Printf("Part 2: %d\n", counter)
}
