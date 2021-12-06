package day5

import (
	"bufio"
	"time"
)

func RunPart1(scanner *bufio.Scanner) (int, time.Duration, time.Duration) {
	inputStart := time.Now()
	input := parseInput(scanner)
	fieldSizeY, fieldSizeX := getFieldSize(input)
	playField := getPlayField(fieldSizeX, fieldSizeY)
	inputDuration := time.Since(inputStart)

	logicStart := time.Now()
	for _, rule := range input {
		if rule.From.Y == rule.To.Y || rule.From.X == rule.To.X {
			playField.drawEasyLine(rule)
		}
	}

	out := 0

	for _, rows := range playField.Field {
		for _, col := range rows {
			if col >= 2 {
				out++
			}
		}
	}
	logicDuration := time.Since(logicStart)

	return out, inputDuration, logicDuration
}

type PlayField struct {
	Field [][]int
}

func (p *PlayField) drawEasyLine(rule *Rule) {
	from := rule.From
	to := rule.To

	if from.X == to.X {
		x := from.X
		y0, y1 := findRange(from.Y, to.Y)
		for y := y0; y <= y1; y++ {
			p.Field[y][x]++
		}
	}

	if from.Y == to.Y {
		y := from.Y
		x0, x1 := findRange(from.X, to.X)
		for x := x0; x <= x1; x++ {
			p.Field[y][x]++
		}
	}
}
