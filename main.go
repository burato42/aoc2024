package main

import (
	"fmt"
	"aoc2024/day1"
)


func main() {
	fmt.Println(day1.CalculateDist(day1.ReadInput("./day1/sample.txt")))
	fmt.Println(day1.CalculateDist(day1.ReadInput("./day1/input.txt")))
	fmt.Println(day1.CalculateSimilarity((day1.ReadInput("./day1/sample.txt"))))
	fmt.Println(day1.CalculateSimilarity((day1.ReadInput("./day1/input.txt"))))
}