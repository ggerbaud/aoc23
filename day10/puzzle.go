package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

const day = "10"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	start, m := makeMaze(lines)
	mp := *m
	max := len(mp) * len(mp[0])
	pp1, pp2 := start, start
	var p1, p2 *pipe
	if start.n != nil {
		start.hasN = true
		p1 = start.n
	}
	if start.s != nil {
		start.hasS = true
		p1 = start.s
	}
	if start.e != nil {
		start.hasE = true
		p1 = start.e
	}
	if start.w != nil {
		start.hasW = true
		p1 = start.w
	}
	p2 = start.other(p1)
	for i := 1; i < max; i++ {
		pp1, p1 = p1, p1.other(pp1)
		pp2, p2 = p2, p2.other(pp2)
		if p1 == p2 {
			return i + 1
		}
	}
	return 0
}

func part2(lines []string) int {
	start, m := makeMaze(lines)
	mp := *m
	// clean maze
	cleanM := make([][]rune, len(mp))
	for i := 0; i < len(mp); i++ {
		cleanM[i] = make([]rune, len(mp[i]))
	}
	cleanM[start.y][start.x] = start.symb
	if start.n != nil {
		start.hasN = true
	}
	if start.s != nil {
		start.hasS = true
	}
	if start.e != nil {
		start.hasE = true
	}
	if start.w != nil {
		start.hasW = true
	}
	var current, prev *pipe
	prev = start
	current = start.other(nil)
	path := make([]point, 0)
	path = append(path, point{x: start.x, y: start.y})
	for current != start {
		path = append(path, point{x: current.x, y: current.y})
		cleanM[current.y][current.x] = current.symb
		prev, current = current, current.other(prev)
	}
	poly := &polygon{path: path}
	poly.computeMinMax()
	total := 0
	for y := 0; y < len(mp); y++ {
		for x := 0; x < len(mp[y]); x++ {
			if cleanM[y][x] != 0 {
				continue
			}
			if poly.isInside(x, y) {
				total++
			}
		}
	}
	return total
}

func makeMaze(lines []string) (*pipe, *maze) {
	m := make(maze, len(lines))
	for i := 0; i < len(lines); i++ {
		m[i] = make([]*pipe, len(lines[i]))
	}
	var start *pipe
	for y, line := range lines {
		for x, c := range line {
			p := m[y][x]
			if p == nil {
				p = &pipe{x: x, y: y}
				m[y][x] = p
			}
			p.symb = c
			var pn, ps, pw, pe *pipe
			xn, yn, xe, ye, xs, ys, xw, yw := -1, -1, -1, -1, -1, -1, -1, -1
			if c == '-' {
				xw, yw, xe, ye = x-1, y, x+1, y
				p.hasW, p.hasE = true, true
			} else if c == '|' {
				xn, yn, xs, ys = x, y-1, x, y+1
				p.hasN, p.hasS = true, true
			} else if c == 'L' {
				xn, yn, xe, ye = x, y-1, x+1, y
				p.hasN, p.hasE = true, true
			} else if c == 'J' {
				xw, yw, xn, yn = x-1, y, x, y-1
				p.hasW, p.hasN = true, true
			} else if c == 'F' {
				xs, ys, xe, ye = x, y+1, x+1, y
				p.hasS, p.hasE = true, true
			} else if c == '7' {
				xw, yw, xs, ys = x-1, y, x, y+1
				p.hasW, p.hasS = true, true
			} else if c == 'S' {
				start = p
				continue
			} else {
				continue
			}
			if yn >= 0 {
				pn = m[yn][xn]
				if pn == nil {
					pn = &pipe{x: xn, y: yn}
					m[yn][xn] = pn
				}
				pn.s = p
				p.n = pn
			}
			if ys > 0 && ys < len(m) {
				ps = m[ys][xs]
				if ps == nil {
					ps = &pipe{x: xs, y: ys}
					m[ys][xs] = ps
				}
				ps.n = p
				p.s = ps
			}
			if xw >= 0 {
				pw = m[yw][xw]
				if pw == nil {
					pw = &pipe{x: xw, y: yw}
					m[yw][xw] = pw
				}
				pw.e = p
				p.w = pw
			}
			if xe > 0 && xe < len(m[y]) {
				pe = m[ye][xe]
				if pe == nil {
					pe = &pipe{x: xe, y: ye}
					m[ye][xe] = pe
				}
				pe.w = p
				p.e = pe
			}
		}
	}
	return start, &m
}

func (p *polygon) isInside(x, y int) bool {
	if x < p.minX || x > p.maxX || y < p.minY || y > p.maxY {
		return false
	}

	inside := false
	for i, j := 0, len(p.path)-1; i < len(p.path); i++ {
		if (p.path[i].y > y) != (p.path[j].y > y) &&
			x < (p.path[j].x-p.path[i].x)*(y-p.path[i].y)/(p.path[j].y-p.path[i].y)+p.path[i].x {
			inside = !inside
		}
		j = i
	}
	return inside
}

func (p *polygon) computeMinMax() {
	p.minX = p.path[0].x
	p.maxX = p.path[0].x
	p.minY = p.path[0].y
	p.maxY = p.path[0].y
	for n := 1; n < len(p.path); n++ {
		q := p.path[n]
		p.minX = utils.Min(q.x, p.minX)
		p.maxX = utils.Max(q.x, p.maxX)
		p.minY = utils.Min(q.y, p.minY)
		p.maxY = utils.Max(q.y, p.maxY)
	}
}

func (p *pipe) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func (p *pipe) other(prev *pipe) *pipe {
	if p.n != prev && p.hasN {
		return p.n
	}
	if p.s != prev && p.hasS {
		return p.s
	}
	if p.e != prev && p.hasE {
		return p.e
	}
	if p.w != prev && p.hasW {
		return p.w
	}
	panic("No other pipe")
}

type (
	maze [][]*pipe
	pipe struct {
		symb                   rune
		x, y                   int
		hasN, hasS, hasE, hasW bool
		n, s, e, w             *pipe
	}

	point struct {
		x, y int
	}
	polygon struct {
		path                   []point
		minX, maxX, minY, maxY int
	}
)
