package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

const day = "3"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	total := 0
	current := 0
	cl := 0
	for j, line := range lines {
		for i, char := range line {
			d, err := strconv.Atoi(string(char))
			if err != nil {
				if cl > 0 {
					if char != '.' {
						total += current
					} else if checkAround(lines, i, j, cl) {
						total += current
					}
					cl = 0
					current = 0
				}
			} else {
				cl++
				current = current*10 + d
			}
		}
		if cl > 0 && checkAround(lines, len(line), j, cl) {
			total += current
			cl = 0
			current = 0
		}
	}
	return total
}

func part2(lines []string) int {
	gears := make(map[int]map[int]int)
	total := 0
	current := 0
	cl := 0
	for j, line := range lines {
		for i, char := range line {
			d, err := strconv.Atoi(string(char))
			if err != nil {
				if cl > 0 {
					has, symb, x, y := getSymbAround(lines, i, j, cl)
					if has && symb == '*' {
						if _, ok := gears[x]; !ok {
							gears[x] = make(map[int]int)
						}
						if v, ok := gears[x][y]; !ok {
							gears[x][y] = current
						} else {
							power := current * v
							total += power
						}
					}
					cl = 0
					current = 0
				}
			} else {
				cl++
				current = current*10 + d
			}
		}
		if cl > 0 {
			has, symb, x, y := getSymbAround(lines, len(line), j, cl)
			if has && symb == '*' {
				if _, ok := gears[x]; !ok {
					gears[x] = make(map[int]int)
				}
				if v, ok := gears[x][y]; !ok {
					gears[x][y] = current
				} else {
					power := current * v
					total += power
				}
			}
		}
	}
	return total
}

func checkAround(lines []string, i, j, cl int) bool {
	x0, x1, y0, y1 := i-cl-1, i, j-1, j+1
	edges := checkChar(lines, x0, y0) ||
		checkChar(lines, x0, j) ||
		checkChar(lines, x0, y1) ||
		checkChar(lines, x1, y0) ||
		checkChar(lines, x1, j) ||
		checkChar(lines, x1, y1)

	for k := 0; k < cl; k++ {
		edges = edges || checkChar(lines, i-k-1, y0) || checkChar(lines, i-k-1, y1)
	}
	return edges
}

func getSymbAround(lines []string, i, j, cl int) (bool, byte, int, int) {
	x0, x1, y0, y1 := i-cl-1, i, j-1, j+1
	has, symb := getSymbChar(lines, x0, y0)
	if has {
		return has, symb, x0, y0
	}
	has, symb = getSymbChar(lines, x0, j)
	if has {
		return has, symb, x0, j
	}
	has, symb = getSymbChar(lines, x0, y1)
	if has {
		return has, symb, x0, y1
	}
	has, symb = getSymbChar(lines, x1, y0)
	if has {
		return has, symb, x1, y0
	}
	has, symb = getSymbChar(lines, x1, j)
	if has {
		return has, symb, x1, j
	}
	has, symb = getSymbChar(lines, x1, y1)
	if has {
		return has, symb, x1, y1
	}
	for k := 0; k < cl; k++ {
		has, symb = getSymbChar(lines, i-k-1, y0)
		if has {
			return has, symb, i - k - 1, y0
		}
		has, symb = getSymbChar(lines, i-k-1, y1)
		if has {
			return has, symb, i - k - 1, y1
		}
	}
	return false, 0, 0, 0
}

func checkChar(lines []string, x, y int) bool {
	if x < 0 || y < 0 || y >= len(lines) || x >= len(lines[y]) {
		return false
	}
	c := lines[y][x]
	return !utils.IsDigit(c) && c != '.'
}

func getSymbChar(lines []string, x, y int) (bool, byte) {
	if x < 0 || y < 0 || y >= len(lines) || x >= len(lines[y]) {
		return false, 0
	}
	c := lines[y][x]
	return !utils.IsDigit(c) && c != '.', c
}
