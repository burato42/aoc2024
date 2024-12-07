package main

import (
	"aoc2024/day7"
	"fmt"
)

func main() {
	fmt.Println(day7.ReadInput("./day7/sample.txt"))
	fmt.Println(day7.CalcCalibration([]int{83, 17, 5}))
	fmt.Println(day7.SumCalibrations(day7.ReadInput("./day7/sample.txt")))
	fmt.Println(day7.SumExtraCalibrations(day7.ReadInput("./day7/sample.txt")))
	fmt.Println(day7.SumCalibrations(day7.ReadInput("./day7/input.txt")))
	fmt.Println(day7.SumExtraCalibrations(day7.ReadInput("./day7/input.txt")))
}
