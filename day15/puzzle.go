package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

const day = "15"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines[0])
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines[0])
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(data string) int {
	total := 0
	for {
		d, l, ok := strings.Cut(data, ",")
		total += hash(d)
		if !ok {
			break
		}
		data = l
	}
	return total
}

func part2(data string) int {
	boxes := make([]*box, 256)
	for i := 0; i < 256; i++ {
		boxes[i] = &box{make([]*lens, 0)}
	}
	for {
		d, instr, ok := strings.Cut(data, ",")
		remove := false
		strength := ""
		var label string
		if strings.HasSuffix(d, "-") {
			label = strings.TrimSuffix(d, "-")
			remove = true
		} else {
			label, strength, _ = strings.Cut(d, "=")
		}
		bhash := hash(label)
		b := boxes[bhash]
		if remove {
			nlenses := make([]*lens, 0)
			for _, l := range b.lenses {
				if l.label != label {
					nlenses = append(nlenses, l)
				}
			}
			b.lenses = nlenses
		} else {
			str := utils.ParseInt(strength)
			nl := true
			for _, l := range b.lenses {
				if l.label == label {
					l.strength = str
					nl = false
				}
			}
			if nl {
				b.lenses = append(b.lenses, &lens{label, str})
			}
		}
		if !ok {
			break
		}
		data = instr
	}
	total := 0
	for i, b := range boxes {
		bpow := i + 1
		for k, l := range b.lenses {
			lpow := k + 1
			total += bpow * l.strength * lpow
		}
	}
	return total
}

func hash(data string) int {
	result := 0
	for _, i := range data {
		result += int(i)
		result *= 17
		result %= 256
	}
	return result
}

type (
	lens struct {
		label    string
		strength int
	}
	box struct {
		lenses []*lens
	}
)
