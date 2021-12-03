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

func printDuration(duration time.Duration) {
	fmt.Println("Time spend:", duration)
}

func printAnswer(answer interface{}) {
	fmt.Println("Answer:", answer)
}

func benchFunction(benchFunction func() interface{}) (interface{}, time.Duration) {
	now := time.Now()
	out := benchFunction()
	duration := time.Since(now)

	return out, duration
}
