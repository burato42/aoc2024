package day11

import (
	"aoc2024/utils"
	"strconv"
)

func GetStonesAfterBlink(input []int) []int {
	res := []int{}
	for _, stone := range input {
		switch {
		case stone == 0:
			res = append(res, 1)
		case len(strconv.Itoa(stone))%2 == 0:
			size := len(strconv.Itoa(stone))
			left := strconv.Itoa(stone)[:size/2]
			right := strconv.Itoa(stone)[size/2:]
			res = append(res, utils.DigitToInt(left), utils.DigitToInt(right))
		default:
			res = append(res, stone*2024)
		}
	}
	return res
}

var memo = make(map[string]int)

func CountStonesAfterNBlinks() int {
	var helper func(int, int) int

	helper = func(iteration int, stone int) int {
		key := strconv.Itoa(stone) + "-" + strconv.Itoa(iteration)
		if val, ok := memo[key]; ok {
			return val
		}

		if iteration == 75 {
			return 1
		}

		switch {
		case stone == 0:
			memo[key] = helper(iteration+1, 1)
			return memo[key]
		case len(strconv.Itoa(stone))%2 == 0:
			size := len(strconv.Itoa(stone))
			left := strconv.Itoa(stone)[:size/2]
			right := strconv.Itoa(stone)[size/2:]
			memo[key] = helper(iteration+1, utils.DigitToInt(left)) + helper(iteration+1, utils.DigitToInt(right))
			return memo[key]
		default:
			memo[key] = helper(iteration+1, stone*2024)
			return memo[key]
		}
	}

	sum := 0
	input := []int{3935565, 31753, 437818, 7697, 5, 38, 0, 123}
	for i := 0; i < len(input); i++ {
		sum += helper(0, input[i])
	}
	return sum

}
