package main

import (
	"aoc2024/day6"
	"aoc2024/utils"
	"fmt"
)

func main() {
	fmt.Println(day6.CountLoops(utils.ReadTextToMatrix(("./day6/input.txt"))))
}
