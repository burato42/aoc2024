package main

import (
	"aoc2024/day16"
	"aoc2024/utils"
	"fmt"
)

func main() {
	maze := utils.ReadTextToMatrix("./day16/sample1.txt")
	for _, line := range maze {
		fmt.Println(line)
	}
	startEnd := day16.FindStartAndEnd(maze)
	path := day16.GetPathAStar(maze, startEnd)
	for _, point := range path {
		maze[point[0]][point[1]] = "*"
	}
	for _, line := range maze {
		fmt.Println(line)
	}
	fmt.Println(day16.CalculateScore(path))

}
