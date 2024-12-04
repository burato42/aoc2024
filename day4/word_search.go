package day4

import (
	"bufio"
	"os"
	"strings"
)

func ReadTextToMatrix(textFile string) [][]string {
	file, err := os.Open(textFile)
	if err != nil {
		panic("Can't read from the file")
	}
	defer file.Close()

	content := [][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		content = append(content, strings.Split(line, ""))
	}

	return content
}

func isInside(v, h, height, width int) bool {
	return v >= 0 && h >= 0 && v < height && h < width
}

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
					if !isInside(v+i*d[0], h+i*d[1], height, width) || target[i] != matrix[v+i*d[0]][h+i*d[1]] {
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

	height := len(matrix)
	width := len(matrix[0])
	counter := 0
	for v, line := range matrix {
		for h := range line {
			if matrix[v][h] == "A" &&
				isInside(v-1, h-1, height, width) &&
				isInside(v-1, h+1, height, width) &&
				isInside(v+1, h-1, height, width) &&
				isInside(v+1, h+1, height, width) &&
				((matrix[v-1][h-1] == matrix[v-1][h+1] && matrix[v-1][h+1] == "S" &&
					matrix[v+1][h-1] == matrix[v+1][h+1] && matrix[v+1][h+1] == "M") ||
					(matrix[v-1][h-1] == matrix[v-1][h+1] && matrix[v-1][h+1] == "M" &&
						matrix[v+1][h-1] == matrix[v+1][h+1] && matrix[v+1][h+1] == "S") ||
					(matrix[v-1][h-1] == matrix[v+1][h-1] && matrix[v-1][h-1] == "S" &&
						matrix[v+1][h+1] == matrix[v-1][h+1] && matrix[v-1][h+1] == "M") ||
					(matrix[v-1][h-1] == matrix[v+1][h-1] && matrix[v-1][h-1] == "M" &&
						matrix[v+1][h+1] == matrix[v-1][h+1] && matrix[v-1][h+1] == "S")) {
				counter++
			}

		}
	}
	return counter
}
