package solutions

import (
	"aoc/solver"
	"fmt"
	"strings"
)

type DayOne struct{ solver.Solves }

func (d DayOne) PartOne() (solution any) {
	lines := d.input()
	prev := d.scan_line(lines[0])
	count := 0

	for _, line := range lines[1:] {
		depth := d.scan_line(line)
		if depth > prev {
			count++
		}
		prev = depth
	}

	solution = count
	return
}

func (d DayOne) PartTwo() (solution any) {
	lines := d.input()
	count := 0
	prev := 10000000
	for i := range lines[:len(lines)-3] {
		sum := 0
		for j := 0; j < 3; j++ {
			sum += d.scan_line(lines[i+j])
		}
		if sum > prev {
			count++
		}
		prev = sum
	}

	solution = count
	return
}

func (d DayOne) input() []string {
	return strings.Split(d.GetInput(), "\n")
}

func (d DayOne) scan_line(line string) (value int) {
	fmt.Sscanf(line, "%d", &value)
	return
}
