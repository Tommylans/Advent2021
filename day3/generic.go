package day3

import (
	"bufio"
)

func parseInput(scanner *bufio.Scanner) (out []string) {
	for scanner.Scan() {
		text := scanner.Text()
		out = append(out, text)
	}

	return
}

func flipBinary(input string) (out string) {
	for _, s := range input {
		if string(s) == "0" {
			out += "1"
		} else {
			out += "0"
		}
	}
	return
}
