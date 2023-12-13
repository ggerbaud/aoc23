package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

const day = "13"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	patterns := makePatterns(lines)
	total := 0
	for _, p := range patterns {
		s, horizontal := p.symetry(0)
		if horizontal {
			total += s * 100
		} else {
			total += s
		}
	}
	return total
}

func part2(lines []string) int {
	patterns := makePatterns(lines)
	total := 0
	for _, p := range patterns {
		s, horizontal := p.symetry(1)
		if horizontal {
			total += s * 100
		} else {
			total += s
		}
	}
	return total
}

func makePatterns(lines []string) []pattern {
	rank := 0
	patterns := make([]pattern, 0)
	current := pattern{rank, make([]row, 0)}
	for _, line := range lines {
		if line == "" {
			patterns = append(patterns, current)
			rank++
			current = pattern{rank, make([]row, 0)}
			continue
		}
		r := make(row, 0)
		for _, char := range line {
			if char == '#' {
				r = append(r, 1)
			} else {
				r = append(r, 0)
			}
		}
		current.rows = append(current.rows, r)
	}
	patterns = append(patterns, current)
	return patterns
}

func (p pattern) rotate() pattern {
	newPattern := make([]row, len(p.rows[0]))
	for i := 0; i < len(p.rows[0]); i++ {
		newPattern[i] = make(row, len(p.rows))
		for j := 0; j < len(p.rows); j++ {
			newPattern[i][j] = p.rows[j][i]
		}
	}
	return pattern{p.rank, newPattern}
}

func (p pattern) symetry(ndiff int) (int, bool) {
	s, ok := p.horizontalSymetry(0, ndiff)
	if ok {
		return s, true
	}
	s, ok = p.rotate().horizontalSymetry(0, ndiff)
	if ok {
		return s, false
	}
	panic("No symetry found for pattern " + strconv.Itoa(p.rank) + " !")
}

func (p pattern) horizontalSymetry(start int, ndiff int) (int, bool) {
	for i := start; i < len(p.rows)-1; i++ {
		if p.numberOfDiffAt(i) == ndiff {
			return i + 1, true
		}
	}
	return 0, false
}

func (p pattern) numberOfDiffAt(s int) int {
	total := 0
	for j := 0; j <= s; j++ {
		if s-j >= 0 && s+1+j < len(p.rows) {
			total += p.rows[s-j].numberOfDiff(p.rows[s+1+j])
		}
	}
	return total
}

func (r row) numberOfDiff(o row) int {
	diff := 0
	for i, v := range r {
		if v != o[i] {
			diff++
		}
	}
	return diff
}

type (
	pattern struct {
		rank int
		rows []row
	}
	row []int
)
