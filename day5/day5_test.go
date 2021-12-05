package day5

import (
	"bufio"
	"os"
	"testing"
)

const (
	inputFile   = "./input_sample.txt"
	part1Answer = 5
	part2Answer = 12
)

func TestRunPart1(t *testing.T) {
	file, _ := os.Open(inputFile)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	out := RunPart1(scanner)
	if out != part1Answer {
		t.Error("Expected:", part1Answer, " Got: ", out)
	}
}

func TestRunPart2(t *testing.T) {
	file, _ := os.Open(inputFile)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	out := RunPart2(scanner)
	if out != part2Answer {
		t.Error("Expected:", part2Answer, " Got: ", out)
	}
}
