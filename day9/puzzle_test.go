package main

import "testing"

func TestPart1(t *testing.T) {
	lines := testData()
	result := part1(lines)
	expect := 114
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
}

func TestPart2(t *testing.T) {
	lines := testData()
	result := part2(lines)
	expect := 2
	if result != expect {
		t.Fatalf("Part2 returns %d, we expect %d", result, expect)
	}
}

func testData() []string {
	return []string{
		"0 3 6 9 12 15",
		"1 3 6 10 15 21",
		"10 13 16 21 30 45",
	}
}
