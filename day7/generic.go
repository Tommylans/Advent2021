package day7

import (
	"bufio"
	"github.com/Tommylans/Advent2021/tools"
	"strconv"
	"strings"
)

type Data struct {
	Data []int
}

func parseInput(scanner *bufio.Scanner) (out []int) {
	scanner.Scan()

	text := scanner.Text()
	split := strings.Split(text, ",")

	for _, s := range split {
		atoi, _ := strconv.Atoi(s)
		out = append(out, atoi)
	}

	return
}

func (d *Data) findCheapestRoute(costFunction func([]int, int) int) (cheapestPos int, cost int) {
	cost = 999999999999
	for pos := 0; pos < tools.Max(d.Data); pos++ {
		currentCost := costFunction(d.Data, pos)
		if currentCost < cost {
			cheapestPos = pos
			cost = currentCost
		}
	}
	return
}
