package day6

import (
	"bufio"
	"time"
)

type FishGroup struct {
	Amount   int
	DaysLeft int
}

type FishGroupManager struct {
	FishGroups []*FishGroup
}

func RunPart2(scanner *bufio.Scanner) (int, time.Duration, time.Duration) {
	inputStart := time.Now()
	input := parseInput(scanner)
	inputDuration := time.Since(inputStart)

	logicStart := time.Now()
	fishes := [9]int{}

	for _, in := range input {
		fishes[in]++
	}

	for i := 0; i < 256; i++ {
		toBeCreated := fishes[0]
		for j := 1; j < len(fishes); j++ {
			fishes[j-1] = fishes[j]
		}

		fishes[6] += toBeCreated
		fishes[8] = toBeCreated
	}

	out := 0
	for _, fishGroup := range fishes {
		out += fishGroup
	}
	logicDuration := time.Since(logicStart)

	return out, inputDuration, logicDuration
}
