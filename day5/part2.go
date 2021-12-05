package day5

import (
	"bufio"
	"math"
)

func RunPart2(scanner *bufio.Scanner) (out int) {
	input := parseInput(scanner)
	fieldSizeY, fieldSizeX := getFieldSize(input)
	playField := getPlayField(fieldSizeX, fieldSizeY)

	for _, rule := range input {
		playField.drawLine(rule)
	}

	for _, rows := range playField.Field {
		for _, col := range rows {
			if col >= 2 {
				out++
			}
		}
	}

	return
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
		p.drawPixel(x, y)
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
