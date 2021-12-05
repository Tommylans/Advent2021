package day5

import (
	"bufio"
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

	playField.print()
	return
}

func (p *PlayField) drawLine(rule *Rule) {
	//x0, x1 := rule.From.X, rule.To.X
	//y0, y1 := rule.From.Y, rule.To.Y

}
