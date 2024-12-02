package day2

import (
	"aoc2024/utils"
	"bufio"
	"math"
	"os"
	"strings"
)

func isSafe(levels []int) bool {
	if len(levels) < 2 {
		return false
	}

	sign := 0
	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]

		if math.Abs(float64(diff)) == 0 || math.Abs(float64(diff)) > 3 {
			return false
		}

		if sign != 0 && diff*sign < 0 {
			return false
		}

		if diff == 0 {
			return false
		} else if diff > 0 {
			sign = 1
		} else {
			sign = -1
		}
	}
	return true
}

func isSafeWithTolerance(levels []int) bool {
	for i := 0; i < len(levels); i++ {
		excluded := make([]int, len(levels)-1)
		copy(excluded[:i], levels[:i])
		copy(excluded[i:], levels[i+1:])
		if isSafe(excluded) {
			return true
		}
	}
	return false
}

func AnalyzeReport(textFile string) (int, int) {
	file, err := os.Open(textFile)
	if err != nil {
		panic("Can't read from the file")
	}
	defer file.Close()

	var counter int
	var counterWithErr int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		levels := []int{}
		reports := strings.Fields(line)
		for _, digit := range reports {
			levels = append(levels, utils.DigitToInt(digit))
		}

		if isSafe(levels) {
			counter++
			counterWithErr++
		} else if isSafeWithTolerance(levels) {
			counterWithErr++
		}
	}

	if err := scanner.Err(); err != nil {
		panic("Scanner error, we can't proceed")
	}

	return counter, counterWithErr

}
