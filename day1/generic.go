package day1

import (
	"bufio"
	"strconv"
)

func parseInput(scanner *bufio.Scanner) (out []int) {
	for scanner.Scan() {
		parsedNumber, _ := strconv.Atoi(scanner.Text())
		out = append(out, parsedNumber)
	}

	return
}
