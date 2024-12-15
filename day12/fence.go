package day12

import (
	"aoc2024/utils"
	"strings"

	"golang.org/x/text/width"
)

func CalcArea(matrix [][]string) map[string]int {
	area := 0
	perimeter := 0
	for i, row := range matrix {
		for j, elem := range row {
			if strings.Contains(elem, ".") {
				
			}

		}
	}
	return area
}

func calcPerims(matrix [][]string, v int, h int, sym string) int {
	hight := len(matrix)
	width := len(matrix[0])
	
	switch {
	case !utils.IsInside(v-1, h-1, hight, width) || 
		!utils.IsInside(v+1, h-1, hight, width) || 
		!utils.IsInside(v - 1, h+1, hight, width) || 
		!utils.IsInside(v+1, h+1, hight, width):
		return 2
	case !utils.IsInside(v+1, h, hight, width) ||
		!utils.IsInside(v-1, h, hight, width) ||
		!utils.IsInside(v, h-1, height, width) ||
		!utils.IsInside(v, h+1, height, width):
		return 1
	case sym != matrix[v+1][h] && sym != matrix[v-1][h]

}

func bfs(matrix [][]string, v int, h int, sym string) int {
	area := 0
	perimeter := 0
	toVisit := [][2]int{{v, h}}
	for len(toVisit) > 0 {
		next := toVisit[0]
		toVisit = toVisit[1:]
		area += 1
		if !utils.IsInside(v - 1, h, len(matrix), len(matrix[0])) 
	}
}