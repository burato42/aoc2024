package main

import (
	"aoc2024/day3"
	"aoc2024/utils"
	"fmt"
)

func main() {
	fmt.Println(day3.SumMuls("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"))
	fmt.Println(day3.SumMulsWithAccuracy("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"))
	fmt.Println(day3.SumMuls(utils.ReadTextFile("./day3/input.txt")))
	fmt.Println(day3.SumMulsWithAccuracy(utils.ReadTextFile("./day3/input.txt")))

}
