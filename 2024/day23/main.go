package main

import (
	"fmt"
	"slices"
	"strings"

	"advent2024/util"
)

func main() {
	//data := util.GetTestByRow("day23")
	data := util.GetDataByRow("day23")

	joins := parse(data)

	part1(joins)
	part2(joins)
}

// alphabetical ordered
type join struct {
	first  string
	second string
}

func parse(data []string) []join {
	joins := []join{}

	for _, d := range data {
		nodes := strings.Split(d, "-")
		slices.Sort(nodes)
		joins = append(joins, join{nodes[0], nodes[1]})
	}
	return joins
}

func getLinkMap(joins []join) map[string][]string {
	links := map[string][]string{}

	for _, j := range joins {
		if _, ok := links[j.first]; !ok {
			links[j.first] = []string{}
		}
		if _, ok := links[j.second]; !ok {
			links[j.second] = []string{}
		}

		links[j.first] = append(links[j.first], j.second)
		links[j.second] = append(links[j.second], j.first)
	}
	return links
}

func commonArray(a, b []string) []string {
	common := []string{}
	for _, s := range a {
		if slices.Contains(b, s) {
			common = append(common, s)
		}
	}
	return common
}

func part1(joins []join) {
	counter := 0

	links := getLinkMap(joins)

	unique := map[string]bool{}

	addUnique := func(a, b, c string) {
		temp := []string{a, b, c}
		slices.Sort(temp)
		unique[fmt.Sprintf("%s-%s-%s", temp[0], temp[1], temp[2])] = true
	}

	for k, l := range links {
		if !strings.HasPrefix(k, "t") {
			continue
		}

		for _, middle := range l {
			for _, third := range links[middle] {
				if third == k {
					continue
				}
				if slices.Contains(l, third) {
					addUnique(k, middle, third)
				}
			}
		}
	}
	counter = len(unique)

	fmt.Printf("Part 1: %d\n", counter)
}

func netKey(a []string) string {
	slices.Sort(a)
	return strings.Join(a, ",")
}

func replaceInArray(a []string, old string, new string) []string {
	i := slices.Index(a, old)
	n := slices.Clone(a)
	if i < 0 {
		return n
	}
	n[i] = new
	return n
}

func part2(joins []join) {
	links := getLinkMap(joins)

	var bestNodeKey string

	for node, nodeLinks := range links {
		net := map[string]int{}
		common := slices.Clone(nodeLinks)
		for _, linkedNode := range nodeLinks {
			cc := replaceInArray(common, linkedNode, node)
			newCommon := replaceInArray(commonArray(cc, links[linkedNode]), node, linkedNode)
			net[netKey(newCommon)]++
			//common = newCommon
		}
		//fmt.Printf("  [%s] -> [%v]\n", node, net)

		for k, v := range net {
			kNum := strings.Count(k, ",") + 2
			bestK := 0
			if len(bestNodeKey) > 0 {
				bestK = strings.Count(bestNodeKey, ",") + 1
			}

			if kNum == v+1 && kNum > bestK {
				bestNodeKey = strings.Join(append(strings.Split(k, ","), node), ",")
			}
		}
	}

	ordered := strings.Split(bestNodeKey, ",")
	slices.Sort(ordered)
	bestNodeKey = strings.Join(ordered, ",")
	kNum := strings.Count(bestNodeKey, ",") + 1

	fmt.Printf("Part 2: [%d] %s\n", kNum, bestNodeKey)
}
