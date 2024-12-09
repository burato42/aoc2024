package main

import (
	"aoc2024/day9"
	"fmt"
)

func main() {
	input := day9.ReadInput("./day9/input.txt")
	// input := day9.ReadInput("./day9/sample.txt")
	filledBlocks := day9.FillBlocks(input)
	// fmt.Println(filledBlocks)
	// compressed := day9.CompressBlocks(filledBlocks)
	// fmt.Println(day9.CalcCheckSum(compressed))
	spare := day9.FindSpareSpace(filledBlocks)
	busy := day9.FindBusySpace(filledBlocks)
	// fmt.Println(busy)
	compr := day9.Mix(filledBlocks, busy, spare)
	fmt.Println(day9.CalcCheckSum(compr))
}
