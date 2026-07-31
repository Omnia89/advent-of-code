package main

import (
	"fmt"
	"math"
	"strings"

	"advent2019/util"
)

const (
	LAYER_W = 25
	LAYER_H = 6
)

// TEST
// const (
// 	LAYER_W = 3
// 	LAYER_H = 2
// )

func main() {
	data := util.GetRawData("day08")
	// data := util.GetRawTest("day08")

	list := util.StringToIntSlice(data, "")

	for i, n := range list {
		if n != util.ToInt(data[i:i+1]) {
			panic(fmt.Sprintf("TROVATO: [%d]\n", i))
		}
	}

	// fmt.Printf("%v\n", list)

	layerSize := (LAYER_H * LAYER_W)
	numLayers := len(data) / layerSize

	fmt.Printf(" len[%d] size[%d] layers[%d]\n", len(list), layerSize, numLayers)
	layers := make([][]int, numLayers)

	for i, n := range list {
		index := i / layerSize
		if index >= numLayers {
			break
		}
		layers[index] = append(layers[index], n)
	}

	part1(layers)
	part2(layers)
}

func count(arr []int) (zeros int, ones int, twos int) {
	for _, v := range arr {
		switch v {
		case 0:
			zeros++
		case 1:
			ones++
		case 2:
			twos++
		}
	}
	return
}

func part1(layers [][]int) {
	counter := 0

	zeros := math.MaxInt
	ones := 0
	twos := 0

	for _, l := range layers {
		z, o, t := count(l)
		if z < zeros {
			zeros = z
			ones = o
			twos = t
		}
	}
	counter = ones * twos

	fmt.Printf("Part 1: %d\n", counter)
}

func printLayer(l []int) {
	var sb strings.Builder

	for i, v := range l {
		c := "#"
		if v == 0 {
			c = " "
		}
		if i%LAYER_W == 0 {
			sb.WriteString("\n")
		}
		sb.WriteString(c)
	}
	sb.WriteString("\n")
	fmt.Printf("%s", sb.String())
}

func part2(layers [][]int) {
	counter := 0

	image := make([]int, LAYER_H*LAYER_W)
	for i := range image {
		image[i] = -1
	}

	for i := range image {
		for _, l := range layers {
			if l[i] == 2 {
				continue
			}
			image[i] = l[i]
			break
		}
	}
	printLayer(image)

	fmt.Printf("Part 2: %d\n", counter)
}
