package day7

import (
	"bufio"
	"math"
	"time"
)

func RunPart1(scanner *bufio.Scanner) (int, time.Duration, time.Duration) {
	inputStart := time.Now()
	input := parseInput(scanner)
	data := Data{Data: input}
	inputDuration := time.Since(inputStart)

	logicStart := time.Now()
	_, cost := data.findCheapestRoute(calcCostStable)
	logicDuration := time.Since(logicStart)

	return cost, inputDuration, logicDuration
}

func calcCostStable(input []int, des int) (total int) {
	for _, pos := range input {
		total += int(math.Abs(float64(pos - des)))
	}
	return
}
