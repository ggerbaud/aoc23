package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

const day = "14"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	l := makeLever(lines)
	ln := l.moveNorth()
	return ln.weight()
}

func part2(lines []string) int {
	cycles := 1000000000
	l := makeLever(lines)
	prevs := make(map[string]int)
	prevs[l.String()] = 0
	for i := 0; i < cycles; i++ {
		l = l.cycle()
		if prev, ok := prevs[l.String()]; ok {
			remain := cycles - i
			size := i - prev
			skip := remain / size
			i += skip * size
		}
		prevs[l.String()] = i
	}
	return l.weight()
}

func makeLever(lines []string) lever {
	result := make(lever, len(lines))
	for i, line := range lines {
		result[i] = make([]int, len(line))
		for j, char := range line {
			if char == '#' {
				result[i][j] = 2
			} else if char == '.' {
				result[i][j] = 0
			} else {
				result[i][j] = 1
			}
		}
	}
	return result
}

func (l lever) weight() int {
	result := 0
	maxW := len(l)
	for _, data := range l {
		for _, value := range data {
			if value == 1 {
				result += maxW
			}
		}
		maxW--
	}
	return result
}

func (l lever) moveWest() lever {
	result := make(lever, len(l))
	for i, data := range l {
		r := make([]int, len(data))
		result[i] = r
		obstacle := 0
		for j := 0; j < len(data); j++ {
			if data[j] == 2 {
				obstacle = j + 1
				r[j] = 2
				continue
			} else if data[j] == 1 {
				r[obstacle] = 1
				obstacle++
			}
		}
	}
	return result
}

func (l lever) moveEast() lever {
	result := make(lever, len(l))
	for i, data := range l {
		r := make([]int, len(data))
		result[i] = r
		obstacle := len(data) - 1
		for j := len(data) - 1; j >= 0; j-- {
			if data[j] == 2 {
				obstacle = j - 1
				r[j] = 2
				continue
			} else if data[j] == 1 {
				r[obstacle] = 1
				obstacle--
			}
		}
	}
	return result
}

func (l lever) moveNorth() lever {
	result := make(lever, len(l))
	for i, data := range l {
		result[i] = make([]int, len(data))
	}
	for i := 0; i < len(l[0]); i++ {
		obstacle := 0
		for j := 0; j < len(l); j++ {
			if l[j][i] == 2 {
				obstacle = j + 1
				result[j][i] = 2
				continue
			} else if l[j][i] == 1 {
				result[obstacle][i] = 1
				obstacle++
			}
		}
	}
	return result
}

func (l lever) moveSouth() lever {
	result := make(lever, len(l))
	for i, data := range l {
		result[i] = make([]int, len(data))
	}
	for i := 0; i < len(l[0]); i++ {
		obstacle := len(l) - 1
		for j := len(l) - 1; j >= 0; j-- {
			if l[j][i] == 2 {
				obstacle = j - 1
				result[j][i] = 2
				continue
			} else if l[j][i] == 1 {
				result[obstacle][i] = 1
				obstacle--
			}
		}
	}
	return result
}

func (l lever) cycle() lever {
	return l.
		moveNorth().
		moveWest().
		moveSouth().
		moveEast()
}

func (l lever) String() string {
	result := ""
	for _, data := range l {
		for _, value := range data {
			if value == 0 {
				result += "."
			} else if value == 1 {
				result += "O"
			} else {
				result += "#"
			}
		}
		result += "\n"
	}
	return result
}

func (l lever) equals(o lever) bool {
	if len(l) != len(o) {
		return false
	}
	for i, data := range l {
		if len(data) != len(o[i]) {
			return false
		}
		for j, value := range data {
			if value != o[i][j] {
				return false
			}
		}
	}
	return true
}

type (
	lever [][]int
)
