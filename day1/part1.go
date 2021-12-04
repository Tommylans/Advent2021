package day1

import "bufio"

func RunPart1(scanner *bufio.Scanner) int {
	input := parseInput(scanner)

	increases := 0

	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			increases++
		}
	}

	return increases
}
