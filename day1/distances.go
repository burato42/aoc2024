package day1

import (
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func digitToInt(digitStr string) int {
	digit, err := strconv.Atoi(digitStr)
	if err != nil {
		panic("We have not a number, we can't proceed")
	}
	return digit
}

func ReadInput(textFile string) ([]int, []int) {
	file, err := os.Open(textFile)
	if err != nil {
		panic("Can't read from the file")
	}
	defer file.Close()

	var left, right []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		distances := strings.Fields(line)

		left = append(left, digitToInt(distances[0]))
		right = append(right, digitToInt(distances[1]))
	}

	if err := scanner.Err(); err != nil {
		panic("Scanner error, we can't proceed")
	}

	return left, right

}

func CalculateDist(left, right []int) int {
	sort.Ints(left)
	sort.Ints(right)
	res := 0
	for i, coord := range left {
		res += int(math.Abs(float64(coord - right[i])))
	}

	return res
}

func CalculateSimilarity(left, right []int) int {
	similarity := 0
	rightCounter := make(map[int]int)

	for _, coord := range right {
		rightCounter[coord] += 1
	}

	for _, coord := range left {
		similarity += coord * rightCounter[coord]
	}

	return similarity
}