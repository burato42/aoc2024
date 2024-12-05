package main

import (
	"aoc2024/day5"
	"fmt"
)

func main() {
	fmt.Println(day5.SumPageNumbers(day5.ReadTextFile("./day5/sample.txt")))
	fmt.Println(day5.SumPageNumbers(day5.ReadTextFile("./day5/input.txt")))
	fmt.Println(day5.SumIncorrectPageNumbers(day5.ReadTextFile("./day5/sample.txt")))
	fmt.Println(day5.SumIncorrectPageNumbers(day5.ReadTextFile("./day5/input.txt")))
}
