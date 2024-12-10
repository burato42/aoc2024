package main

import (
	"aoc2024/day10"
	"aoc2024/utils"
	"fmt"
)

func main() {
	input := utils.ReadTextToMatrix("./day10/input.txt")
	intInput := utils.StringToIntMatrix(input)
	fmt.Println(day10.Trail(intInput)) 
}
