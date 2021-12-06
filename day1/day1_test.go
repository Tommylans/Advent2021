package day1

import (
	"bufio"
	"os"
	"testing"
)

const (
	inputFile   = "./input_sample.txt"
	part1Answer = 7
	part2Answer = 5
)

func TestRunPart1(t *testing.T) {
	file, _ := os.Open(inputFile)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	out, _, _ := RunPart1(scanner)
	if out != part1Answer {
		t.Error("Expected:", part1Answer, " Got: ", out)
	}
}

func TestRunPart2(t *testing.T) {
	file, _ := os.Open(inputFile)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	out, _, _ := RunPart2(scanner)
	if out != part2Answer {
		t.Error("Expected:", part2Answer, " Got: ", out)
	}
}
