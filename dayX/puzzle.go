package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

const day = "X"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	return 0
}

func part2(lines []string) int {
	return 0
}
