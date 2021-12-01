package day1

import "fmt"

func Start2() {
	bakfiets := parseInput()

	var avg []int

	for i := 2; i < len(bakfiets); i++ {
		avg = append(avg, bakfiets[i]+bakfiets[i-1]+bakfiets[i-2])
	}

	increases := 0

	for i := 1; i < len(avg); i++ {
		if avg[i] > avg[i-1] {
			increases++
		}
	}

	fmt.Println("Day1:p2 Answer", increases)
}
