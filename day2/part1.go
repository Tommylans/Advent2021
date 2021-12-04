package day2

import "bufio"

func RunPart1(scanner *bufio.Scanner) int {
	input := parseInput(scanner)

	hor := 0
	vert := 0

	for _, move := range input {
		hor += move.Horizontal
		vert += move.Vertical
	}

	return hor * vert
}
