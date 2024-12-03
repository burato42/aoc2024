package day3

import (
	"aoc2024/utils"
	"regexp"
)

func SumMuls(input string) int {
	res := 0
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		res += utils.DigitToInt(match[1]) * utils.DigitToInt(match[2])
	}
	return res
}

func SumMulsWithAccuracy(input string) int {
	res := 0
	include := true
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|don't\(\)|do\(\)`)
	matches := re.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		if match[0] == "do()" {
			include = true
			continue
		} else if match[0] == "don't()" {
			include = false
			continue
		}

		if include {
			res += utils.DigitToInt(match[1]) * utils.DigitToInt(match[2])
		}

	}
	return res
}
