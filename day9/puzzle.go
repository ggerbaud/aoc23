package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

const day = "9"

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
		data := utils.ListOfNumbers(line, " ")
		series := []history{data}
		dh := &dhistory{series: series}
		dh.derive()
		p := dh.predictNext()
		total += p
	}
	return total
}

func part2(lines []string) int {
	total := 0
	for _, line := range lines {
		data := utils.ListOfNumbers(line, " ")
		series := []history{data}
		dh := &dhistory{series: series}
		dh.derive()
		p := dh.predictPrev()
		total += p
	}
	return total
}

func (dh *dhistory) predictNext() int {
	pred := make([]int, len(dh.series))
	pred[len(pred)-1] = 0
	for i := 1; i < len(dh.series); i++ {
		data := dh.series[len(dh.series)-1-i]
		pred[len(pred)-i-1] = pred[len(pred)-i] + data[len(data)-1]
	}
	return pred[0]
}

func (dh *dhistory) predictPrev() int {
	pred := make([]int, len(dh.series))
	pred[len(pred)-1] = 0
	for i := 1; i < len(dh.series); i++ {
		data := dh.series[len(dh.series)-1-i]
		pred[len(pred)-i-1] = data[0] - pred[len(pred)-i]
	}
	return pred[0]
}

func (dh *dhistory) derive() {
	last := dh.series[len(dh.series)-1]
	for !last.isZero() {
		nserie := make(history, 0)
		for i := 0; i < len(last)-1; i++ {
			nserie = append(nserie, last[i+1]-last[i])
		}
		dh.series = append(dh.series, nserie)
		last = nserie
	}
}

func (h history) isZero() bool {
	for _, v := range h {
		if v != 0 {
			return false
		}
	}
	return true
}

type (
	history  []int
	dhistory struct {
		series     []history
		prediction []int
	}
)
