package day3

import (
	"bufio"
	"strconv"
	"time"
)

func RunPart1(scanner *bufio.Scanner) (int, time.Duration, time.Duration) {
	inputStart := time.Now()
	input := parseInput(scanner)
	inputDuration := time.Since(inputStart)

	logicStart := time.Now()
	gammaRateRaw := findGammaRate(input)
	gammaRate, _ := strconv.ParseInt(gammaRateRaw, 2, 64)

	epsilonRate, _ := strconv.ParseInt(flipBinary(gammaRateRaw), 2, 64)
	logicDuration := time.Since(logicStart)

	return int(gammaRate * epsilonRate), inputDuration, logicDuration
}

func findGammaRate(input []string) (out string) {
	ones := make([]int, len(input[0]))
	zeros := make([]int, len(input[0]))

	for _, s := range input {
		for i := 0; i < len(input[0]); i++ {
			if string(s[i]) == "1" {
				ones[i] += 1
			} else {
				zeros[i] += 1
			}
		}
	}

	for i := 0; i < len(ones); i++ {
		if ones[i] > zeros[i] {
			out += "1"
		} else {
			out += "0"
		}
	}

	return
}
