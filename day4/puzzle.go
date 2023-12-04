package main

import (
	"advent/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const day = "4"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	cards := getCards(lines)
	total := part1(cards)
	// 18619
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(cards)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(cards map[int]card) int {
	total := 0
	for _, c := range cards {
		wins := c.howManyWins()
		if wins > 0 {
			total += int(math.Pow(2, float64(wins-1)))
		}
	}
	return total
}

func part2(cards map[int]card) int {
	deck := make(map[int]int)
	for i, _ := range cards {
		deck[i] = 1
	}
	for i := 0; i < len(cards); i++ {
		w := cards[i].howManyWins()
		for j := 0; j < w; j++ {
			deck[j+i+1] += deck[i]
		}
	}
	total := 0
	for _, v := range deck {
		total += v
	}
	return total
}

func getCards(lines []string) map[int]card {
	cards := make(map[int]card)
	for _, line := range lines {
		data := strings.Split(line, ":")
		id := utils.ParseInt(strings.TrimPrefix(data[0], "Card "))
		lists := strings.Split(data[1], "|")
		win := utils.ListOfNumbers(lists[0], " ")
		have := utils.ListOfNumbers(lists[1], " ")
		cards[id] = card{win: win, have: have}
	}
	return cards
}

type (
	card struct {
		win  []int
		have []int
	}
)

func (c card) howManyWins() int {
	wins := 0
	for _, i := range c.have {
		if _, ok := utils.Index(i, c.win); ok {
			wins++
		}
	}
	return wins
}
