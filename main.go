package main

import (
	"bufio"
	"fmt"
	"github.com/Tommylans/Advent2021/day1"
	"github.com/Tommylans/Advent2021/day2"
	"github.com/Tommylans/Advent2021/day3"
	"github.com/Tommylans/Advent2021/day4"
	"github.com/Tommylans/Advent2021/day5"
	"os"
	"strings"
	"time"
)

type Exercise struct {
	Title   string
	Dataset string
	Func    func(scanner *bufio.Scanner) (int, time.Duration, time.Duration)
}

func main() {
	exercises := []*Exercise{
		{Title: "Day 1 Part 1", Func: day1.RunPart1, Dataset: "./day1/input.txt"},
		{Title: "Day 1 Part 2", Func: day1.RunPart2, Dataset: "./day1/input.txt"},

		{Title: "Day 2 Part 1", Func: day2.RunPart1, Dataset: "./day2/input.txt"},
		{Title: "Day 2 Part 2", Func: day2.RunPart2, Dataset: "./day2/input.txt"},

		{Title: "Day 3 Part 1", Func: day3.RunPart1, Dataset: "./day3/input.txt"},
		{Title: "Day 3 Part 2", Func: day3.RunPart2, Dataset: "./day3/input.txt"},

		{Title: "Day 4 Part 1", Func: day4.RunPart1, Dataset: "./day4/input.txt"},
		{Title: "Day 4 Part 2", Func: day4.RunPart2, Dataset: "./day4/input.txt"},

		{Title: "Day 5 Part 1", Func: day5.RunPart1, Dataset: "./day5/input.txt"},
		{Title: "Day 5 Part 2", Func: day5.RunPart2, Dataset: "./day5/input.txt"},
	}

	for _, exercise := range exercises {
		printHeader(exercise.Title)

		file, _ := os.ReadFile(exercise.Dataset)
		reader := strings.NewReader(string(file))
		scanner := bufio.NewScanner(reader)

		out, initDuration, answerDuration := exercise.Func(scanner)

		printAnswer(out)
		printDuration("init", initDuration)
		printDuration("logic", answerDuration)
		fmt.Println()
	}
}
