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
