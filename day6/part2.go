package day6

import (
	"bufio"
	"time"
)

func RunPart2(scanner *bufio.Scanner) (int, time.Duration, time.Duration) {
	inputStart := time.Now()
	input := parseInput(scanner)
	inputDuration := time.Since(inputStart)

	logicStart := time.Now()
	out := simulateFishGroup(input, 256)
	logicDuration := time.Since(logicStart)

	return out, inputDuration, logicDuration
}
