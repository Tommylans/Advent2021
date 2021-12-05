package day4

import "bufio"

func RunPart1(scanner *bufio.Scanner) int {
	cards, drawn := parseInput(scanner)

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
				return card.SumUnmarked() * pickedNumber
			}
		}
	}

	return 0
}

func (c *Card) CheckLineBingo() bool {
	for _, bingoNumbers := range c.Field {
		bingo := true
		for _, bingoNumber := range bingoNumbers {
			if !bingoNumber.Marked {
				bingo = false
				break
			}
		}

		if bingo {
			return true
		}
	}

	for x := 0; x < len(c.Field[0]); x++ {
		bingo := true

		for y := 0; y < len(c.Field); y++ {
			if !c.Field[y][x].Marked {
				bingo = false
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
