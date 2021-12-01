package main

import (
	"fmt"
	"github.com/Tommylans/Advent2021/day1"
	"time"
)

func main() {
	start := time.Now()
	day1.StartPart1()
	fmt.Println("Day1", time.Since(start))

	start = time.Now()
	day1.StartPart2()
	fmt.Println("Day1:p2", time.Since(start))

}
