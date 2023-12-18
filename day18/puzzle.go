package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

const day = "18"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	instrs := makeDigInstrs1(lines)
	return handleInstrs(instrs)
}

func part2(lines []string) int {
	instrs := makeDigInstrs2(lines)
	return handleInstrs(instrs)
}

func handleInstrs(instrs []digInstr) int {
	lines := makeLines(instrs)
	total := 0
	for i := 0; i < len(lines); i++ {
		total += (lines[i].p1.y + lines[i].p2.y) * (lines[i].p1.x - lines[i].p2.x)
	}
	for _, l := range lines {
		total += l.length()
	}
	total /= 2
	total++
	return total
}

func makeDigInstrs1(lines []string) []digInstr {
	instrs := make([]digInstr, len(lines))
	for i, line := range lines {
		instrs[i] = makeDigInstr1(line)
	}
	return instrs
}

func makeDigInstrs2(lines []string) []digInstr {
	instrs := make([]digInstr, len(lines))
	for i, line := range lines {
		instrs[i] = makeDigInstr2(line)
	}
	return instrs
}
func makeDigInstr1(line string) digInstr {
	dir := rune(line[0])
	rdsit, _, _ := strings.Cut(line[1:], "(")
	dist, _ := strconv.Atoi(strings.TrimSpace(rdsit))
	return digInstr{dir, dist}
}
func makeDigInstr2(line string) digInstr {
	_, hexa, _ := strings.Cut(line[1:], "(")
	dist, _ := strconv.ParseInt(hexa[1:6], 16, 64)
	switch hexa[6] {
	case '0':
		return digInstr{'R', int(dist)}
	case '1':
		return digInstr{'D', int(dist)}
	case '2':
		return digInstr{'L', int(dist)}
	case '3':
		return digInstr{'U', int(dist)}
	}
	return digInstr{'R', 0}
}

func makeLines(instrs []digInstr) []line {
	lines := make([]line, 0)
	x1, x2, y1, y2 := 0, 0, 0, 0
	for _, instr := range instrs {
		switch instr.dir {
		case 'R':
			x2 = x1 + instr.dist
		case 'L':
			x2 = x1 - instr.dist
		case 'U':
			y2 = y1 - instr.dist
		case 'D':
			y2 = y1 + instr.dist
		}
		lines = append(lines, line{point{x1, y1}, point{x2, y2}})
		x1, y1 = x2, y2
	}
	return lines
}

func (l line) length() int {
	return utils.Abs(l.p1.x-l.p2.x) + utils.Abs(l.p1.y-l.p2.y)
}

type (
	digInstr struct {
		dir  rune
		dist int
	}

	point struct {
		x, y int
	}
	line struct {
		p1, p2 point
	}
)
