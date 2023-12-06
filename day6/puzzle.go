package main

import (
	"advent/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const day = "6"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	races := getRaces(lines)
	total := part1(races)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	r := getRace(lines)
	total = part2(r)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(races []race) int {
	total := 1
	for _, r := range races {
		total *= r.countSolutions()
	}
	return total
}

func part2(r race) int {
	return r.countSolutions()
}

func getRaces(lines []string) []race {
	if len(lines) != 2 {
		panic("Invalid input")
	}
	times := utils.ListOfNumbers(strings.TrimPrefix(lines[0], "Time: "), " ")
	dists := utils.ListOfNumbers(strings.TrimPrefix(lines[1], "Distance: "), " ")
	if len(times) != len(dists) {
		panic("Invalid input")
	}
	races := make([]race, 0)
	for i, time := range times {
		races = append(races, race{time: time, record: dists[i]})
	}
	return races
}

func getRace(lines []string) race {
	if len(lines) != 2 {
		panic("Invalid input")
	}
	time, _ := strconv.Atoi(strings.ReplaceAll(strings.TrimPrefix(lines[0], "Time: "), " ", ""))
	dist, _ := strconv.Atoi(strings.ReplaceAll(strings.TrimPrefix(lines[1], "Distance: "), " ", ""))
	return race{time: time, record: dist}
}

func (r race) countSolutions() int {
	x1, x2 := r.solutions()
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	x2 = math.Floor(x2)
	x1 = math.Ceil(x1)
	return int(x2) - int(x1) + 1
}

func (r race) solutions() (float64, float64) {
	var a, b, c float64 = -1, float64(r.time), float64(-r.record - 1)
	var d = b*b - 4*a*c
	if d < 0 {
		return 0, 0
	}
	x1 := (-b + math.Sqrt(d)) / (2 * a)
	x2 := (-b - math.Sqrt(d)) / (2 * a)
	return x1, x2
}

type (
	race struct {
		time   int
		record int
	}
)
