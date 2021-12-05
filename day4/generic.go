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

	//492.042µs
	//976.392µs

	for scanner.Scan() {
		card := &Card{
			Field: [5][5]*BingoNumber{},
		}
		for i := 0; i < 5; i++ {
			input := strings.TrimSpace(strings.ReplaceAll(scanner.Text(), "  ", " "))
			rowNumbers := strings.Split(input, " ")

			for j, rowNumber := range rowNumbers {
				parsedNumber, _ := strconv.Atoi(rowNumber)
				card.Field[i][j] = &BingoNumber{
					Label:  parsedNumber,
					Marked: false,
				}
			}
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
	Field    [5][5]*BingoNumber
}
