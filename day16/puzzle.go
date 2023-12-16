package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

const day = "16"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	c := makeCave(lines)
	return calcEnergy(c, &beam{dir: 'e', x: -1, y: 0})
}

func part2(lines []string) int {
	c0 := makeCave(lines)
	max := 0
	for i, _ := range c0 {
		c := makeCave(lines)
		e := calcEnergy(c, &beam{dir: 'e', x: -1, y: i})
		if e > max {
			max = e
		}
		c = makeCave(lines)
		e = calcEnergy(c, &beam{dir: 'w', x: len(c[0]), y: i})
		if e > max {
			max = e
		}
	}
	for i, _ := range c0[0] {
		c := makeCave(lines)
		e := calcEnergy(c, &beam{dir: 's', x: i, y: -1})
		if e > max {
			max = e
		}
		c = makeCave(lines)
		e = calcEnergy(c, &beam{dir: 'n', x: i, y: len(c)})
		if e > max {
			max = e
		}
	}
	return max
}

func calcEnergy(c cave, b *beam) int {
	beams := make([]*beam, 0)
	beams = append(beams, b)
start:
	for len(beams) > 0 {
		b := beams[0]
		beams = beams[1:]
		for {
			switch b.dir {
			case 'e':
				if b.x < len(c[0])-1 {
					b.x++
					next := c[b.y][b.x]
					if !next.east {
						next.east = true
						if next.r == '|' {
							beams = append(beams, &beam{dir: 's', x: b.x, y: b.y})
							beams = append(beams, &beam{dir: 'n', x: b.x, y: b.y})
							goto start
						} else if next.r == '/' {
							b.dir = 'n'
						} else if next.r == '\\' {
							b.dir = 's'
						}
					} else {
						goto start
					}
				} else {
					goto start
				}
			case 'w':
				if b.x > 0 {
					b.x--
					next := c[b.y][b.x]
					if !next.west {
						next.west = true
						if next.r == '|' {
							beams = append(beams, &beam{dir: 's', x: b.x, y: b.y})
							beams = append(beams, &beam{dir: 'n', x: b.x, y: b.y})
							goto start
						} else if next.r == '/' {
							b.dir = 's'
						} else if next.r == '\\' {
							b.dir = 'n'
						}
					} else {
						goto start
					}
				} else {
					goto start
				}
			case 'n':
				if b.y > 0 {
					b.y--
					next := c[b.y][b.x]
					if !next.north {
						next.north = true
						if next.r == '-' {
							beams = append(beams, &beam{dir: 'e', x: b.x, y: b.y})
							beams = append(beams, &beam{dir: 'w', x: b.x, y: b.y})
							goto start
						} else if next.r == '/' {
							b.dir = 'e'
						} else if next.r == '\\' {
							b.dir = 'w'
						}
					} else {
						goto start
					}
				} else {
					goto start
				}
			case 's':
				if b.y < len(c)-1 {
					b.y++
					next := c[b.y][b.x]
					if !next.south {
						next.south = true
						if next.r == '-' {
							beams = append(beams, &beam{dir: 'e', x: b.x, y: b.y})
							beams = append(beams, &beam{dir: 'w', x: b.x, y: b.y})
							goto start
						} else if next.r == '/' {
							b.dir = 'w'
						} else if next.r == '\\' {
							b.dir = 'e'
						}
					} else {
						goto start
					}
				} else {
					goto start
				}
			}
		}
	}
	total := 0
	for _, row := range c {
		for _, t := range row {
			if t.north || t.south || t.east || t.west {
				total++
			}
		}
	}
	return total
}

func makeCave(lines []string) cave {
	c := make(cave, len(lines))
	for i, line := range lines {
		c[i] = make([]*tile, len(line))
		for j, r := range line {
			c[i][j] = &tile{r: r}
		}
	}
	return c
}

type (
	cave [][]*tile
	tile struct {
		r                        rune
		north, south, east, west bool
	}
	beam struct {
		x, y int
		dir  rune
	}
)
