package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

const day = "8"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	instr := lines[0]
	n := makeNodes(lines[2:])["AAA"]
	steps := 0
	idx := 0
	for {
		if idx >= len(instr) {
			idx = 0
		}
		steps++
		switch instr[idx] {
		case 'R':
			n = n.right
		case 'L':
			n = n.left
		}
		if n.label == "ZZZ" {
			return steps
		}
		idx++
	}
}

func part2(lines []string) int {
	instr := lines[0]
	players := startNodes(makeNodes(lines[2:]), &instr)
	n := len(players)
	cycles := make([]*cycle, n)
	for i, p := range players {
		ch := findCycle(p)
		cycles[i] = ch
	}
	current := 0
	steps := 1
	for current < n {
		current = 0
	again:
		for _, c := range cycles {
			for _, end := range c.ends {
				if steps < end {
					steps = end
					continue again
				}
			}
			sfc := (steps-c.init)%c.length + c.init
			for _, end := range c.ends {
				if sfc == end {
					current++
					continue again
				}
			}
			for _, end := range c.ends {
				if sfc < end {
					delta := end - sfc
					steps += delta
					continue again
				}
			}
			k := (steps - c.init) % c.length
			steps = c.init + (k+1)*(c.length)
		}
	}
	return steps
}

func startNodes(nodes map[string]*node, instr *string) []*player {
	var starts []*player
	for _, n := range nodes {
		if n.start {
			starts = append(starts, &player{n: n, step: 0, instr: instr})
		}
	}
	return starts
}

func makeNodes(lines []string) map[string]*node {
	nodes := make(map[string]*node)
	for _, line := range lines {
		label, outs, _ := strings.Cut(line, " = ")
		n, ok := nodes[label]
		if !ok {
			n = makeNode(label)
			nodes[label] = n
		}
		left, right, _ := strings.Cut(outs, ", ")
		left = left[1:]
		right = right[:len(right)-1]
		nl, okl := nodes[left]
		if !okl {
			nl = makeNode(left)
			nodes[left] = nl
		}
		n.left = nl
		nr, okr := nodes[right]
		if !okr {
			nr = makeNode(right)
			nodes[right] = nr
		}
		n.right = nr
	}
	return nodes
}

func makeNode(label string) *node {
	n := &node{label: label}
	if strings.HasSuffix(label, "A") {
		n.start = true
	}
	if strings.HasSuffix(label, "Z") {
		n.end = true
	}
	return n
}

func findCycle(p *player) *cycle {
	dict := make(map[int]map[*node]*chainon)
	dict[0] = make(map[*node]*chainon)
	ch0 := &chainon{n: p.n}
	ch := ch0
	dict[0][p.n] = ch0
	idx := 1
	for {
		p.next()
		idxdata, ok := dict[idx]
		if !ok {
			idxdata = make(map[*node]*chainon)
			dict[idx] = idxdata
		}
		idx = (idx + 1) % len(*p.instr)
		chn, ok := idxdata[p.n]
		if !ok {
			chn = &chainon{n: p.n}
			ch.next = chn
			idxdata[p.n] = chn
			ch = chn
		} else {
			ch.next = chn
			init, length := ch0.length()
			return &cycle{first: ch0, init: init, length: length, ends: ch0.ends()}
		}
	}
}

func (p *player) next() {
	idx := p.step % len(*p.instr)
	switch (*p.instr)[idx] {
	case 'R':
		p.n = p.n.right
	case 'L':
		p.n = p.n.left
	}
	p.step++
}

func (c *chainon) ends() []int {
	data := make(map[*chainon]bool)
	ends := make([]int, 0)
	data[c] = true
	l := 1
	n := c.next
	for _, ok := data[n]; !ok; _, ok = data[n] {
		data[n] = true
		if n.n.end {
			ends = append(ends, l)
		}
		l++
		n = n.next
	}
	return ends
}

func (c *chainon) length() (int, int) {
	data := make(map[*chainon]int)
	data[c] = 0
	l := 1
	n := c.next
	for _, ok := data[n]; !ok; _, ok = data[n] {
		data[n] = l
		l++
		n = n.next
	}
	k, _ := data[n]
	return k, l - k
}

type (
	node struct {
		label      string
		left       *node
		right      *node
		start, end bool
	}

	player struct {
		n     *node
		step  int
		instr *string
	}

	cycle struct {
		first  *chainon
		init   int
		length int
		ends   []int
	}

	chainon struct {
		n    *node
		next *chainon
	}
)
