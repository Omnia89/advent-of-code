package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"

	"advent2016/util"
)

func main() {
	data := util.GetDataByRow("day05")
	// data := util.GetTestByRow("day05")

	part1(data)
	part2(data)
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func part1(data []string) {
	prefix := data[0]
	psw := ""

	i := 1
	for len(psw) < 8 {
		temp := fmt.Sprintf("%s%d", prefix, i)
		emdi5 := getMD5Hash(temp)
		if strings.HasPrefix(emdi5, "00000") {
			psw += emdi5[5:6]
		}
		i++
	}

	fmt.Printf("Part 1: %s\n", psw)
}

func replaceChar(s string, c string, pos int) string {
	if pos == len(s)-1 {
		return s[:pos] + c
	}
	return s[:pos] + c + s[pos+1:]
}

func part2(data []string) {
	prefix := data[0]
	psw := "________"

	i := 1
	for strings.Contains(psw, "_") {
		temp := fmt.Sprintf("%s%d", prefix, i)
		emdi5 := getMD5Hash(temp)
		if strings.HasPrefix(emdi5, "00000") {
			if strings.Contains("01234567", emdi5[5:6]) {
				position := util.ToInt(emdi5[5:6])
				if psw[position] == '_' {
					psw = replaceChar(psw, emdi5[6:7], position)
				}
			}
		}
		i++
	}
	fmt.Printf("Part 2: %s\n", psw)
}
