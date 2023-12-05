package main

import "testing"

func TestPart1(t *testing.T) {
	lines := testData()
	result := part1(lines)
	expect := 142
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
}

func TestPart2(t *testing.T) {
	lines := testData2()
	result := part2(lines)
	expect := 281
	if result != expect {
		t.Fatalf("Part2 returns %d, we expect %d", result, expect)
	}
}

func TestFisrtLastFunc(t *testing.T) {
	testFirstLast(t, "1abc2", 1, 2)
	testFirstLast(t, "pqr3stu8vwx", 3, 8)
	testFirstLast(t, "a1b2c3d4e5f", 1, 5)
	testFirstLast(t, "treb7uchet", 7, 7)
}

func testFirstLast(t *testing.T, s string, first, last int) {
	f, l := getFirstLast(s)
	if f != first || l != last {
		t.Fatalf("getFirstLast returns %d and %d, we expect %d and %d", f, l, first, last)
	}
}

func testData() []string {
	return []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
}

func testData2() []string {
	return []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}
}
