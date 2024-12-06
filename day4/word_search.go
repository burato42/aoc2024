package day4

import (
	"aoc2024/utils"
	"reflect"
)


func FindWords(matrix [][]string) int {
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, -1}, {-1, 1}}
	target := []string{"X", "M", "A", "S"}

	height := len(matrix)
	width := len(matrix[0])
	violates := false
	counter := 0
	for v, line := range matrix {
		for h := range line {
			for _, d := range directions {
				violates = false
				for i := 0; i < len(target); i++ {
					if !utils.IsInside(v+i*d[0], h+i*d[1], height, width) || target[i] != matrix[v+i*d[0]][h+i*d[1]] {
						violates = true
						break
					}
				}
				if !violates {
					counter++
				}
			}
		}
	}
	return counter
}

func FindXWords(matrix [][]string) int {
	patterns := [][]string{{"S", "S", "M", "M"}, {"S", "M", "S", "M"}, {"M", "M", "S", "S"}, {"M", "S", "M", "S"}}
	height := len(matrix)
	width := len(matrix[0])
	counter := 0
	for v, line := range matrix {
		for h := range line {
			options := []string{}
			if matrix[v][h] == "A" &&
				utils.IsInside(v-1, h-1, height, width) &&
				utils.IsInside(v-1, h+1, height, width) &&
				utils.IsInside(v+1, h-1, height, width) &&
				utils.IsInside(v+1, h+1, height, width) {
					options = append(options, matrix[v-1][h-1], matrix[v-1][h+1], matrix[v+1][h-1], matrix[v+1][h+1])
					for _, pattern := range patterns {
						if reflect.DeepEqual(pattern, options) {
							counter++
							break
						}
					}
			}

		}
	}
	return counter
}
