package main

import "testing"

func TestPart1(t *testing.T) {
	lines := testData()
	data := makeData(lines)
	result := part1(data)
	expect := 8
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
}

func TestPart2(t *testing.T) {
	lines := testData()
	data := makeData(lines)
	result := part2(data)
	expect := 2286
	if result != expect {
		t.Fatalf("Part2 returns %d, we expect %d", result, expect)
	}
}

func TestGamePower(t *testing.T) {
	lines := testData()
	data := makeData(lines)
	doTestGamePower(t, data[1], 48)
	doTestGamePower(t, data[2], 12)
	doTestGamePower(t, data[3], 1560)
	doTestGamePower(t, data[4], 630)
	doTestGamePower(t, data[5], 36)
}

func doTestGamePower(t *testing.T, game []map[string]int, expect int) {
	result := gamePower(game)
	if result != expect {
		t.Fatalf("GamePower returns %d, we expect %d", result, expect)
	}
}

func testData() []string {
	return []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}
}
