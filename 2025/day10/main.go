package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"advent2025/util"
)

func main() {
	data := util.GetDataByRow("day10") // 443 troppo basso
	// data := util.GetTestByRow("day10")

	problems := []Problem{}
	for _, s := range data {
		if string(s[0]) == "#" {
			continue
		}
		problems = append(problems, parseProblem(s))
	}

	debugger := false

	part1(problems, debugger)
	// part2(data)
}

type Problem struct {
	Lights  []bool
	Buttons [][]int
	Voltage []int
}

func (p Problem) toString() string {
	str := "["
	l := []string{}
	for _, ll := range p.Lights {
		s := "0"
		if ll {
			s = "1"
		}
		l = append(l, s)
	}
	str += strings.Join(l, " ") + "] - "
	bts := []string{}
	for _, b := range p.Buttons {
		t := "("
		for _, idx := range b {
			t += fmt.Sprintf("%d,", idx)
		}
		bts = append(bts, t[:len(t)-1]+")")
	}
	str += strings.Join(bts, " ")
	return str
}

func parseProblem(row string) Problem {
	parts := strings.Split(row, " ")
	lights := make([]bool, len(parts[0])-2)
	for i, c := range parts[0][1 : len(parts[0])-1] {
		lights[i] = string(c) == "#"
	}

	buttons := [][]int{}
	other := []int{}
	for _, p := range parts[1:] {
		arr := util.StringToIntSlice(p[1:len(p)-1], ",")
		if string(p[0]) == "{" {
			other = arr
		} else {
			buttons = append(buttons, arr)
		}
	}
	return Problem{
		Lights:  lights,
		Buttons: buttons,
		Voltage: other,
	}
}

func printBoolArray(arr []bool) string {
	s := []string{}
	for _, t := range arr {
		if t {
			s = append(s, "1")
		} else {
			s = append(s, "0")
		}
	}
	return strings.Join(s, " ")
}

func printMatrix(matrix [][]bool) {
	fmt.Printf("----------------\n")
	for _, r := range matrix {
		fmt.Printf("%s\n", printBoolArray(r))
	}
	fmt.Printf("----------------\n")
}

func buildMatrix(p Problem) [][]bool {
	matrix := [][]bool{}
	for r, l := range p.Lights {
		matrix = append(matrix, make([]bool, len(p.Buttons)+1))
		for c, b := range p.Buttons {
			matrix[r][c] = slices.Contains(b, r)
		}
		matrix[r][len(p.Buttons)] = l
	}
	return matrix
}

func xor(a, b bool) bool {
	return a != b
}

func xorRow(target []bool, additive []bool) []bool {
	ris := make([]bool, len(target))
	for i := range len(target) {
		ris[i] = xor(target[i], additive[i])
	}
	return ris
}

func getBinaryValue(number, index int) bool {
	pow := int(math.Pow(2, float64(index)))
	return number&pow == pow
}

func evaluateRow(row []bool, originalRow []bool, result bool) bool {
	firstFound := false
	value := true
	for i, r := range originalRow {
		if !firstFound {
			if r {
				firstFound = true
			}
		} else {
			if r {
				value = xor(value, row[i])
			}
		}
	}

	return value == result
}

func countTrue(row []bool) int {
	c := 0
	for _, r := range row {
		if r {
			c++
		}
	}
	return c
}

func part1(problems []Problem, debug bool) {
	counter := 0
	for _, prob := range problems {
		if debug {
			fmt.Printf("prob - %s\n", prob.toString())
		}

		indexButton := 0
		matrix := buildMatrix(prob)
		if debug {
			printMatrix(matrix)
		}
		row := 0
		column := 0
		pivots := map[int]int{} // [rowIdx]btnIdx
		btnPivotIndexes := []int{}
		searchingPivots := true
		for searchingPivots {
			// Ignore the "false cross" case, it should be impossible
			if column >= len(matrix[0])-1 || row >= len(matrix) {
				// the last column is the lights reference, no pivot to find
				searchingPivots = false
				continue
			}
			if !matrix[row][column] {
				// Find the pivot
				index := -1
				for r := row + 1; r < len(prob.Lights); r++ {
					if matrix[r][column] {
						index = r
						break
					}
				}
				if index != -1 {
					// swap rows
					matrix[row], matrix[index] = matrix[index], matrix[row]
				} else {
					// Skip the column
					column++
					indexButton++
					continue
				}
			}
			// xor-ing all the "true" rows
			for r := 0; r < len(matrix); r++ {
				if r == row || !matrix[r][column] {
					continue
				}
				matrix[r] = xorRow(matrix[r], matrix[row])
			}
			pivots[row] = indexButton
			btnPivotIndexes = append(btnPivotIndexes, indexButton)
			column++
			indexButton++
			row++
		}
		if debug {
			printMatrix(matrix)
		}

		minResult := math.MaxInt
		minButtons := []bool{}
		variableButtonsIndex := []int{}
		for i := range prob.Buttons {
			if !slices.Contains(btnPivotIndexes, i) {
				variableButtonsIndex = append(variableButtonsIndex, i)
			}
		}

		combinations := int(math.Pow(2, float64(len(variableButtonsIndex))))
		if debug {
			fmt.Printf("pivot %v\n", pivots)
			fmt.Printf("varBtn %v\n", variableButtonsIndex)
		}
		for i := range combinations {

			results := make([]bool, len(prob.Buttons))
			for j, v := range variableButtonsIndex {
				results[v] = getBinaryValue(i, j)
			}

			for ii, matrixRow := range matrix {
				compiledRow := slices.Clone(matrixRow[:len(matrixRow)-1])
				for j, v := range variableButtonsIndex {
					compiledRow[v] = getBinaryValue(i, j)
				}

				btnIndex, ok := pivots[ii]
				if ok {
					results[btnIndex] = evaluateRow(compiledRow, matrixRow[:len(matrixRow)-1], matrixRow[len(matrixRow)-1])
					if debug {
						fmt.Printf("     - btnIndex [%d] - row [%s] = [%t]\n", btnIndex, printBoolArray(compiledRow), results[btnIndex])
					}
				}
			}

			count := countTrue(results)

			if debug {
				fmt.Printf("  - riga temp [%d/%d] - %v - [%d]\n", i, combinations, printBoolArray(results), count)
			}
			if count < minResult {
				minResult = count
				minButtons = results
			}
		}
		if debug {
			fmt.Printf("res [%d] - buttons [%s]\n", minResult, printBoolArray(minButtons))
		}
		counter += minResult
	}

	fmt.Printf("Part 1: %d\n", counter)
}
