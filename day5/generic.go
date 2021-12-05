package day5

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

type Rule struct {
	From *Coordinate
	To   *Coordinate
}

func parseInput(scanner *bufio.Scanner) (out []*Rule) {
	for scanner.Scan() {
		in := scanner.Text()
		split := strings.Split(in, " -> ")

		out = append(out, &Rule{
			From: ToCoordinate(split[0]),
			To:   ToCoordinate(split[1]),
		})
	}

	return
}

func ToCoordinate(string string) *Coordinate {
	split := strings.Split(string, ",")

	x, _ := strconv.Atoi(split[0])
	y, _ := strconv.Atoi(split[1])

	return &Coordinate{
		X: x,
		Y: y,
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

func (p *PlayField) drawPixel(x int, y int) {
	p.Field[y][x]++
}

func (p PlayField) print() {
	for _, ints := range p.Field {
		fmt.Println(ints)
	}
}
