package day6

import (
	"bufio"
)

func parseInput(scanner *bufio.Scanner) (out []string) {
	for scanner.Scan() {
		out = append(out, scanner.Text())
	}

	return
}
