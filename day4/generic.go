package day4

import (
	"bufio"
	"os"
)

func parseInput() (out []string) {
	file, _ := os.Open("./day4/input_test.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		out = append(out, text)
	}

	return
}
