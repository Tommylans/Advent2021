package day2

import (
	"bufio"
	"time"
)

func RunPart2(scanner *bufio.Scanner) (int, time.Duration, time.Duration) {
	inputStart := time.Now()
	input := parseInput(scanner)
	inputDuration := time.Since(inputStart)

	hor := 0
	vert := 0
	aim := 0

	logicStart := time.Now()
	for _, move := range input {
		hor += move.Horizontal
		aim += move.Vertical

		if move.Horizontal > 0 {
			vert += aim * move.Horizontal
		}
	}
	logicDuration := time.Since(logicStart)

	return hor * vert, inputDuration, logicDuration
}
