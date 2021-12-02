package day2

import "strconv"

func StartPart2() string {
	input := parseInput()

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

	return strconv.Itoa(hor * vert)
}
