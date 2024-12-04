package main

import (
	"aoc2024/day4"
	"fmt"
)

func main() {
	// fmt.Println(day4.FindWords(day4.ReadTextToMatrix("./day4/sample.txt")))
	// fmt.Println(day4.FindWords(day4.ReadTextToMatrix("./day4/input.txt")))
	fmt.Println(day4.FindXWords(day4.ReadTextToMatrix("./day4/sample.txt")))
	fmt.Println(day4.FindXWords(day4.ReadTextToMatrix("./day4/input.txt")))
}
