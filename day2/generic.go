package day2

import (
	"bufio"
	"strconv"
	"strings"
)

type Move struct {
	Horizontal int
	Vertical   int
}

func parseInput(scanner *bufio.Scanner) (out []Move) {
	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, " ")

		horizontal := 0
		vertical := 0

		move := split[0]

		stepSize, _ := strconv.Atoi(split[1])

		if move == "forward" {
			horizontal += stepSize
		} else if move == "down" {
			vertical += stepSize
		} else if move == "up" {
			vertical -= stepSize
		}

		out = append(out, Move{
			Horizontal: horizontal,
			Vertical:   vertical,
		})
	}

	return
}
