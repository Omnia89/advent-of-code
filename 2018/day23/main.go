package main

import (
	"container/heap"
	"fmt"
	"regexp"

	"advent2018/util"
)

func main() {
	data := util.GetDataByRow("day23")
	// data := util.GetTestByRow("day23")

	list := parse(data)

	part1(list)
	part2(list)
}

type Point struct {
	x int
	y int
	z int
}

type Bot struct {
	p      Point
	radius int
}

func parse(data []string) []Bot {
	bs := []Bot{}

	capture := regexp.MustCompile(`<(\-?\d+),(\-?\d+),(\-?\d+)>.*?(\d+)`)

	for _, s := range data {
		matches := capture.FindStringSubmatch(s)
		bs = append(bs, Bot{
			p:      Point{util.ToInt(matches[1]), util.ToInt(matches[2]), util.ToInt(matches[3])},
			radius: util.ToInt(matches[4]),
		})
	}

	return bs
}

func part1(bots []Bot) {
	counter := 0

	var stronger Bot

	for _, b := range bots {
		if b.radius > stronger.radius {
			stronger = b
		}
	}

	for _, b := range bots {
		delta := util.IntAbs(b.p.x - stronger.p.x)
		delta += util.IntAbs(b.p.y - stronger.p.y)
		delta += util.IntAbs(b.p.z - stronger.p.z)

		if delta <= stronger.radius {
			counter++
		}
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func boxDistance(b *Box, p Point) int {
	dist := 0
	if p.x > b.xMax || p.x < b.xMin {
		if p.x > b.xMax {
			dist += p.x - b.xMax
		} else {
			dist += b.xMin - p.x
		}
	}
	if p.y > b.yMax || p.y < b.yMin {
		if p.y > b.yMax {
			dist += p.y - b.yMax
		} else {
			dist += b.yMin - p.y
		}
	}
	if p.z > b.zMax || p.z < b.zMin {
		if p.z > b.zMax {
			dist += p.z - b.zMax
		} else {
			dist += b.zMin - p.z
		}
	}
	return dist
}

func touchBox(b *Box, p Bot) bool {
	dist := boxDistance(b, p.p)

	return dist <= p.radius
}

func countTouches(box *Box, bots []Bot) {
	box.points = 0
	for _, b := range bots {
		if touchBox(box, b) {
			box.points += 1
		}
	}

	box.minDist = boxDistance(box, Point{0, 0, 0})
}

func splitBox(b *Box) []Box {
	boxes := make([]Box, 0, 8)
	halfX := b.xMin + (b.xMax-b.xMin)/2
	halfY := b.yMin + (b.yMax-b.yMin)/2
	halfZ := b.zMin + (b.zMax-b.zMin)/2

	boxes = append(boxes, Box{
		xMin: b.xMin,
		xMax: halfX,
		yMin: b.yMin,
		yMax: halfY,
		zMin: b.zMin,
		zMax: halfZ,
	})
	if halfX < b.xMax {
		boxes = append(boxes, Box{
			xMin: halfX + 1,
			xMax: b.xMax,
			yMin: b.yMin,
			yMax: halfY,
			zMin: b.zMin,
			zMax: halfZ,
		})
	}
	if halfY < b.yMax {
		boxes = append(boxes, Box{
			xMin: b.xMin,
			xMax: halfX,
			yMin: halfY + 1,
			yMax: b.yMax,
			zMin: b.zMin,
			zMax: halfZ,
		})
	}
	if halfX < b.xMax && halfY < b.yMax {
		boxes = append(boxes, Box{
			xMin: halfX + 1,
			xMax: b.xMax,
			yMin: halfY + 1,
			yMax: b.yMax,
			zMin: b.zMin,
			zMax: halfZ,
		})
	}

	if halfZ < b.zMax {
		boxes = append(boxes, Box{
			xMin: b.xMin,
			xMax: halfX,
			yMin: b.yMin,
			yMax: halfY,
			zMin: halfZ + 1,
			zMax: b.zMax,
		})
		if halfX < b.xMax {
			boxes = append(boxes, Box{
				xMin: halfX + 1,
				xMax: b.xMax,
				yMin: b.yMin,
				yMax: halfY,
				zMin: halfZ + 1,
				zMax: b.zMax,
			})
		}
		if halfY < b.yMax {
			boxes = append(boxes, Box{
				xMin: b.xMin,
				xMax: halfX,
				yMin: halfY + 1,
				yMax: b.yMax,
				zMin: halfZ + 1,
				zMax: b.zMax,
			})
		}
		if halfX < b.xMax && halfY < b.yMax {
			boxes = append(boxes, Box{
				xMin: halfX + 1,
				xMax: b.xMax,
				yMin: halfY + 1,
				yMax: b.yMax,
				zMin: halfZ + 1,
				zMax: b.zMax,
			})
		}
	}
	return boxes
}

func part2(bots []Bot) {
	counter := 0

	firstBox := Box{}

	for _, b := range bots {
		if b.p.x-b.radius < firstBox.xMin {
			firstBox.xMin = b.p.x - b.radius
		}
		if b.p.x+b.radius > firstBox.xMax {
			firstBox.xMax = b.p.x + b.radius
		}

		if b.p.y-b.radius < firstBox.yMin {
			firstBox.yMin = b.p.y - b.radius
		}
		if b.p.y+b.radius > firstBox.yMax {
			firstBox.yMax = b.p.y + b.radius
		}

		if b.p.z-b.radius < firstBox.zMin {
			firstBox.zMin = b.p.z - b.radius
		}
		if b.p.z+b.radius > firstBox.zMax {
			firstBox.zMax = b.p.z + b.radius
		}
		firstBox.points++
	}

	queue := make(BoxQueue, 0)
	heap.Init(&queue)
	heap.Push(&queue, &firstBox)

	var final Box
	for queue.Len() > 0 {
		box := heap.Pop(&queue).(*Box)
		volume := box.volume()
		if volume == 0 {
			panic("da faq")
		}
		if volume == 1 {
			final = *box
			break
		}
		boxes := splitBox(box)
		for _, b := range boxes {
			countTouches(&b, bots)
			heap.Push(&queue, &b)
		}
	}

	counter = final.minDist

	fmt.Printf("Part 2: %d\n", counter)
}
