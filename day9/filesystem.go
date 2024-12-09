package day9

import (
	"aoc2024/utils"
	"bufio"
	"os"
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
	left, right := 0, len(blocks) - 1
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
		if digit == "." {
			continue
		}
		checkSum += i * utils.DigitToInt(digit)
	}
	return checkSum
}

type Empty struct {
	start, size int
}
type Busy struct {
	start, size int
	value string
}

func FindSpareSpace(blocks []string) []Empty {
	spareAreas := []Empty{}
	start := -1
	size := 0
	for i, digit := range blocks {
		if digit == "." && start == -1 {
			start = i
		} else if  digit == "." && start != -1 {
			continue
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
	spareAreas := []Busy{}
	prev := ""
	start := -1
	for i, digit := range blocks {
		if digit != "." && prev == "" {
			start = i
			prev = digit
		} else if  digit != "." && prev == digit {
			continue
		} else if digit == "." && prev != "" {
			size := i - start
			spareAreas = append(spareAreas, Busy{start, size, blocks[i - 1]})
			start = -1
			prev = ""
		} else if digit == "." && prev == "" {
			continue
		} else if digit != "." && prev != digit {
			size := i - start
			spareAreas = append(spareAreas, Busy{start, size, blocks[i - 1]})
			start = i - 1
			prev = digit
		}
	}
	if start != -1 {
		spareAreas = append(spareAreas, Busy{start + 1, len(blocks) - start - 1, blocks[len(blocks) - 1]})
	}
	return spareAreas
}


func Mix(blocks []string, busy []Busy, empty []Empty) []string {
	for i := len(busy) - 1; i >= 0; i-- {
		for j := 0; j < len(empty) - 1; j++ {
			if busy[i].size <= empty[j].size {
				for k := 0; k < busy[i].size; k++ {
					blocks[empty[j].start] = busy[i].value
					empty[j].start++
					empty[j].size--
				}
				for k := 0; k < busy[i].size; k++ {
					blocks[busy[i].start + k] = "."
				}
				break
			}
		}
	}
	return blocks
}