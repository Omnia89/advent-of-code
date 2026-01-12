package main

import (
	"fmt"

	"advent2024/util"
)

func main() {
	data := util.GetDataByRow("day09")
	// data := util.GetTestByRow("day09")

	intData := util.StringToIntSlice(data[0], "")

	part1(intData)
	part2(intData)
}

func part1(data []int) {
	sum := 0

	lastDataIndex := 0
	for i := len(data) - 1; i >= 0; i-- {
		if i%2 == 1 {
			continue
		}
		if data[i] != 0 {
			lastDataIndex = i
			break
		}
	}

	fileBlocks := 0
	for i := 0; i < len(data); i += 2 {
		fileBlocks += data[i]
	}

	compactDisk := []int{}
externalFor:
	for i, v := range data {
		if i > lastDataIndex {
			break externalFor
		}
		isData := i%2 == 0
		if isData {
			val := i / 2
			for range v {
				compactDisk = append(compactDisk, val)
			}
		} else {
			for j := 0; j < v; j++ {
				if i > lastDataIndex {
					break externalFor
				}
				if data[lastDataIndex] == 0 {
					// go to next data
					j--
					lastDataIndex -= 2
					continue
				}
				val := lastDataIndex / 2
				compactDisk = append(compactDisk, val)
				data[lastDataIndex]--
			}
		}
	}
	for i, v := range compactDisk {
		sum += i * v
	}
	fmt.Printf("fileBlocks [%d] len [%d]\n", fileBlocks, len(compactDisk))

	fmt.Printf("Part 1: %d\n", sum)
}

func part2(data []int) {
	sum := 0

	fmt.Printf("Part 2: %d\n", sum)
}
