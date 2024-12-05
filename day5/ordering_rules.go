package day5

import (
	"aoc2024/utils"
	"bufio"
	"os"
	"slices"
	"sort"
	"strings"
)

func ReadTextFile(textFile string) (map[int][]int, [][]int) {
	file, err := os.Open(textFile)
	if err != nil {
		panic("Can't read from the file")
	}
	defer file.Close()

	seqs := [][]int{}
	scanner := bufio.NewScanner(file)
	rules := make(map[int][]int)
	read_rules := true
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			read_rules = false
			continue
		}

		if read_rules {
			fromTo := strings.Split(line, "|")
			from := utils.DigitToInt(fromTo[0])
			to := utils.DigitToInt(fromTo[1])
			if res, ok := rules[from]; ok {
				rules[from] = append(res, to)
			} else {
				rules[from] = []int{to}
			}
		} else {
			order := []int{}
			for _, digit := range strings.Split(line, ",") {
				order = append(order, utils.DigitToInt(digit))
			}
			seqs = append(seqs, order)
		}
		
	}

	return rules, seqs
}

func IsCorrectOrder(rules map[int][]int, line []int) bool {
	size := len(line)
	for i := 0; i < size - 1; i++ {
		if to, ok := rules[line[i]]; !ok || !slices.Contains(to, line[i+1]) {
			return false
		}
	}
	return true
}

func SumPageNumbers(rules map[int][]int, lines [][]int) int {
	res := 0
	for _, line := range lines {
		if IsCorrectOrder(rules, line) {
			res += line[len(line)/2]
		}
	}
	return res
}

func FixOrder(rules map[int][]int, line []int) []int {
	for !IsCorrectOrder(rules, line) {
		for i := 0; i < len(line) - 1; i++ {
			if to, ok := rules[line[i]]; !ok || !slices.Contains(to, line[i+1]) {
				line[i], line[i+1] = line[i+1], line[i]
			}
		}
	}
	return line
}

func SumIncorrectPageNumbers(rules map[int][]int, lines [][]int) int {
	res := 0
	for _, line := range lines {
		if !IsCorrectOrder(rules, line) {
			// newLine := FixOrder(rules, line)
			newLine := Sort(rules, line)
			res += newLine[len(line)/2]
		}
	}
	return res
}

func Sort(rules map[int][]int, line []int) []int {
	sort.SliceStable(line, func(i, j int) bool {
		if to, ok := rules[line[i]]; ok && slices.Contains(to, line[j]) {
			return true
		}
		return false
	})
	return line
}