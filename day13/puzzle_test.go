package main

import "testing"

func TestPart1(t *testing.T) {
	lines := testData()
	result := part1(lines)
	expect := 405
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
	lines = testData2()
	result = part1(lines)
	expect = 400
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
}

func TestPart2(t *testing.T) {
	lines := testData()
	result := part2(lines)
	expect := 400
	if result != expect {
		t.Fatalf("Part2 returns %d, we expect %d", result, expect)
	}
}

func testData() []string {
	return []string{
		"#.##..##.",
		"..#.##.#.",
		"##......#",
		"##......#",
		"..#.##.#.",
		"..##..##.",
		"#.#.##.#.",
		"",
		"#...##..#",
		"#....#..#",
		"..##..###",
		"#####.##.",
		"#####.##.",
		"..##..###",
		"#....#..#",
	}
}
func testData2() []string {
	return []string{
		".#..",
		"##..",
		"##..",
		"..#.",
		"..#.",
		"##..",
		"##..",
		".#..",
	}
}
