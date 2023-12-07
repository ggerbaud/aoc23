package main

import (
	"advent/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const day = "7"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	g := makeGame(lines, false)
	total := part1(g)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	jg := makeGame(lines, true)
	total = part2(jg)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(g game) int {
	sortedGame := g.Copy()
	sort.Sort(sortedGame)
	total := 0
	for i, sg := range sortedGame.p {
		total += sg.bid * (i + 1)
	}
	return total
}

func part2(g game) int {
	sortedGame := g.Copy()
	sort.Sort(sortedGame)
	total := 0
	for i, sg := range sortedGame.p {
		total += sg.bid * (i + 1)
	}
	return total
}

func makeGame(lines []string, joker bool) game {
	g := make(plays, len(lines))
	for i, line := range lines {
		g[i] = makePlay(line)
	}
	return game{g, joker}
}

func makePlay(line string) play {
	data := strings.Split(line, " ")
	return play{makeHand(data[0]), utils.ParseInt(data[1])}
}

func makeHand(data string) hand {
	h := make(hand, len(data))
	for i, c := range data {
		h[i] = makeCard(c)
	}
	return h
}

func makeCard(c rune) card {
	return card(c)
}

func (g game) Copy() game {
	_copyP := make(plays, len(g.p))
	for i, p := range g.p {
		_copyP[i] = p
	}
	return game{_copyP, g.joker}
}

func (g game) Len() int {
	return len(g.p)
}

func (g game) Less(i, j int) bool {
	ci, cj := g.p[i].cards, g.p[j].cards
	si, sj := ci.Strength(g.joker), cj.Strength(g.joker)
	if si == sj {
		for k := 0; k < len(ci); k++ {
			if ci[k].value(g.joker) != cj[k].value(g.joker) {
				return ci[k].value(g.joker) < cj[k].value(g.joker)
			}
		}
		return false
	}
	return si < sj
}

func (g game) Swap(i, j int) {
	g.p[i], g.p[j] = g.p[j], g.p[i]
}

func (h hand) Strength(joker bool) strength {
	freq := make(map[card]int)
	nj := 0
	for _, c := range h {
		if v, ok := freq[c]; ok {
			freq[c] = v + 1
		} else {
			freq[c] = 1
		}
	}
	if joker {
		nj = freq[CJ]
		delete(freq, CJ)
	}
	high := 0
	for _, i := range freq {
		if i > high {
			high = i
		}
	}
	high += nj
	if len(freq) <= 1 {
		return FiveOfAKind
	} else if len(freq) == 2 {
		if high == 4 {
			return FourOfAKind
		} else {
			return FullHouse
		}
	} else if len(freq) == 3 {
		if high == 3 {
			return ThreeOfAKind
		} else {
			return TwoPairs
		}
	} else if len(freq) == 4 {
		return OnePair
	}
	return HighCard
}

func (c card) value(joker bool) int {
	switch c {
	case C2:
		return 2
	case C3:
		return 3
	case C4:
		return 4
	case C5:
		return 5
	case C6:
		return 6
	case C7:
		return 7
	case C8:
		return 8
	case C9:
		return 9
	case CT:
		return 10
	case CJ:
		if joker {
			return 1
		}
		return 11
	case CQ:
		return 12
	case CK:
		return 13
	case CA:
		return 14
	}
	return 0
}

type (
	strength int
	card     rune
	hand     []card
	play     struct {
		cards hand
		bid   int
	}
	plays []play

	game struct {
		p     plays
		joker bool
	}
)

const (
	HighCard     strength = iota
	OnePair      strength = iota
	TwoPairs     strength = iota
	ThreeOfAKind strength = iota
	FullHouse    strength = iota
	FourOfAKind  strength = iota
	FiveOfAKind  strength = iota
	C2           card     = '2'
	C3           card     = '3'
	C4           card     = '4'
	C5           card     = '5'
	C6           card     = '6'
	C7           card     = '7'
	C8           card     = '8'
	C9           card     = '9'
	CT           card     = 'T'
	CJ           card     = 'J'
	CQ           card     = 'Q'
	CK           card     = 'K'
	CA           card     = 'A'
)
