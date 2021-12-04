package day2

import "bufio"

func RunPart2(scanner *bufio.Scanner) int {
	input := parseInput(scanner)

	hor := 0
	vert := 0
	aim := 0

	for _, move := range input {
		hor += move.Horizontal
		aim += move.Vertical

		if move.Horizontal > 0 {
			vert += aim * move.Horizontal
		}
	}

	return hor * vert
}
