package day1

import (
	"bufio"
	"os"
	"strconv"
)

func parseInput() (out []int) {
	file, _ := os.Open("./day1/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parsedNumber, _ := strconv.Atoi(scanner.Text())
		out = append(out, parsedNumber)
	}

	return
}
