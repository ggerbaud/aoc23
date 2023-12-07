package main

import "testing"

func TestPart1(t *testing.T) {
	lines := testData()
	g := makeGame(lines, false)
	result := part1(g)
	expect := 6440
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
}

func TestPart2(t *testing.T) {
	lines := testData()
	g := makeGame(lines, true)
	result := part2(g)
	expect := 5905
	if result != expect {
		t.Fatalf("Part2 returns %d, we expect %d", result, expect)
	}
}

func testData() []string {
	return []string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}
}
