package day2

import "strconv"

func StartPart1() string {
	input := parseInput()

	hor := 0
	vert := 0

	for _, move := range input {
		hor += move.Horizontal
		vert += move.Vertical
	}

	return strconv.Itoa(hor * vert)
}
