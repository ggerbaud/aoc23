package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

const day = "19"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	wfs, idx := makeWorkflows(lines)
	in := wfs["in"]
	parts := makeParts(lines[idx:])
	accepted, rejected := make([]part, 0), make([]part, 0)
	wfs["A"] = workflow{name: "A", end: func(p part) { accepted = append(accepted, p) }}
	wfs["R"] = workflow{name: "R", end: func(p part) { rejected = append(rejected, p) }}
	for _, p := range parts {
		in.evaluate(p)
	}
	total := 0
	for _, p := range accepted {
		total += p["x"] + p["m"] + p["a"] + p["s"]
	}
	return total
}

func part2(lines []string) int {
	wfs, _ := makeWorkflows(lines)
	wfs["A"] = workflow{name: "A"}
	wfs["R"] = workflow{name: "R"}
	in := wfs["in"]
	c := constraint{minX: 1, maxX: 4000, minM: 1, maxM: 4000, minA: 1, maxA: 4000, minS: 1, maxS: 4000}
	accepted, _ := in.constraints(c)
	total := 0
	for _, constr := range accepted {
		total += constr.size()
	}
	return total
}

func makeWorkflows(lines []string) (map[string]workflow, int) {
	workflows := make(map[string]workflow)
	for i, line := range lines {
		if line == "" {
			return workflows, i + 1
		}
		wf := makeWorkflow(line)
		wf.dict = workflows
		workflows[wf.name] = wf
	}
	return workflows, -1
}
func makeWorkflow(line string) workflow {
	name, remain, _ := strings.Cut(line, "{")
	rulesData := strings.Split(remain[:len(remain)-1], ",")
	rules := make([]rule, 0)
	for _, data := range rulesData {
		rules = append(rules, makeRule(data))
	}
	return workflow{name: name, rules: rules}
}

func makeParts(lines []string) []part {
	parts := make([]part, len(lines))
	for i, line := range lines {
		parts[i] = makePart(line)
	}
	return parts
}

func makePart(line string) part {
	p := part{}
	pairs := strings.Split(strings.Trim(line, "{}"), ",")
	for _, pair := range pairs {
		k, v, _ := strings.Cut(pair, "=")
		p[k] = utils.ParseInt(v)
	}
	return p
}

func makeRule(data string) rule {
	test, out, cplx := strings.Cut(data, ":")
	dr := rule{out: test, test: func(p part) bool { return true }, splitConstraint: func(c constraint) (constraint, constraint) { return c, constraint{invalid: true} }}
	if !cplx {
		return dr
	}
	labelGt, valueGtData, greater := strings.Cut(test, ">")
	if greater {
		valueGt := utils.ParseInt(valueGtData)
		return rule{
			out:  out,
			test: func(p part) bool { return p[labelGt] > valueGt },
			splitConstraint: func(c constraint) (constraint, constraint) {
				return c.handleMin(labelGt, valueGt+1), c.handleMax(labelGt, valueGt)
			},
		}
	}
	labelLt, valueLtData, lesser := strings.Cut(test, "<")
	if lesser {
		valueLt := utils.ParseInt(valueLtData)
		return rule{
			out:  out,
			test: func(p part) bool { return p[labelLt] < valueLt },
			splitConstraint: func(c constraint) (constraint, constraint) {
				return c.handleMax(labelLt, valueLt-1), c.handleMin(labelLt, valueLt)
			},
		}
	}
	return dr
}

func (wf workflow) evaluate(p part) {
	for _, r := range wf.rules {
		if r.test(p) {
			wf.dict[r.out].evaluate(p)
			return
		}
	}
	if wf.end != nil {
		wf.end(p)
	}
}

func (wf workflow) constraints(source constraint) ([]constraint, []constraint) {
	if len(wf.rules) == 0 {
		if wf.name == "A" {
			return []constraint{source}, []constraint{}
		} else {
			return []constraint{}, []constraint{source}
		}
	}
	accepted, rejected := make([]constraint, 0), make([]constraint, 0)
	current := source
	for _, rul := range wf.rules {
		c1, c2 := rul.constraints(current)
		if !c1.invalid {
			a, r := wf.dict[rul.out].constraints(c1)
			accepted = append(accepted, a...)
			rejected = append(rejected, r...)
		}
		if !c2.invalid {
			current = c2
		} else {
			break
		}
	}
	return accepted, rejected
}

func (r rule) constraints(source constraint) (constraint, constraint) {
	return r.splitConstraint(source)
}

func (c constraint) handleMin(label string, value int) constraint {
	if c.invalid {
		return c
	}
	_min, _max := 1, 4000
	switch label {
	case "x":
		_min, _max = c.minX, c.maxX
	case "m":
		_min, _max = c.minM, c.maxM
	case "a":
		_min, _max = c.minA, c.maxA
	case "s":
		_min, _max = c.minS, c.maxS
	}
	if value > _max {
		return constraint{invalid: true}
	}
	if value > _min {
		_min = value
	}
	c2 := constraint{minX: c.minX, maxX: c.maxX, minM: c.minM, maxM: c.maxM, minA: c.minA, maxA: c.maxA, minS: c.minS, maxS: c.maxS}
	switch label {
	case "x":
		c2.minX = _min
	case "m":
		c2.minM = _min
	case "a":
		c2.minA = _min
	case "s":
		c2.minS = _min
	}
	return c2
}

func (c constraint) handleMax(label string, value int) constraint {
	if c.invalid {
		return c
	}
	_min, _max := 1, 4000
	switch label {
	case "x":
		_min, _max = c.minX, c.maxX
	case "m":
		_min, _max = c.minM, c.maxM
	case "a":
		_min, _max = c.minA, c.maxA
	case "s":
		_min, _max = c.minS, c.maxS
	}
	if value < _min {
		return constraint{invalid: true}
	}
	if value < _max {
		_max = value
	}
	c2 := constraint{minX: c.minX, maxX: c.maxX, minM: c.minM, maxM: c.maxM, minA: c.minA, maxA: c.maxA, minS: c.minS, maxS: c.maxS}
	switch label {
	case "x":
		c2.maxX = _max
	case "m":
		c2.maxM = _max
	case "a":
		c2.maxA = _max
	case "s":
		c2.maxS = _max
	}
	return c2
}

func (c constraint) size() int {
	return (c.maxX - c.minX + 1) * (c.maxM - c.minM + 1) * (c.maxA - c.minA + 1) * (c.maxS - c.minS + 1)
}

type (
	part     map[string]int
	workflow struct {
		name  string
		dict  map[string]workflow
		rules []rule
		end   func(part)
	}
	rule struct {
		test            func(part) bool
		splitConstraint func(constraint) (constraint, constraint)
		out             string
	}

	constraint struct {
		invalid                                        bool
		minX, maxX, minM, maxM, minA, maxA, minS, maxS int
	}
)
