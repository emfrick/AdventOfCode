package main

// AdventOfCodeInterface accepts a string filename, and returns the answer
// to the solution or an error.
type AdventOfCodeInterface func(string) (string, error)

// AdventOfCode structure to hold the callable function as well as the input file
type AdventOfCode struct {
	run  AdventOfCodeInterface
	file string
}

// InputForDay defines the input files for each day (starting with index 1 for day 1)
var InputForDay = [...]string{
	"",
	"day01.input",
	"day02.input",
	"day03.input",
	"day04.input",
	"day05.input",
	"day06.input",
	"day07.input",
	"day08.input",
	"day09.input",
	"day10.input",
	"day11.input",
	"day12.input",
	"day13.input",
	"day14.input",
	"day15.input",
	"day16.input",
	"day17.input",
	"day18.input",
	"day19.input",
	"day20.input",
	"day21.input",
	"day22.input",
	"day23.input",
	"day24.input",
	"day25.input",
}

// AocFunctions define the array of callable
// functions paired with their inputs
var AocFunctions = [...]AdventOfCode{
	// {AocD1P1, InputForDay[1]},
	// {AocD1P2, InputForDay[1]},
	// {AocD2P1, InputForDay[2]},
	// {AocD2P2, InputForDay[2]},
	{AocD3P1, InputForDay[3]},
	{AocD3P2, InputForDay[3]},
}
