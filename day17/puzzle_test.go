package main

import "testing"

func TestPart1(t *testing.T) {
	lines := testData()
	result := part1(lines)
	expect := 102
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
}

func TestPart2(t *testing.T) {
	lines := testData()
	result := part2(lines)
	expect := 94
	if result != expect {
		t.Fatalf("Part2 returns %d, we expect %d", result, expect)
	}
}

func testData() []string {
	return []string{
		"2413432311323",
		"3215453535623",
		"3255245654254",
		"3446585845452",
		"4546657867536",
		"1438598798454",
		"4457876987766",
		"3637877979653",
		"4654967986887",
		"4564679986453",
		"1224686865563",
		"2546548887735",
		"4322674655533",
	}
}
