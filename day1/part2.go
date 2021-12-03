package day1

func RunPart2() interface{} {
	input := parseInput()

	var avgList []int

	for i := 2; i < len(input); i++ {
		avgList = append(avgList, input[i]+input[i-1]+input[i-2])
	}

	increases := 0

	for i := 1; i < len(avgList); i++ {
		if avgList[i] > avgList[i-1] {
			increases++
		}
	}

	return increases
}
