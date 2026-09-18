package main

import (
	"advent2019/util"
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

var stdinReader = bufio.NewReader(os.Stdin)

func main() {
	data := util.GetRawData("day13")
	//data := util.GetRawTest("day13")

	list := util.StringToIntSlice(data, ",")
	list2 := slices.Clone(list)

	part1(list)
	part2(list2)
}

type Point struct {
	x int
	y int
}

func runArcade(data []int, display map[Point]int) {
	inChan := make(chan int, 3)
	outChan := make(chan int, 3)

	intcode := NewIntcodeChannel(data, inChan, outChan, false)

	go intcode.run(0)

	for {

		x, ok := <-outChan
		if !ok {
			break
		}

		y, ok := <-outChan
		if !ok {
			break
		}

		t, ok := <-outChan
		if !ok {
			break
		}

		display[Point{x, y}] = t
	}
}

func part1(data []int) {
	counter := 0

	display := map[Point]int{}

	runArcade(data, display)

	for _, v := range display {
		if v == 2 {
			counter++
		}
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func printDisplay(display map[Point]int, code *IntCode) {
	maxX, maxY := 43, 23

	grid := make([][]string, maxY)
	for i := range grid {
		grid[i] = make([]string, maxX)
		for j := range grid[i] {
			grid[i][j] = " "
		}
	}
	for p, v := range display {
		if v == 0 {
			continue
		}
		char := ""
		switch v {
		case 1:
			char = "#"
		case 2:
			char = "X"
		case 3:
			char = "="
		case 4:
			char = "o"
		default:
			char = "!"
		}
		grid[p.y][p.x] = char
	}
	var sb strings.Builder
	for _, row := range grid {
		sb.WriteString(strings.Join(row, ""))
		sb.WriteString("\n")
	}
	sb.WriteString(fmt.Sprintf("input: %v\n", code.inputArray))

	//fmt.Print("\033[26A\033[J")
	fmt.Println("-----------------------------------------------")
	fmt.Println(sb.String())
	//stdinReader.ReadString('\n')
	//time.Sleep(1000 * time.Millisecond)
}

func part2(data []int) {
	display := map[Point]int{}
	data[0] = 2

	intcode := NewIntcodeArray(data, nil, false)
	intcode.inputArray = []int{0}

	var endSetup bool

	var ballX, paddleX int
	score := 0

	for intcode.ip < len(intcode.program) {

		// input only when the program aspect an input, to avoid wrong input calculation
		if endSetup && len(intcode.inputArray) == 0 && intcode.program[intcode.ip]%100 == 3 {
			move := 0
			if ballX < paddleX {
				move = -1
			} else if ballX > paddleX {
				move = 1
			}
			intcode.inputArray = append(intcode.inputArray, move)
		}

		if intcode.next() {
			break
		}

		if len(intcode.outputArray) == 3 {
			x, y, t := intcode.outputArray[0], intcode.outputArray[1], intcode.outputArray[2]
			intcode.outputArray = []int{}

			if x == -1 && y == 0 {
				score = t
				continue
			}
			if !endSetup && x == 42 && y == 22 {
				endSetup = true
			}

			display[Point{x, y}] = t
			//if endSetup {
			//	printDisplay(display, intcode)
			//}

			if t == 3 {
				paddleX = x
			}
			if t == 4 {
				ballX = x
			}
		}

	}

	printDisplay(display, intcode)

	fmt.Printf("Part 2: %d\n", score)
}
