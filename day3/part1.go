package day3

import (
	"bufio"
	"strconv"
)

func RunPart1(scanner *bufio.Scanner) int {
	input := parseInput(scanner)

	gammaRateRaw := findGammaRate(input)
	gammaRate, _ := strconv.ParseInt(gammaRateRaw, 2, 64)

	epsilonRate, _ := strconv.ParseInt(flipBinary(gammaRateRaw), 2, 64)

	return int(gammaRate * epsilonRate)
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
