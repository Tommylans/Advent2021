package day1

func RunPart1() interface{} {
	input := parseInput()

	increases := 0

	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			increases++
		}
	}

	return increases
}
