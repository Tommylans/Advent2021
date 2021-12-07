package tools

func Max(in []int) (max int) {
	for _, i := range in {
		if i > max {
			max = i
		}
	}
	return
}

func Sum(in []int) (sum int) {
	for _, i := range in {
		sum += i
	}
	return
}
