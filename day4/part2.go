package day4

import (
	"bufio"
)

func RunPart2(scanner *bufio.Scanner) int {
	cards, drawn := parseInput(scanner)

	pickedNumber, lastBingo := findLastBingo(drawn, cards)

	return lastBingo.SumUnmarked() * pickedNumber
}

// 1.080907ms

func findLastBingo(drawn []int, cards []*Card) (int, *Card) {
	bingoCards := 0

	for _, pickedNumber := range drawn {
		for _, card := range cards {
			for _, bingoNumbers := range card.Field {
				for _, bingoNumber := range bingoNumbers {
					if pickedNumber == bingoNumber.Label {
						bingoNumber.Marked = true
					}
				}
			}

			if !card.HadBingo && card.CheckLineBingo() {
				card.HadBingo = true
				bingoCards++

				if bingoCards == len(cards) {
					return pickedNumber, card
				}
			}
		}
	}

	return 0, &Card{}
}
