package day5

import (
	"bufio"
)

func RunPart1(scanner *bufio.Scanner) (out int) {
	input := parseInput(scanner)
	fieldSizeY, fieldSizeX := getFieldSize(input)
	playField := getPlayField(fieldSizeX, fieldSizeY)

	for _, rule := range input {
		playField.drawEasyLine(rule)
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
			p.drawPixel(x, y)
		}
	}

	if from.Y == to.Y {
		y := from.Y
		x0, x1 := findRange(from.X, to.X)
		for x := x0; x <= x1; x++ {
			p.drawPixel(x, y)
		}
	}
}
