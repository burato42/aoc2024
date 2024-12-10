package day9

import (
	"aoc2024/utils"
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

func ReadInput(textFile string) []int {
	file, err := os.Open(textFile)
	if err != nil {
		panic("Can't read from the file")
	}
	defer file.Close()

	res := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		digits := strings.Split(line, "")
		numbers := []int{}
		for _, digit := range digits {
			numbers = append(numbers, utils.DigitToInt(digit))
		}
		res = append(res, numbers...)
	}
	return res
}

func FillBlocks(blocks []int) []string {
	filled := []string{}

	id := 0
	for {
		for i := 0; i < len(blocks); i++ {
			if i%2 == 0 {
				for j := 0; j < blocks[i]; j++ {
					filled = append(filled, strconv.Itoa(id))
				}
				id++
			} else {
				for j := 0; j < blocks[i]; j++ {
					filled = append(filled, ".")
				}
			}
		}
		if id >= len(blocks)/2 {
			break
		}

	}
	return filled
}

func CompressBlocks(blocks []string) []string {
	left, right := 0, len(blocks)-1
	for left < right {
		if blocks[left] != "." {
			left++
		} else if blocks[right] == "." {
			right--
		} else {
			blocks[left], blocks[right] = blocks[right], blocks[left]
		}
	}
	return blocks
}

func CalcCheckSum(blocks []string) int {
	checkSum := 0
	for i, digit := range blocks {
		if digit != "." {
			checkSum += i * utils.DigitToInt(digit)
		}
	}
	return checkSum
}

type Empty struct {
	start, size int
}
type Busy struct {
	start, size int
	value       string
}

func FindSpareSpace(blocks []string) []Empty {
	spareAreas := []Empty{}
	start := -1
	size := 0
	for i, digit := range blocks {
		if digit == "." && start == -1 {
			start = i
		} else if digit != "." && start != -1 {
			size = i - start
			spareAreas = append(spareAreas, Empty{start, size})
			start = -1
		}
	}
	if start != -1 {
		spareAreas = append(spareAreas, Empty{start, len(blocks) - start})
	}
	return spareAreas
}

func FindBusySpace(blocks []string) []Busy {
	fileAreas := []Busy{}
	prev := "."
	start := 0
	for i, digit := range blocks {
		if digit != prev && prev == "." {
			start = i
			prev = digit
		} else if digit != prev && digit != "." {
			fileAreas = append(fileAreas, Busy{start, i - start, prev})
			start = i
			prev = digit
		} else if digit != prev && digit == "." {
			fileAreas = append(fileAreas, Busy{start, i - start, prev})
			prev = "."
		}
	}
	if prev != "." {
		fileAreas = append(fileAreas, Busy{start, len(blocks) - start, prev})
	}

	return fileAreas
}

func Mix(blocks []string, busy []Busy, empty []Empty) []string {
	for i := len(busy) - 1; i >= 0; i-- {
		for j := 0; ; {
			if busy[i].size <= empty[j].size && empty[j].start+empty[j].size-1 < busy[i].start {
				for k := 0; k < busy[i].size; k++ {
					blocks[empty[j].start] = busy[i].value
					blocks[busy[i].start+k] = "."
					empty[j].start++
					empty[j].size--
				}
				if empty[j].size == 0 {
					empty = slices.Delete(empty, j, j+1)
				}
				break
			}
			j++
			if j >= len(empty) {
				break
			}
		}
	}
	return blocks
}
