package day4

import "bufio"

func RunPart2(scanner *bufio.Scanner) int {
	cards, drawn := parseInput(scanner)

	pickedNumber, lastBingo := findLastBingo(drawn, cards)

	return lastBingo.SumUnmarked() * pickedNumber
}

func findLastBingo(drawn []int, cards []*Card) (int, *Card) {
	var bingoCards []*Card

	for _, pickedNumber := range drawn {
		for _, card := range cards {
			for _, bingoNumbers := range card.Field {
				for _, bingoNumber := range bingoNumbers {
					if pickedNumber == bingoNumber.Label {
						bingoNumber.Marked = true
					}
				}
			}

			if card.CheckLineBingo() && !card.HadBingo {
				card.HadBingo = true
				bingoCards = append(bingoCards, card)

				if len(bingoCards) == len(cards) {
					return pickedNumber, card
				}
			}
		}
	}
	return 0, &Card{}
}
