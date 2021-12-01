package main

import (
	"fmt"
	"github.com/Tommylans/Advent2021/day1"
	"time"
)

func main() {
	start := time.Now()
	day1.Start()
	fmt.Println("Day1", time.Since(start))

	start = time.Now()
	day1.Start2()
	fmt.Println("Day1:p2", time.Since(start))

}
