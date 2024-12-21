package main

import (
	"aoc2024/day17"
)

func main() {
	registers, program := day17.ReadInput("./day17/input.txt")
	day17.FixCorruption(registers, program)

}
