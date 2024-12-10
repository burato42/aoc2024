package day10

import (
	"aoc2024/utils"
)


func Trail(matrix [][]int) int {
	var dfs func([][]int, int, int)
	res := [][2]int{}
	
	dfs = func (matrix [][]int, row int, col int) {

		if matrix[row][col] == 9 {
			res = append(res, [2]int{row, col})
			return
		}
		dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		for _, d := range dirs {
			if utils.IsInside(row + d[0], col + d[1], len(matrix), len(matrix[0])) && 
			matrix[row + d[0]][col + d[1]] - matrix[row][col] == 1 {
				dfs(matrix, row + d[0], col + d[1])
			}
		}
	}
	
	sum := 0
	for row, line := range matrix {
		for col, el := range line {
			if el == 0 {
				dfs(matrix, row, col)
				// res = utils.RemoveDuplicates(res) // uncomment for a task 1
				sum += len(res)
				res = [][2]int{}
			}
		}
	}

	return sum
}

