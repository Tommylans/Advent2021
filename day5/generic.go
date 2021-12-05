package day4

import (
	"bufio"
	"strconv"
	"strings"
)

func parseInput(scanner *bufio.Scanner) (out []*Card, numbers []int) {
	scanner.Scan()

	numbersSplit := strings.Split(scanner.Text(), ",")
	for _, s := range numbersSplit {
		atoi, _ := strconv.Atoi(s)
		numbers = append(numbers, atoi)
	}

	scanner.Scan()

	for scanner.Scan() {
		card := &Card{
			Field: [][]*BingoNumber{},
		}
		for i := 0; i < 5; i++ {
			input := strings.TrimSpace(strings.ReplaceAll(scanner.Text(), "  ", " "))
			rowNumbers := strings.Split(input, " ")
			rowNumbersParsed := make([]*BingoNumber, 5)

			for j, rowNumber := range rowNumbers {
				parsedNumber, _ := strconv.Atoi(rowNumber)
				rowNumbersParsed[j] = &BingoNumber{
					Label:  parsedNumber,
					Marked: false,
				}
			}

			card.Field = append(card.Field, rowNumbersParsed)
			scanner.Scan()
		}

		out = append(out, card)
	}

	return
}

type BingoNumber struct {
	Label  int
	Marked bool
}

type Card struct {
	HadBingo bool
	Field    [][]*BingoNumber
}
