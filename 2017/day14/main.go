package main

import (
	"fmt"
	"slices"
	"strings"

	"advent2017/util"
)

func main() {
	data := util.GetDataByRow("day14")
	//data := util.GetTestByRow("day14")

	part1(data[0])
	part2(data[0])
}

func buildArray(n int) []int {
	r := make([]int, n)

	for i := range r {
		r[i] = i
	}

	return r
}

func swap(values *[]int, w int, start int, end int) {
	half := w / 2
	n := len(*values)

	for i := range half {
		a := (start + i) % n
		b := ((end-i)%n + n) % n
		(*values)[a], (*values)[b] = (*values)[b], (*values)[a]
	}
}

func hashKnot(data string) string {
	list := []byte(data)
	list = append(list, []byte{17, 31, 73, 47, 23}...)

	values := buildArray(256)

	position := 0
	skip := 0
	for range 64 {
		for _, b := range list {
			n := int(b)
			if n > 1 {
				swap(&values, n, position, (position+n-1)%len(values))
			}
			position += skip + n
			position %= len(values)
			skip++
		}
	}

	res := []int{}
	t := 0
	for t < 16 {
		i := t * 16
		v := values[i]
		for j := i + 1; j < i+16; j++ {
			v ^= values[j]
		}
		res = append(res, v)
		t++
	}

	var sb strings.Builder
	for _, n := range res {
		fmt.Fprintf(&sb, "%02x", n)
	}

	return sb.String()
}

func hexToBinArray(val string) []int {
	conv := map[rune]string{
		'0': "0000",
		'1': "0001",
		'2': "0010",
		'3': "0011",
		'4': "0100",
		'5': "0101",
		'6': "0110",
		'7': "0111",
		'8': "1000",
		'9': "1001",
		'a': "1010",
		'b': "1011",
		'c': "1100",
		'd': "1101",
		'e': "1110",
		'f': "1111",
	}

	var sb strings.Builder

	for _, c := range val {
		sb.WriteString(conv[c])
	}

	return util.StringToIntSlice(sb.String(), "")
}

func getGrid(data string) [][]int {
	bits := [][]int{}

	for i := range 128 {
		s := fmt.Sprintf("%s-%d", data, i)
		hash := hashKnot(s)

		bits = append(bits, []int{})

		bits[i] = hexToBinArray(hash)
	}
	return bits
}

func part1(data string) {
	counter := 0

	bits := getGrid(data)

	for _, v := range bits {
		for _, n := range v {
			counter += n
		}
	}

	fmt.Printf("Part 1: %d\n", counter)
}

type Point struct {
	x int
	y int
}

func getNears(p Point) []Point {
	ps := []Point{}

	if p.x > 0 {
		ps = append(ps, Point{p.x - 1, p.y})
	}
	if p.x < 127 {
		ps = append(ps, Point{p.x + 1, p.y})
	}

	if p.y > 0 {
		ps = append(ps, Point{p.x, p.y - 1})
	}
	if p.y < 127 {
		ps = append(ps, Point{p.x, p.y + 1})
	}
	return ps
}

func part2(data string) {
	counter := 0

	bits := getGrid(data)

	notDone := []Point{}
	for y, v := range bits {
		for x, n := range v {
			if n == 1 {
				notDone = append(notDone, Point{x, y})
			}
		}
	}

	for len(notDone) > 1 {
		var p Point
		q := []Point{notDone[0]}

		for len(q) > 0 {
			p, q = q[0], q[1:]
			nears := getNears(p)

			if slices.Contains(notDone, p) {
				notDone = slices.DeleteFunc(notDone, func(a Point) bool {
					return a == p
				})
			}

			for _, pp := range nears {
				if slices.Contains(notDone, pp) {
					if !slices.Contains(q, pp) {
						q = append(q, pp)
						notDone = slices.DeleteFunc(notDone, func(a Point) bool {
							return a == pp
						})
					}
				}
			}
		}
		counter++
	}

	fmt.Printf("Part 2: %d\n", counter)
}
