package main

import "testing"

func TestPart1(t *testing.T) {
	lines := testData()
	result := part1(lines)
	expect := 4
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
	lines = testData2()
	result = part1(lines)
	expect = 8
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
}

func TestPart2(t *testing.T) {
	lines := testData3()
	result := part2(lines)
	expect := 4
	if result != expect {
		t.Fatalf("Part2 returns %d, we expect %d", result, expect)
	}
	lines = testData3bis()
	result = part2(lines)
	expect = 4
	if result != expect {
		t.Fatalf("Part2 returns %d, we expect %d", result, expect)
	}
	lines = testData4()
	result = part2(lines)
	expect = 8
	if result != expect {
		t.Fatalf("Part2 returns %d, we expect %d", result, expect)
	}
	lines = testData5()
	result = part2(lines)
	expect = 10
	if result != expect {
		t.Fatalf("Part2 returns %d, we expect %d", result, expect)
	}
}

func testData() []string {
	return []string{
		"-L|F7",
		"7S-7|",
		"L|7||",
		"-L-J|",
		"L|-JF",
	}
}

func testData2() []string {
	return []string{
		"7-F7-",
		".FJ|7",
		"SJLL7",
		"|F--J",
		"LJ.LJ",
	}
}

func testData3() []string {
	return []string{
		"...........",
		".S-------7.",
		".|F-----7|.",
		".||.....||.",
		".||.....||.",
		".|L-7.F-J|.",
		".|..|.|..|.",
		".L--J.L--J.",
		"...........",
	}
}

func testData3bis() []string {
	return []string{
		"..........",
		".S------7.",
		".|F----7|.",
		".||....||.",
		".||....||.",
		".|L-7F-J|.",
		".|..||..|.",
		".L--JL--J.",
		"..........",
	}
}

func testData4() []string {
	return []string{
		".F----7F7F7F7F-7....",
		".|F--7||||||||FJ....",
		".||.FJ||||||||L7....",
		"FJL7L7LJLJ||LJ.L-7..",
		"L--J.L7...LJS7F-7L7.",
		"....F-J..F7FJ|L7L7L7",
		"....L7.F7||L7|.L7L7|",
		".....|FJLJ|FJ|F7|.LJ",
		"....FJL-7.||.||||...",
		"....L---J.LJ.LJLJ...",
	}
}

func testData5() []string {
	return []string{
		"FF7FSF7F7F7F7F7F---7",
		"L|LJ||||||||||||F--J",
		"FL-7LJLJ||||||LJL-77",
		"F--JF--7||LJLJ7F7FJ-",
		"L---JF-JLJ.||-FJLJJ7",
		"|F|F-JF---7F7-L7L|7|",
		"|FFJF7L7F-JF7|JL---7",
		"7-L-JL7||F7|L7F-7F7|",
		"L.L7LFJ|||||FJL7||LJ",
		"L7JLJL-JLJLJL--JLJ.L",
	}
}
