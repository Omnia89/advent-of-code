package main

import (
	"fmt"
	"slices"

	"advent2024/util"
)

func main() {
	data := util.GetDataByRow("day09")
	// data := util.GetTestByRow("day09")

	intData := util.StringToIntSlice(data[0], "")
	intData2 := slices.Clone(intData)

	part1(intData)
	part2(intData2)
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

type FreeSpace struct {
	Index  int
	Length int
}

func part2(data []int) {
	sum := 0

	freeSpaces := []FreeSpace{}
	compactDisk := []int{}

	index := 0
	for i, v := range data {
		isData := i%2 == 0
		val := i / 2
		if !isData {
			val = -1
			freeSpaces = append(freeSpaces, FreeSpace{index, v})
		}
		for range v {
			compactDisk = append(compactDisk, val)
		}
		index += v
	}

	lastI := len(compactDisk) - 1
	lastValue := compactDisk[lastI]
	for i := lastI - 1; i >= 0; i-- {
		if compactDisk[lastI] == compactDisk[i] {
			if compactDisk[i] == -1 {
				lastI = i
			}
			continue
		}

		// i[valA] lastI[-1]
		if compactDisk[i] != -1 && compactDisk[lastI] == -1 {
			lastI = i
			continue
		}

		// i[valA] lastI[valB]
		// i[-1] lastI[valA]

		if compactDisk[lastI] > lastValue {
			lastI = i
			continue
		}

		lastValue = compactDisk[lastI]

		space := lastI - i
		for k, f := range freeSpaces {
			if f.Index > i {
				continue
			}
			if f.Length >= space {
				freeSpaces[k].Length -= space
				freeSpaces[k].Index += space
				for v := range space {
					compactDisk[f.Index+v] = compactDisk[lastI]
					compactDisk[i+1+v] = -1
				}
				break
			}
		}
		lastI = i
	}
	for i, v := range compactDisk {
		if v != -1 {
			sum += i * v
		}
	}

	fmt.Printf("Part 2: %d\n", sum)
}
