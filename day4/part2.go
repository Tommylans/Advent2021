package day4

import (
	"bufio"
	"time"
)

func RunPart2(scanner *bufio.Scanner) (int, time.Duration, time.Duration) {
	inputStart := time.Now()
	cards, drawn := parseInput(scanner)
	inputDuration := time.Since(inputStart)

	logicStart := time.Now()
	pickedNumber, lastBingo := findLastBingo(drawn, cards)
	logicDuration := time.Since(logicStart)

	return lastBingo.SumUnmarked() * pickedNumber, inputDuration, logicDuration
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
