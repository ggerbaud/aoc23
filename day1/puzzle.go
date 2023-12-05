package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

const day = "1"

var digitsMap = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	total := 0
	for _, line := range lines {
		total += getFirstLastInt(line)
	}
	return total
}

func part2(lines []string) int {
	total := 0
	for _, line := range lines {
		total += getFirstLastInt(line)
	}
	return total
}

func getFirstLastInt(s string) int {
	first, last := getFirstLast(s)
	return first*10 + last
}

func getFirstLast(s string) (int, int) {
	return getFirstDigit(s), getLastDigit(s)
}

func getFirstDigit(s string) int {
	if len(s) == 0 {
		return -1
	}
	i, hasIt := getFirstLetteredDigit(s)
	if hasIt {
		return i
	}
	c := s[0]
	if c >= '0' && c <= '9' {
		d, _ := strconv.Atoi(string(c))
		return d
	}
	return getFirstDigit(s[1:])
}

func getLastDigit(s string) int {
	if len(s) == 0 {
		return -1
	}
	i, hasIt := getLastLetteredDigit(s)
	if hasIt {
		return i
	}
	c := s[len(s)-1]
	if c >= '0' && c <= '9' {
		d, _ := strconv.Atoi(string(c))
		return d
	}
	return getLastDigit(s[:len(s)-1])
}

func getFirstLetteredDigit(s string) (int, bool) {
	if len(s) == 0 {
		return -1, false
	}
	for s2, i := range digitsMap {
		if strings.HasPrefix(s, s2) {
			return i, true
		}
	}
	return -1, false
}

func getLastLetteredDigit(s string) (int, bool) {
	if len(s) == 0 {
		return -1, false
	}
	for s2, i := range digitsMap {
		if strings.HasSuffix(s, s2) {
			return i, true
		}
	}
	return -1, false
}
