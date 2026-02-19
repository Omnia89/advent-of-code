package main

import (
	"fmt"
	"maps"
	"math"
	"slices"
	"strconv"
	"strings"

	"advent2024/util"
)

func main() {
	// data := util.GetTestByRow("day24")
	data := util.GetDataByRow("day24")

	problem := parse(data)

	part1(problem)
	part2(problem)
}

type Gate struct {
	gateType string
	input1   string
	input2   string
	output   string
}

type Problem struct {
	values map[string]int
	gates  []Gate
}

func (p Problem) resetValues(x, y int) {
	for k := range p.values {
		if !strings.HasPrefix(k, "x") || !strings.HasPrefix(k, "y") {
			delete(p.values, k)
		} else {
			p.values[k] = 0
		}
	}
	xBin := util.StringToIntSlice(strconv.FormatInt(int64(x), 2), "")
	for i, v := range xBin {
		k := fmt.Sprintf("x%02d", i)
		p.values[k] = v
	}
	yBin := util.StringToIntSlice(strconv.FormatInt(int64(y), 2), "")
	for i, v := range yBin {
		k := fmt.Sprintf("y%02d", i)
		p.values[k] = v
	}
}

func parse(data []string) Problem {
	doneValues := false
	values := map[string]int{}
	gates := []Gate{}

	for _, s := range data {
		if s == "" {
			doneValues = true
			continue
		}
		if !doneValues {
			k, v, _ := strings.Cut(s, ":")
			values[k] = util.ToInt(v)
		} else {
			gate, output, _ := strings.Cut(s, " -> ")
			parts := strings.Split(gate, " ")
			gates = append(gates, Gate{
				strings.TrimSpace(parts[1]),
				strings.TrimSpace(parts[0]),
				strings.TrimSpace(parts[2]),
				output,
			})
		}
	}

	return Problem{
		values,
		gates,
	}
}

func exec(gateType string, input1 int, input2 int) int {
	switch gateType {
	case "AND":
		return input1 & input2
	case "OR":
		return input1 | input2
	case "XOR":
		return input1 ^ input2
	}
	return 0
}

func solve(problem Problem) int {
	queue := problem.gates

	var g Gate
	for len(queue) > 0 {
		g, queue = queue[0], queue[1:]

		input1, ok1 := problem.values[g.input1]
		if !ok1 {
			queue = append(queue, g)
			continue
		}
		input2, ok2 := problem.values[g.input2]
		if !ok2 {
			queue = append(queue, g)
			continue
		}

		problem.values[g.output] = exec(g.gateType, input1, input2)
	}

	zValues := []int{}

	i := 0
	for {
		zKey := fmt.Sprintf("z%02d", i)
		v, ok := problem.values[zKey]
		if !ok {
			break
		}
		zValues = append(zValues, v)
		i++
	}

	counter := 0
	for i, v := range zValues {
		counter += v * int(math.Pow(2, float64(i)))
	}
	return counter
}

func part1(problem Problem) {
	counter := 0

	counter = solve(problem)

	fmt.Printf("Part 1: %d\n", counter)
}

func checkLastLayer(gate Gate, significantZ string) string {
	// All z come from XOR, except last (more significant), that's an OR
	if gate.output == significantZ {
		if gate.gateType != "OR" {
			return gate.output
		}
	} else {
		if gate.gateType != "XOR" {
			return gate.output
		}
	}

	// Special case: z00 always comes from x00 XOR y00
	if gate.output == "z00" {
		if !slices.Contains([]string{"x00", "y00"}, gate.input1) || !slices.Contains([]string{"x00", "y00"}, gate.input2) {
			return gate.output
		}
		return ""
	}

	// No other gate comes directly from x or y
	if strings.HasPrefix(gate.input1, "x") || strings.HasPrefix(gate.input1, "y") || strings.HasPrefix(gate.input2, "x") || strings.HasPrefix(gate.input2, "y") {
		return gate.output
	}
	return ""
}

func checkLayer(gate Gate, gateInput1 Gate, gateInput2 Gate) string {
	switch gate.gateType {
	case "OR":
		// OR gates always have ANDs as inputs
		if gateInput1.gateType != "AND" {
			return gateInput1.output
		}
		if gateInput2.gateType != "AND" {
			return gateInput2.output
		}
	case "AND":
		// No two ANDs in a row (unless its x00)
		if gateInput1.gateType == "AND" && !slices.Contains([]string{gateInput1.input1, gateInput1.input2}, "x00") {
			return gateInput1.output
		}
		if gateInput2.gateType == "AND" && !slices.Contains([]string{gateInput2.input1, gateInput2.input2}, "x00") {
			return gateInput2.output
		}
	case "XOR":
		// XORs that have OR and XOR as inputs, must output in z
		if gateInput1.gateType == "XOR" && gateInput2.gateType == "OR" || gateInput2.gateType == "XOR" && gateInput1.gateType == "OR" {
			if !strings.HasPrefix(gate.output, "z") {
				return gate.output
			}
		}
	}
	return ""
}

func part2(problem Problem) {
	gatesByOutput := map[string]Gate{}

	queue := []string{}

	for _, g := range problem.gates {
		gatesByOutput[g.output] = g
		if strings.HasPrefix(g.output, "z") {
			queue = append(queue, g.output)
		}
	}
	significantZ := fmt.Sprintf("z%02d", len(queue)-1)

	wrongGates := map[string]bool{}

	var out string
	for len(queue) > 0 {
		out, queue = queue[0], queue[1:]
		isZ := strings.HasPrefix(out, "z")
		if gate, ok := gatesByOutput[out]; ok {
			if isZ {
				if v := checkLastLayer(gate, significantZ); v != "" {
					wrongGates[v] = true
				}
			} else {
				in1Gate := gatesByOutput[gate.input1]
				in2Gate := gatesByOutput[gate.input2]
				if v := checkLayer(gate, in1Gate, in2Gate); v != "" {
					wrongGates[v] = true
				}
			}
			queue = append(queue, gate.input1, gate.input2)
		}
	}

	gates := slices.Collect(maps.Keys(wrongGates))
	slices.Sort(gates)

	solution := strings.Join(gates, ",")

	fmt.Printf("Part 2: %s\n", solution)
}
