package main

import "testing"

func TestPart1(t *testing.T) {
	lines := testData()
	result := part1(lines)
	expect := 136
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
}

func TestPart2(t *testing.T) {
	lines := testData()
	result := part2(lines)
	expect := 0
	if result != expect {
		t.Fatalf("Part2 returns %d, we expect %d", result, expect)
	}
}

func TestMoveNorth(t *testing.T) {
	lines := testData()
	l := makeLever(lines)
	result := l.moveNorth()
	eLines := dataMoveNorthExpect()
	expect := makeLever(eLines)
	if !result.equals(expect) {
		t.Fatalf("moveNorth returns:\n%s\n\nwe expect:\n%s", result, expect)
	}
}

func TestMoveWest(t *testing.T) {
	lines := testData()
	l := makeLever(lines)
	result := l.moveWest()
	eLines := dataMoveWestExpect()
	expect := makeLever(eLines)
	if !result.equals(expect) {
		t.Fatalf("moveWest returns:\n%s\n\nwe expect:\n%s", result, expect)
	}
}

func TestMoveEast(t *testing.T) {
	lines := testData()
	l := makeLever(lines)
	result := l.moveEast()
	eLines := dataMoveEastExpect()
	expect := makeLever(eLines)
	if !result.equals(expect) {
		t.Fatalf("moveEast returns:\n%s\n\nwe expect:\n%s", result, expect)
	}
}

func TestMoveSouth(t *testing.T) {
	lines := testData()
	l := makeLever(lines)
	result := l.moveSouth()
	eLines := dataMoveSouthExpect()
	expect := makeLever(eLines)
	if !result.equals(expect) {
		t.Fatalf("moveSouth returns:\n%s\n\nwe expect:\n%s", result, expect)
	}
}

func TestMoveCycle(t *testing.T) {
	lines := testData()
	l := makeLever(lines)
	result := l.cycle()
	c1Lines := dataOneCycleExpect()
	expect := makeLever(c1Lines)
	if !result.equals(expect) {
		t.Fatalf("cycle returns:\n%s\n\nwe expect:\n%s", result, expect)
	}
	result = result.cycle()
	c2Lines := dataTwoCycleExpect()
	expect = makeLever(c2Lines)
	if !result.equals(expect) {
		t.Fatalf("cycle returns:\n%s\n\nwe expect:\n%s", result, expect)
	}
	result = result.cycle()
	c3Lines := dataThreeCycleExpect()
	expect = makeLever(c3Lines)
	if !result.equals(expect) {
		t.Fatalf("cycle returns:\n%s\n\nwe expect:\n%s", result, expect)
	}
}

func testData() []string {
	return []string{
		"O....#....",
		"O.OO#....#",
		".....##...",
		"OO.#O....O",
		".O.....O#.",
		"O.#..O.#.#",
		"..O..#O..O",
		".......O..",
		"#....###..",
		"#OO..#....",
	}
}

func dataMoveWestExpect() []string {
	return []string{
		"O....#....",
		"OOO.#....#",
		".....##...",
		"OO.#OO....",
		"OO......#.",
		"O.#O...#.#",
		"O....#OO..",
		"O.........",
		"#....###..",
		"#OO..#....",
	}
}

func dataMoveEastExpect() []string {
	return []string{
		"....O#....",
		".OOO#....#",
		".....##...",
		".OO#....OO",
		"......OO#.",
		".O#...O#.#",
		"....O#..OO",
		".........O",
		"#....###..",
		"#..OO#....",
	}
}

func dataMoveNorthExpect() []string {
	return []string{
		"OOOO.#.O..",
		"OO..#....#",
		"OO..O##..O",
		"O..#.OO...",
		"........#.",
		"..#....#.#",
		"..O..#.O.O",
		"..O.......",
		"#....###..",
		"#....#....",
	}
}

func dataMoveSouthExpect() []string {
	return []string{
		".....#....",
		"....#....#",
		"...O.##...",
		"...#......",
		"O.O....O#O",
		"O.#..O.#.#",
		"O....#....",
		"OO....OO..",
		"#OO..###..",
		"#OO.O#...O",
	}
}
func dataOneCycleExpect() []string {
	return []string{
		".....#....",
		"....#...O#",
		"...OO##...",
		".OO#......",
		".....OOO#.",
		".O#...O#.#",
		"....O#....",
		"......OOOO",
		"#...O###..",
		"#..OO#....",
	}
}

func dataTwoCycleExpect() []string {
	return []string{
		".....#....",
		"....#...O#",
		".....##...",
		"..O#......",
		".....OOO#.",
		".O#...O#.#",
		"....O#...O",
		".......OOO",
		"#..OO###..",
		"#.OOO#...O",
	}
}

func dataThreeCycleExpect() []string {
	return []string{
		".....#....",
		"....#...O#",
		".....##...",
		"..O#......",
		".....OOO#.",
		".O#...O#.#",
		"....O#...O",
		".......OOO",
		"#...O###.O",
		"#.OOO#...O",
	}
}
