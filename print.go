package main

import (
	"fmt"
	"strings"
	"time"
)

const (
	dashes = 10
)

func printHeader(text string) {
	line := strings.Repeat("-", dashes)
	fmt.Println(line+"[", text, "]"+line)
}

func printDuration(prefix string, duration time.Duration) {
	fmt.Println("Time spend on "+prefix+":", duration)
}

func printAnswer(answer interface{}) {
	fmt.Println("Answer:", answer)
}
