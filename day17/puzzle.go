package main

import (
	"advent/utils"
	"fmt"
	"image"
	"math"
	"strconv"
)

const day = "17"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	hm, out := makeHeatmap(lines)
	return hm.solve(out, 1, 3)
}

func part2(lines []string) int {
	hm, out := makeHeatmap(lines)
	return hm.solve(out, 4, 10)
}

func (hm heatmap) solve(out image.Point, min, max int) int {
	queue, seen := utils.PriorityQueue[path]{}, map[path]struct{}{}
	queue.GPush(path{pt: image.Point{}, dir: image.Point{X: 1}}, 0)
	queue.GPush(path{pt: image.Point{}, dir: image.Point{Y: 1}}, 0)

	for len(queue) > 0 {
		p, loss := queue.GPop()

		if p.pt == out {
			return loss
		}
		if _, ok := seen[p]; ok {
			continue
		}
		seen[p] = struct{}{}

		for i := -max; i <= max; i++ {
			n := p.pt.Add(p.dir.Mul(i))
			if _, ok := hm[n]; !ok || i > -min && i < min {
				continue
			}
			nl, s := 0, int(math.Copysign(1, float64(i)))
			for j := s; j != i+s; j += s {
				nl += hm[p.pt.Add(p.dir.Mul(j))]
			}
			queue.GPush(path{pt: n, dir: image.Point{X: p.dir.Y, Y: p.dir.X}}, loss+nl)
		}
	}
	return -1
}

func makeHeatmap(lines []string) (heatmap, image.Point) {
	hm := make(heatmap)
	end := image.Point{}
	for j, line := range lines {
		for i, c := range line {
			hm[image.Point{X: i, Y: j}] = int(c - '0')
			end = image.Point{X: i, Y: j}
		}
	}
	return hm, end
}

type (
	heatmap map[image.Point]int
	path    struct {
		pt, dir image.Point
	}
)
