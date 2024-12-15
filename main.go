package main

import (
	"aoc2024/day13"
	"fmt"
)

func main() {
	machines := day13.ReadInput("./day13/input.txt")
	fmt.Println(day13.Combine2(machines))
}
