package template

import (
	"bufio"
	"time"
)

func RunPart2(scanner *bufio.Scanner) (int, time.Duration, time.Duration) {
	inputStart := time.Now()
	inputDuration := time.Since(inputStart)

	logicStart := time.Now()
	logicDuration := time.Since(logicStart)

	return 0, inputDuration, logicDuration
}
