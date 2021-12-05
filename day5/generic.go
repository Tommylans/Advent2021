package day5

import (
	"bufio"
	"strconv"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

type Rule struct {
	From *Coordinate
	To   *Coordinate
}

func parseInput(scanner *bufio.Scanner) (out []*Rule) {
	for scanner.Scan() {
		in := scanner.Text()
		split := strings.Split(in, " -> ")

		out = append(out, &Rule{
			From: ToCoordinate(split[0]),
			To:   ToCoordinate(split[1]),
		})
	}

	return
}

func ToCoordinate(string string) *Coordinate {
	split := strings.Split(string, ",")

	x, _ := strconv.Atoi(split[0])
	y, _ := strconv.Atoi(split[1])

	return &Coordinate{
		X: x,
		Y: y,
	}
}
