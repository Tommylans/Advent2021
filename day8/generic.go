package day8

import (
	"bufio"
	"strings"
)

func parseInput(scanner *bufio.Scanner) (out [][2][]string) {
	for scanner.Scan() {
		test := strings.SplitN(scanner.Text(), " | ", 2)
		lists := [2][]string{}

		for i, s := range test {
			split := strings.Split(s, " ")

			lists[i] = split
		}

		out = append(out, lists)
	}

	return
}
