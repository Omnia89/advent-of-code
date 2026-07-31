package main

import (
	"fmt"
	"slices"
	"sync"

	"advent2019/util"
)

func main() {
	data := util.GetRawData("day07")
	// data := util.GetRawTest("day07")

	list := util.StringToIntSlice(data, ",")

	part1(list)
	part2(list)
}

func permutations(n int) [][]int {
	base := make([]int, n+1)
	for i := range base {
		base[i] = i
	}

	var result [][]int
	permutation_r(base, 0, &result)
	return result
}

func permutation_r(arr []int, k int, res *[][]int) {
	if k == len(arr) {
		cc := make([]int, len(arr))
		copy(cc, arr)
		*res = append(*res, cc)
		return
	}

	for i := k; i < len(arr); i++ {
		arr[k], arr[i] = arr[i], arr[k]
		permutation_r(arr, k+1, res)
		arr[k], arr[i] = arr[i], arr[k]
	}
}

func part1(list []int) {
	counter := 0

	maxThrust := 0
	phases := permutations(4)

	for _, phase := range phases {

		amp1 := IntCode{
			ip:      0,
			program: slices.Clone(list),
			inputs:  []int{phase[0], 0},
		}
		amp1.run(0)

		amp2 := IntCode{
			ip:      0,
			program: slices.Clone(list),
			inputs:  []int{phase[1], amp1.outputs[0]},
		}
		amp2.run(0)

		amp3 := IntCode{
			ip:      0,
			program: slices.Clone(list),
			inputs:  []int{phase[2], amp2.outputs[0]},
		}
		amp3.run(0)

		amp4 := IntCode{
			ip:      0,
			program: slices.Clone(list),
			inputs:  []int{phase[3], amp3.outputs[0]},
		}
		amp4.run(0)

		amp5 := IntCode{
			ip:      0,
			program: slices.Clone(list),
			inputs:  []int{phase[4], amp4.outputs[0]},
		}
		amp5.run(0)

		thrust := amp5.outputs[0]

		// fmt.Printf(" phase%v - [%d]\n", phase, thrust)

		if thrust > maxThrust {
			maxThrust = thrust
		}
	}
	counter = maxThrust

	fmt.Printf("Part 1: %d\n", counter)
}

func addArray(arr []int, n int) {
	for i := range arr {
		arr[i] += n
	}
}

func getChannels(phase []int) []chan int {
	chans := make([]chan int, 5)
	for i := range chans {
		chans[i] = make(chan int, 3)
		chans[i] <- phase[i]
	}
	chans[0] <- 0

	return chans
}

func part2(list []int) {
	counter := 0

	maxThrust := 0
	phases := permutations(4)

	for _, phase := range phases {
		addArray(phase, 5)
		chans := getChannels(phase)

		amps := []*IntCodeChan{}
		for i := range 5 {
			amp := IntCodeChan{
				ip:      0,
				program: slices.Clone(list),
				inputs:  chans[i],
				outputs: chans[(i+1)%5],
			}
			amps = append(amps, &amp)
		}

		var wg sync.WaitGroup
		for i := range amps {
			wg.Add(1)
			go func(a *IntCodeChan) {
				defer wg.Done()
				a.run(0)
			}(amps[i])
		}
		wg.Wait()

		// fmt.Printf(" outs[%d]\n", len(amps[4].outputs))

		// thrust := <-amps[4].outputs

		// fmt.Printf(" phase%v - [%d]\n", phase, thrust)

		for thrust := range amps[4].outputs {
			if thrust > maxThrust {
				maxThrust = thrust
			}
		}
	}
	counter = maxThrust
	fmt.Printf("Part 2: %d\n", counter)
}
