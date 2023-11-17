package main

import (
	solutions "aoc/solutions/2021"
	"aoc/solver"
	"embed"
	"flag"
	"fmt"
	"time"
)

//go:embed input/*
var inputs embed.FS

func main() {
	day_flag := flag.Int("day", time.Now().Day(), "Day to run")
	flag.Parse()
	day := *day_flag

	println("Solving Day", day)

	var solver solver.Solver
	switch day {
	case 1:
		solver = &solutions.DayOne{}
	case 2:
		solver = &solutions.DayTwo{}
	}
	solver.SetInput(input(day))

	var start = time.Now()
	fmt.Println("Solution part1:\t", solver.PartOne())
	fmt.Println("Elapsed:\t", time.Since(start))

	start = time.Now()
	fmt.Println("Solution part2:\t", solver.PartTwo())
	fmt.Println("Elapsed:\t", time.Since(start))
}

func input(id int) string {
	data, something := inputs.ReadFile(fmt.Sprintf("input/2021/%02d", id))
	if something != nil {
		panic(something)
	}
	return string(data)
}
