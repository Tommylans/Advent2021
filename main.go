package main

import (
	"fmt"
	"github.com/Tommylans/Advent2021/day1"
	"github.com/Tommylans/Advent2021/day2"
	"time"
)

func main() {

	start := time.Now()
	day1part1 := day1.StartPart1()
	fmt.Println("Day1:p1", time.Since(start))
	fmt.Println("Day1:p1 answer", day1part1)

	start = time.Now()
	day1part2 := day1.StartPart2()
	fmt.Println("Day1:p2", time.Since(start))
	fmt.Println("Day1:p2 answer", day1part2)

	start = time.Now()
	day2part1 := day2.StartPart1()
	fmt.Println("Day2:p1", time.Since(start))
	fmt.Println("Day2:p1 answer", day2part1)

	start = time.Now()
	day2part2 := day2.StartPart2()
	fmt.Println("Day2:p2", time.Since(start))
	fmt.Println("Day2:p2 answer", day2part2)

}
