package main

import "testing"

func TestPart1(t *testing.T) {
	lines := testData()
	races := getRaces(lines)
	result := part1(races)
	expect := 288
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
}

func TestPart2(t *testing.T) {
	lines := testData()
	r := getRace(lines)
	result := part2(r)
	expect := 71503
	if result != expect {
		t.Fatalf("Part2 returns %d, we expect %d", result, expect)
	}
}

func testData() []string {
	return []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}
}
