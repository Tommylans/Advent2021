package day4

import (
	"bufio"
	"time"
)

func RunPart1(scanner *bufio.Scanner) (int, time.Duration, time.Duration) {
	inputStart := time.Now()
	cards, drawn := parseInput(scanner)
	inputduration := time.Since(inputStart)

	logicStart := time.Now()
	for _, pickedNumber := range drawn {
		for _, card := range cards {
			for _, bingoNumbers := range card.Field {
				for _, bingoNumber := range bingoNumbers {
					if pickedNumber == bingoNumber.Label {
						bingoNumber.Marked = true
					}
				}
			}

			if card.CheckLineBingo() {
				logicDuration := time.Since(logicStart)
				return card.SumUnmarked() * pickedNumber, inputduration, logicDuration
			}
		}
	}

	return 0, inputduration, time.Since(logicStart)
}

func (c *Card) CheckLineBingo() bool {
	xSize := len(c.Field[0])
	ySize := len(c.Field)

	for y := 0; y < ySize; y++ {
		bingo := true

		for x := 0; x < xSize; x++ {
			if !c.Field[y][x].Marked {
				bingo = false
				break
			}
		}

		if bingo {
			return true
		}
	}

	for x := 0; x < xSize; x++ {
		bingo := true

		for y := 0; y < ySize; y++ {
			if !c.Field[y][x].Marked {
				bingo = false
				break
			}
		}

		if bingo {
			return true
		}
	}
	return false
}

func (c *Card) SumUnmarked() int {
	out := 0

	for _, bingoNumbers := range c.Field {
		for _, bingoNumber := range bingoNumbers {
			if !bingoNumber.Marked {
				out += bingoNumber.Label
			}
		}
	}
	return out
}
