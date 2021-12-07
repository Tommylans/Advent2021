package template

import (
	"bufio"
)

func parseInput(scanner *bufio.Scanner) (out []string) {
	for scanner.Scan() {
		out = append(out, scanner.Text())
	}

	return
}
