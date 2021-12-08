package day8

import (
	"bufio"
	"time"
)

func RunPart1(scanner *bufio.Scanner) (int, time.Duration, time.Duration) {
	inputStart := time.Now()
	input := parseInput(scanner)
	inputDuration := time.Since(inputStart)

	logicStart := time.Now()
	out := 0

	for _, i := range input {
		strings := i[1]
		for _, i2 := range strings {
			if len(i2) == 2 || len(i2) == 3 || len(i2) == 4 || len(i2) == 7 {
				out++
			}
		}
	}

	logicDuration := time.Since(logicStart)

	return out, inputDuration, logicDuration
}
