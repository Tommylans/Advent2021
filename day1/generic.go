package day1

import (
	"os"
	"strconv"
	"strings"
)

func parseInput() []int {
	data, _ := os.ReadFile("./day1/input.txt")
	fileContent := string(data)

	splitted := strings.Split(fileContent, "\n")
	bakfiets := make([]int, len(splitted))
	for i, s := range splitted {
		bakfiets[i], _ = strconv.Atoi(s)
	}

	return bakfiets
}
