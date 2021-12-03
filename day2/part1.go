package day2

func RunPart1() interface{} {
	input := parseInput()

	hor := 0
	vert := 0

	for _, move := range input {
		hor += move.Horizontal
		vert += move.Vertical
	}

	return hor * vert
}
