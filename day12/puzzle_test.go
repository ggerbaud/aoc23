package main

import (
	"advent/utils"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	lines := testData()
	expects := []int{1, 4, 1, 1, 4, 10}
	for i, line := range lines {
		testPartial1(line, expects[i], t)
	}
	result := part1(lines)
	expect := 21
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
}

func TestPart2(t *testing.T) {
	lines := testData()
	expects := []int{1, 16384, 1, 16, 2500, 506250, 1}
	for i, line := range lines {
		testPartial2(line, expects[i], t)
	}
	result := part2(lines)
	expect := 525152
	if result != expect {
		t.Fatalf("Part2 returns %d, we expect %d", result, expect)
	}
}

func testPartial1(line string, expect int, t *testing.T) {
	springs, groups, _ := strings.Cut(line, " ")
	grps := utils.ListOfNumbers(groups, ",")
	ctrl := 0
	for _, grp := range grps {
		ctrl += grp + 1
	}
	ctrl -= 1
	result := solveLine(springs, grps, ctrl)
	if result != expect {
		t.Fatalf("solveLine returns %d, we expect %d", result, expect)
	}
}

func testPartial2(line string, expect int, t *testing.T) {
	springs, groups, _ := strings.Cut(line, " ")
	springs, groups = unfold(springs, groups)
	grps := utils.ListOfNumbers(groups, ",")
	ctrl := 0
	for _, grp := range grps {
		ctrl += grp + 1
	}
	ctrl -= 1
	result := solveLine(springs, grps, ctrl)
	if result != expect {
		t.Fatalf("solveLine returns %d, we expect %d", result, expect)
	}
}

func testData() []string {
	return []string{
		"???.### 1,1,3",
		".??..??...?##. 1,1,3",
		"?#?#?#?#?#?#?#? 1,3,1,6",
		"????.#...#... 4,1,1",
		"????.######..#####. 1,6,5",
		"?###???????? 3,2,1",
	}
}
