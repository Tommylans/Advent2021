package main

import (
	"fmt"
	"github.com/Tommylans/Advent2021/day1"
	"github.com/Tommylans/Advent2021/day2"
	"github.com/Tommylans/Advent2021/day3"
	"github.com/Tommylans/Advent2021/day4"
)

type Exercise struct {
	Title string
	Func  func() interface{}
}

func main() {
	exercises := []Exercise{
		{Title: "Day 1 Part 1", Func: day1.RunPart1},
		{Title: "Day 1 Part 2", Func: day1.RunPart2},
		{Title: "Day 2 Part 1", Func: day2.RunPart1},
		{Title: "Day 2 Part 2", Func: day2.RunPart2},
		{Title: "Day 3 Part 1", Func: day3.RunPart1},
		{Title: "Day 3 Part 2", Func: day3.RunPart2},
		{Title: "Day 4 Part 1", Func: day4.RunPart1},
		{Title: "Day 4 Part 2", Func: day4.RunPart2},
		//{Title: "Day 5 Part 1", Func: },
		//{Title: "Day 5 Part 2", Func: },
		//{Title: "Day 6 Part 1", Func: },
	}

	for _, exercise := range exercises {
		printHeader(exercise.Title)
		out, duration := benchFunction(exercise.Func)
		printAnswer(out)
		printDuration(duration)
		fmt.Println()
	}
}
