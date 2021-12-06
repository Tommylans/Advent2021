package day2

import (
	"bufio"
	"time"
)

func RunPart1(scanner *bufio.Scanner) (int, time.Duration, time.Duration) {
	inputStart := time.Now()
	input := parseInput(scanner)
	inputDuration := time.Since(inputStart)

	logicStart := time.Now()
	hor := 0
	vert := 0

	for _, move := range input {
		hor += move.Horizontal
		vert += move.Vertical
	}
	logicDuration := time.Since(logicStart)

	return hor * vert, inputDuration, logicDuration
}
