package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

const day = "12"

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
		springs, groups, _ := strings.Cut(line, " ")
		grps := utils.ListOfNumbers(groups, ",")
		ctrl := 0
		for _, grp := range grps {
			ctrl += grp + 1
		}
		ctrl -= 1
		total += solveLine(springs, grps, ctrl)
	}
	return total
}

func part2(lines []string) int {
	total := 0
	for _, line := range lines {
		springs, groups, _ := strings.Cut(line, " ")
		springs, groups = unfold(springs, groups)
		grps := utils.ListOfNumbers(groups, ",")
		ctrl := 0
		for _, grp := range grps {
			ctrl += grp + 1
		}
		ctrl -= 1
		total += solveLine(springs, grps, ctrl)
	}
	return total
}

func unfold(springs, groups string) (string, string) {
	rsprings, rgroups := springs, groups
	for i := 0; i < 4; i++ {
		rsprings += "?" + springs
		rgroups += "," + groups
	}
	return rsprings, rgroups
}

func solveLine(springs string, groups []int, ctrl int) int {
	sgroups := fmt.Sprintf("%v", groups)
	if m, ok := memory[springs]; ok {
		if res, ok := m[sgroups]; ok {
			return res
		}
	}
	if len(groups) == 0 {
		if strings.Contains(springs, "#") {
			return 0
		}
		return 1
	}
	if len(springs) < ctrl {
		return 0
	}
	if len(springs) == ctrl {

	}
	for strings.HasPrefix(springs, ".") {
		springs = springs[1:]
	}
	group := groups[0]
	total := 0
	for i := 0; i <= len(springs)-ctrl; i++ {
		nsprings, ok := insertDash(springs, i, group)
		if ok {
			total += solveLine(nsprings, groups[1:], ctrl-group-1)
		}
	}
	m, ok := memory[springs]
	if !ok {
		m = make(map[string]int)
		memory[springs] = m
	}
	m[sgroups] = total
	return total
}

func insertDash(springs string, prefix, n int) (string, bool) {
	for i := 0; i < prefix; i++ {
		if springs[i] == '#' {
			return "", false
		}
	}
	for i := prefix; i < prefix+n; i++ {
		if springs[i] == '.' {
			return "", false
		}
	}
	if len(springs) > prefix+n {
		if springs[prefix+n] == '#' {
			return "", false
		}
		return springs[prefix+n+1:], true
	}
	return springs[prefix+n:], true
}

var memory = make(map[string]map[string]int)
