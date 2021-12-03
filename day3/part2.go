package day3

import (
	"strconv"
)

func StartPart2() string {
	input := parseInput()

	oxygen, _ := strconv.ParseInt(scrub(input, false), 2, 64)
	co2, _ := strconv.ParseInt(scrub(input, true), 2, 64)

	return strconv.Itoa(int(oxygen * co2))
}

func scrub(input []string, least bool) string {
	rest := input

	for i := 0; i < len(input[0]); i++ {
		ones := 0
		zeros := 0
		var zeroEntries []string
		var oneEntries []string

		for _, s := range rest {
			if string(s[i]) == "1" {
				oneEntries = append(oneEntries, s)
				ones += 1
			} else {
				zeroEntries = append(zeroEntries, s)
				zeros += 1
			}
		}

		if !least {
			if ones >= zeros {
				rest = oneEntries
			} else {
				rest = zeroEntries
			}
		} else {
			if len(rest) == 1 {
				break
			}

			if ones < zeros {
				rest = oneEntries
			} else {
				rest = zeroEntries
			}
		}

	}

	return rest[0]
}
