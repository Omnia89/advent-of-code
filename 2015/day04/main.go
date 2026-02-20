package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"

	"advent2015/util"
)

func main() {
	data := util.GetDataByRow("day04")
	// data := util.GetTestByRow("day04")

	part1(data)
	part2(data)
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func part1(data []string) {
	counter := 0

	prefix := data[0]

	i := 1
	for {
		temp := fmt.Sprintf("%s%d", prefix, i)
		emdi5 := getMD5Hash(temp)
		if strings.HasPrefix(emdi5, "00000") {
			break
		}
		i++
	}
	counter = i

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data []string) {
	counter := 0

	prefix := data[0]

	i := 1
	for {
		temp := fmt.Sprintf("%s%d", prefix, i)
		emdi5 := getMD5Hash(temp)
		if strings.HasPrefix(emdi5, "000000") {
			break
		}
		i++
	}
	counter = i
	fmt.Printf("Part 2: %d\n", counter)
}
