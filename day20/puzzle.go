package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

const day = "20"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	pushes := 1000
	modules := makeModules(lines)
	highs, lows := 0, 0
	queue := make([]pulse, 0)
	for i := 0; i < pushes; i++ {
		queue = append(queue, pulse{high: false, src: "button", dest: "broadcaster"})
		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]
			if p.high {
				highs++
			} else {
				lows++
			}
			mo := modules[p.dest]
			queue = append(queue, mo.signal(p)...)
		}
	}
	return highs * lows
}

func part2(lines []string) int {
	modules := makeModules(lines)
	last := inputs(modules, "rx")[0]
	conjs := inputs(modules, last)
	cycles := make(map[string]int)
	queue := make([]pulse, 0)
	pushes := 0
	for len(cycles) < len(conjs) {
		pushes++
		queue = append(queue, pulse{high: false, src: "button", dest: "broadcaster"})
		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]
			mo := modules[p.dest]
			queue = append(queue, mo.signal(p)...)
			for _, c := range conjs {
				if _, ok := cycles[c]; !ok && modules[last].memory[c] {
					cycles[c] = pushes
				}
			}
		}
	}
	lcm := 1
	for _, v := range cycles {
		lcm = LCM(lcm, v)
	}

	return lcm
}

func makeModules(lines []string) map[string]*module {
	modules := make(map[string]*module)
	for _, line := range lines {
		namekind, outs, _ := strings.Cut(line, " -> ")
		name := namekind[1:]
		kind := 'u'
		if namekind[0] == '%' {
			kind = 'f'
		} else if namekind[0] == '&' {
			kind = 'c'
		} else if namekind == "broadcaster" {
			name = "broadcaster"
			kind = 'b'
		} else {
			name = namekind
		}
		outputs := strings.Split(outs, ", ")
		modules[name] = &module{
			kind:    kind,
			name:    name,
			state:   false,
			outputs: outputs,
			memory:  make(map[string]bool),
		}
	}
	for name, m := range modules {
		for _, output := range m.outputs {
			mo, ok := modules[output]
			if !ok {
				mo = &module{
					kind:    'u',
					name:    output,
					state:   false,
					outputs: make([]string, 0),
					memory:  make(map[string]bool),
				}
				modules[output] = mo
			}
			mo.memory[name] = false
		}
	}
	return modules
}

func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func inputs(modules map[string]*module, target string) []string {
	results := make([]string, 0)
	for _, m := range modules {
		for _, output := range m.outputs {
			if output == target {
				results = append(results, m.name)
			}
		}
	}
	return results
}

func (m *module) signal(p pulse) []pulse {
	pulses := make([]pulse, 0)
	m.memory[p.src] = p.high
	high := p.high
	switch m.kind {
	case 'f':
		if !p.high {
			m.state = !m.state
			high = m.state
		} else {
			return pulses
		}
	case 'c':
		high = false
		for _, v := range m.memory {
			if !v {
				high = true
				break
			}
		}
	}
	for _, o := range m.outputs {
		pulses = append(pulses, pulse{high: high, src: m.name, dest: o})
	}
	return pulses
}

type (
	module struct {
		kind    rune
		name    string
		state   bool
		outputs []string
		memory  map[string]bool
	}
	pulse struct {
		high bool
		src  string
		dest string
	}
)
