package main

import "testing"

func TestPart1(t *testing.T) {
	lines := testData()
	result := part1(lines)
	expect := 1320
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
}

func TestPart2(t *testing.T) {
	lines := testData()
	result := part2(lines)
	expect := 145
	if result != expect {
		t.Fatalf("Part2 returns %d, we expect %d", result, expect)
	}
}

func TestHash(t *testing.T) {
	data := testHashData()
	for s, expect := range data {
		result := hash(s)
		if result != expect {
			t.Fatalf("Hash for %s returns %d, we expect %d", s, result, expect)
		}
	}
}

func testData() string {
	return "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"
}

func testHashData() map[string]int {
	return map[string]int{
		"HASH": 52,
		"rn=1": 30,
		"rn":   0,
		"cm-":  253,
		"cm":   0,
		"qp=3": 97,
		"cm=2": 47,
		"qp-":  14,
		"pc=4": 180,
		"ot=9": 9,
		"ab=5": 197,
		"pc-":  48,
		"pc=6": 214,
		"ot=7": 231,
	}
}
