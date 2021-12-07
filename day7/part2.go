package day7

import (
	"bufio"
	"math"
	"time"
)

func RunPart2(scanner *bufio.Scanner) (int, time.Duration, time.Duration) {
	inputStart := time.Now()
	input := parseInput(scanner)
	data := Data{Data: input}
	inputDuration := time.Since(inputStart)

	logicStart := time.Now()
	_, cost := data.findCheapestRoute(calcCostRising)
	logicDuration := time.Since(logicStart)

	return cost, inputDuration, logicDuration
}

func calcCostRising(input []int, des int) (total int) {
	for _, pos := range input {
		x := int(math.Abs(float64(pos - des)))

		total += x * (x + 1) / 2
	}
	return
}
