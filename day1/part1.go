package day1

import (
	"strconv"
)

func StartPart1() string {
	input := parseInput()

	increases := 0

	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			increases++
		}
	}

	return strconv.Itoa(increases)
}
