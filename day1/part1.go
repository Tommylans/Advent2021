package day1

import (
	"bufio"
	"time"
)

func RunPart1(scanner *bufio.Scanner) (int, time.Duration, time.Duration) {
	inputStart := time.Now()
	input := parseInput(scanner)
	inputDuration := time.Since(inputStart)

	logicStart := time.Now()

	increases := 0
	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			increases++
		}
	}

	logicDuration := time.Since(logicStart)

	return increases, inputDuration, logicDuration
}
