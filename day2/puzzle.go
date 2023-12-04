package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

const day = "2"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	data := makeData(lines)
	total := part1(data)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(data)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(data map[int][]map[string]int) int {
	control := map[string]int{"red": 12, "green": 13, "blue": 14}
	total := 0
	for id, game := range data {
		if checkPossible(game, control) {
			total += id
		}
	}
	return total
}

func part2(data map[int][]map[string]int) int {
	total := 0
	for _, game := range data {
		total += gamePower(game)
	}
	return total
}

func makeData(lines []string) map[int][]map[string]int {
	games := make(map[int][]map[string]int)
	for _, line := range lines {
		x := strings.Split(line, ":")
		gameId, err := strconv.Atoi(strings.Replace(x[0], "Game ", "", -1))
		utils.CheckErrorP(err)
		checks := strings.Split(x[1], ";")
		counts := make([]map[string]int, 0)
		for _, check := range checks {
			dices := strings.Split(check, ",")
			count := make(map[string]int)
			for _, dice := range dices {
				for _, color := range colors {
					if strings.HasSuffix(dice, color) {
						count[color], err = strconv.Atoi(strings.TrimSpace(strings.ReplaceAll(dice, color, "")))
					}
				}
			}
			counts = append(counts, count)
		}
		games[gameId] = counts
	}
	return games
}

func checkPossible(game []map[string]int, control map[string]int) bool {
	for _, counts := range game {
		for color, n := range control {
			if counts[color] > n {
				return false
			}
		}
	}
	return true
}

func gamePower(game []map[string]int) int {
	minis := minimums(game)
	power := 1
	for _, n := range minis {
		power *= n
	}
	return power
}

func minimums(game []map[string]int) map[string]int {
	minis := map[string]int{"red": 0, "green": 0, "blue": 0}
	for _, counts := range game {
		for color, n := range counts {
			minis[color] = utils.Max(n, minis[color])
		}
	}
	return minis
}

var (
	colors = []string{"red", "blue", "green"}
)
