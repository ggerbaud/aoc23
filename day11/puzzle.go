package main

import (
	"advent/utils"
	"fmt"
	"math"
	"strconv"
)

const day = "11"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines, 1000000)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	galaxies := makeData(lines, 2)
	total := 0
	for j := 0; j < len(galaxies)-1; j++ {
		for i := j + 1; i < len(galaxies); i++ {
			gi, gj := galaxies[i], galaxies[j]
			total += int(math.Abs(float64(gi.x-gj.x)) + math.Abs(float64(gi.y-gj.y)))
		}
	}
	return total
}

func part2(lines []string, factor int) int {
	galaxies := makeData(lines, factor)
	total := 0
	for j := 0; j < len(galaxies)-1; j++ {
		for i := j + 1; i < len(galaxies); i++ {
			gi, gj := galaxies[i], galaxies[j]
			total += int(math.Abs(float64(gi.x-gj.x)) + math.Abs(float64(gi.y-gj.y)))
		}
	}
	return total
}

func makeData(lines []string, factor int) []*galaxy {
	galaxies := make([]*galaxy, 0)
	cols := make(map[int]bool)
	rows := make(map[int]bool)
	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				galaxies = append(galaxies, &galaxy{x, y})
				cols[x] = true
				rows[y] = true
			}
		}
	}
	y, x := 0, 0
	for j := 0; j < len(lines); j++ {
		if _, ok := rows[j]; !ok {
			for _, g := range galaxies {
				if (g.y - y) > j {
					g.y += factor - 1
				}
			}
			y += factor - 1
		}
	}
	for i := 0; i < len(lines[0]); i++ {
		if _, ok := cols[i]; !ok {
			for _, g := range galaxies {
				if (g.x - x) > i {
					g.x += factor - 1
				}
			}
			x += factor - 1
		}
	}
	return galaxies
}

type (
	galaxy struct {
		x, y int
	}
)
