package day2

func RunPart2() interface{} {
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

	return hor * vert
}
