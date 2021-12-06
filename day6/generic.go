package day6

import (
	"bufio"
	"strconv"
	"strings"
)

func parseInput(scanner *bufio.Scanner) (out []int) {
	scanner.Scan()

	text := scanner.Text()
	split := strings.Split(text, ",")

	for _, s := range split {
		atoi, _ := strconv.Atoi(s)
		out = append(out, atoi)
	}

	return
}

func simulateFishGroup(input []int, days int) int {
	fishes := [9]int{}

	for _, in := range input {
		fishes[in]++
	}

	for i := 0; i < days; i++ {
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
	return out
}
