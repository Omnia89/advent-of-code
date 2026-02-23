package main

import (
	"fmt"
	"strings"

	"advent2015/util"
)

func main() {
	data := util.GetDataByRow("day11")
	//data := util.GetTestByRow("day11")

	part1(data[0])
	part2(data[0])
}

func replaceChar(s string, index int, char string) string {
	return s[:index] + char + s[index+1:]
}

func nextChar(c string) string {
	alpha := "abcdefghjkmnpqrstuvwxyz"
	index := strings.Index(alpha, c)
	if index < 0 {
		panic("Invalid char")
	}

	index = (index + 1) % len(alpha)
	return string(alpha[index])
}

func nextPsw(psw string) string {
	reverseIndex := 1
	nPsw := psw

	for {
		if reverseIndex > len(nPsw) {
			return fmt.Sprintf("a%s", nPsw)
		}
		n := nextChar(string(nPsw[len(nPsw)-reverseIndex]))
		nPsw = replaceChar(nPsw, len(nPsw)-reverseIndex, n)
		if n == "a" {
			reverseIndex++
		} else {
			break
		}
	}
	return nPsw
}

func checkPsw(psw string) bool {
	alpha := "abcdefghijklmnopqrstuvwxyz"

	doubleCounter := 0
	seq := false

	skipDoubleCheck := false

	for i := 1; i < len(psw); i++ {
		if skipDoubleCheck {
			skipDoubleCheck = false
		} else {
			if psw[i] == psw[i-1] {
				doubleCounter++
				skipDoubleCheck = true
			}
		}

		if i == 1 || seq {
			continue
		}

		idx := strings.Index(alpha, string(psw[i]))
		if idx < 2 {
			continue
		}
		if alpha[idx-1] == psw[i-1] && alpha[idx-2] == psw[i-2] {
			seq = true
		}
	}

	return seq && doubleCounter >= 2
}

func part1(psw string) {
	p := nextPsw(psw)
	for !checkPsw(p) {
		p = nextPsw(p)
	}

	fmt.Printf("Part 1: %s\n", p)
}

func part2(psw string) {
	p := nextPsw(psw)
	for !checkPsw(p) {
		p = nextPsw(p)
	}

	p = nextPsw(p)
	for !checkPsw(p) {
		p = nextPsw(p)
	}

	fmt.Printf("Part 2: %s\n", p)
}
