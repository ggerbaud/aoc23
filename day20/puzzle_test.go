package main

import "testing"

func TestPart1(t *testing.T) {
	lines := testData()
	result := part1(lines)
	expect := 32000000
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
	lines = testData2()
	result = part1(lines)
	expect = 11687500
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
}

func TestPart2(t *testing.T) {
	lines := testData2()
	result := part2(lines)
	expect := 2
	if result != expect {
		t.Fatalf("Part2 returns %d, we expect %d", result, expect)
	}
}

func testData() []string {
	return []string{
		"broadcaster -> a, b, c",
		"%a -> b",
		"%b -> c",
		"%c -> inv",
		"&inv -> a",
	}
}

func testData2() []string {
	return []string{
		"broadcaster -> a",
		"%a -> inv, con",
		"&inv -> b",
		"%b -> con",
		"&con -> rx",
	}
}
