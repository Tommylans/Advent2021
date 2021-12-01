package day1

import (
	"fmt"
)

func StartPart1() {
	input := parseInput()

	increases := 0

	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			increases++
		}
	}

	fmt.Println("Day1 Answer", increases)
}
