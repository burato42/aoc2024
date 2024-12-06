package day6

import (
	"aoc2024/utils"
	"strings"
)


type Coordinate struct {
	v, h, dirV, dirH int
}


func TurnRight(dirV, dirH int) (int, int) {
	switch {
	case dirV == 1 && dirH == 0:
		return 0, -1
	case dirV == 0 && dirH == -1:
		return -1, 0
	case dirV == -1 && dirH == 0:
		return 0, 1 
	case dirV == 0 && dirH == 1:
		return 1, 0
	default:
		panic("this direction doesn't exist") 
	}
}


func FindGuardAndDir(matrix [][]string) (int, int, int, int) {
	height := len(matrix)
	width := len(matrix[0])

	for v := 0; v < height; v++ {
		for h := 0; h < width; h++ {
			switch matrix[v][h] {
			case "v":
				return v, h, 1, 0
			case ">":
				return v, h, 0, 1
			case "<":
				return v, h, 0, -1
			case "^":
				return v, h, -1, 0
			}
		}
	}

	return 0, 0, 0, 0
}

func CountSteps(matrix [][]string) int {
	v, h, dirV, dirH := FindGuardAndDir(matrix)

	steps := 0
	for {
		if matrix[v][h] != "X" {
			matrix[v][h] = "X"
			steps++
		}
		if utils.IsInside(v + dirV, h + dirH, len(matrix), len(matrix[0])) && matrix[v + dirV][h + dirH] != "#" {
			v += dirV
			h += dirH
		} else if utils.IsInside(v + dirV, h + dirH, len(matrix), len(matrix[0])) && matrix[v + dirV][h + dirH] == "#" {
			dirV, dirH = TurnRight(dirV, dirH)
		} else if !utils.IsInside(v + dirV, h + dirH, len(matrix), len(matrix[0])) {
			return steps
		}
	}
}

func HasLoop(matrix [][]string) bool {
	v, h, dirV, dirH := FindGuardAndDir(matrix)
	visited := make(map[Coordinate]bool)

	for {
		if _, ok := visited[Coordinate{v, h, dirV, dirH}]; ok {
			return true
		} else {
			visited[Coordinate{v, h, dirV, dirH}] = true
		}

		if utils.IsInside(v + dirV, h + dirH, len(matrix), len(matrix[0])) && matrix[v + dirV][h + dirH] != "#" {
			v += dirV
			h += dirH
		} else if utils.IsInside(v + dirV, h + dirH, len(matrix), len(matrix[0])) && matrix[v + dirV][h + dirH] == "#" {
			dirV, dirH = TurnRight(dirV, dirH)
		} else if !utils.IsInside(v + dirV, h + dirH, len(matrix), len(matrix[0])) {
			return false
		}
	}
}

func AddObstacles(matrix [][]string) [][][]string {
    height := len(matrix)
    width := len(matrix[0])
    matrices := [][][]string{}

    for v := 0; v < height; v++ {
        for h := 0; h < width; h++ {
            if !strings.Contains("#^><v", matrix[v][h]) {
                matrix[v][h] = "#"

                // Create a deep copy of the matrix
                newMatrix := make([][]string, height)
                for i := 0; i < height; i++ {
                    newMatrix[i] = make([]string, width)
                    copy(newMatrix[i], matrix[i])
                }

                matrices = append(matrices, newMatrix)
                matrix[v][h] = "."
            }
        }
    }

    return matrices

}

func CountLoops(matrix [][]string) int {
	counter := 0
	for _, mtrx := range AddObstacles(matrix) {
		if HasLoop(mtrx) {
			counter++
		}
	}
	return counter
}