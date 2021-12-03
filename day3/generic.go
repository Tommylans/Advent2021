package day3

import (
	"bufio"
	"os"
)

func parseInput() (out []string) {
	file, _ := os.Open("./day3/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
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
