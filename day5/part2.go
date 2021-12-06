package day5

import (
	"bufio"
	"math"
	"time"
)

func RunPart2(scanner *bufio.Scanner) (int, time.Duration, time.Duration) {
	inputStart := time.Now()
	input := parseInput(scanner)
	fieldSizeY, fieldSizeX := getFieldSize(input)
	playField := getPlayField(fieldSizeX, fieldSizeY)
	inputDuration := time.Since(inputStart)

	logicStart := time.Now()
	for _, rule := range input {
		playField.drawLine(rule)
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

func (p *PlayField) drawLine(rule *Rule) {
	x0, x1 := rule.From.X, rule.To.X
	y0, y1 := rule.From.Y, rule.To.Y

	if y0 == y1 || x0 == x1 {
		p.drawEasyLine(rule)
	} else {
		p.drawDiagLine(rule)
	}
}

func (p *PlayField) drawDiagLine(rule *Rule) {
	x0, x1 := rule.From.X, rule.To.X
	y0, y1 := rule.From.Y, rule.To.Y

	delta := x0 - x1
	dirX := Sign(delta)
	dirY := Sign(y0 - y1)

	for d := 0; d <= Abs(delta); d++ {
		x := x1 + d*dirX
		y := y1 + d*dirY
		p.Field[y][x]++
	}
}

func Sign(in int) int {
	if in < 0 {
		return -1
	}
	return 1
}

func Abs(in int) int {
	return int(math.Abs(float64(in)))
}
