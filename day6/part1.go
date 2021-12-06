package day6

import (
	"bufio"
	"time"
)

func RunPart1(scanner *bufio.Scanner) (int, time.Duration, time.Duration) {
	inputStart := time.Now()
	input := parseInput(scanner)
	inputDuration := time.Since(inputStart)

	logicStart := time.Now()
	out := simulateFishGroup(input, 80)
	logicDuration := time.Since(logicStart)

	return out, inputDuration, logicDuration
}
