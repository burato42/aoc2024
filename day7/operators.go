package day7

import (
	"aoc2024/utils"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadInput(textFile string) [][]int {
	file, err := os.Open(textFile)
	if err != nil {
		panic("Can't read from the file")
	}
	defer file.Close()

	res := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		numbers := []int{}
		for i, digit := range fields {
			if i == 0 {
				numbers = append(numbers, utils.DigitToInt(strings.TrimRight(digit, ":")))
			} else {
				numbers = append(numbers, utils.DigitToInt(digit))
			}
		}
		res = append(res, numbers)
	}
	return res
}

func CalcCalibration(numbers []int) int {
	var helper func(int, []int) bool
	finalRes := numbers[0]

	helper = func(res int, nums []int) bool {
		if res > finalRes {
			return false
		}
		if len(nums) == 0 && res != finalRes {
			return false
		}
		if res == finalRes && len(nums) == 0 {
			return true
		}

		if res == 0 {
			return helper(res+nums[0], nums[1:])
		}
		return helper(res+nums[0], nums[1:]) || helper(res*nums[0], nums[1:])
	}

	if helper(0, numbers[1:]) {
		return finalRes
	}
	return 0
}

func concat(a, b int) int {
	return utils.DigitToInt(strconv.Itoa(a) + strconv.Itoa(b))
}

func CalcExtraCalibration(numbers []int) int {
	var helper func(int, []int) bool
	finalRes := numbers[0]

	helper = func(res int, nums []int) bool {
		if res > finalRes {
			return false
		}
		if len(nums) == 0 && res != finalRes {
			return false
		}
		if res == finalRes && len(nums) == 0 {
			return true
		}

		if res == 0 {
			return helper(res+nums[0], nums[1:])
		}
		return helper(res+nums[0], nums[1:]) ||
			helper(res*nums[0], nums[1:]) ||
			helper(concat(res, nums[0]), nums[1:])
	}

	if helper(0, numbers[1:]) {
		return finalRes
	}
	return 0
}

func SumCalibrations(inputs [][]int) int {
	sum := 0
	for _, input := range inputs {
		sum += CalcCalibration(input)
	}
	return sum
}

func SumExtraCalibrations(inputs [][]int) int {
	sum := 0
	for _, input := range inputs {
		sum += CalcExtraCalibration(input)
	}
	return sum
}
