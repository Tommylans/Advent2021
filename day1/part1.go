package day1

import (
	"fmt"
)

func Start() {
	bakfiets := parseInput()

	increases := 0

	for i := 1; i < len(bakfiets); i++ {
		if bakfiets[i] > bakfiets[i-1] {
			increases++
		}
	}

	fmt.Println("Day1 Answer", increases)
}
