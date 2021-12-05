package day5

import (
	"bufio"
	"fmt"
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

	playField.print()

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

func (p *PlayField) drawPixel(x int, y int) {
	p.Field[y][x]++
}

func (p PlayField) print() {
	for _, ints := range p.Field {
		fmt.Println(ints)
	}
}

func findRange(a int, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}

func getPlayField(x int, y int) *PlayField {
	rows := make([][]int, y+2)
	for i := 0; i < len(rows); i++ {
		rows[i] = make([]int, x+2)
	}
	return &PlayField{Field: rows}
}

func getFieldSize(rules []*Rule) (x int, y int) {
	for _, rule := range rules {
		if rule.To.Y > y {
			y = rule.To.Y
		}
		if rule.From.Y > y {
			y = rule.From.Y
		}
		if rule.To.X > x {
			x = rule.To.X
		}
		if rule.From.X > x {
			x = rule.From.X
		}
	}

	return
}
