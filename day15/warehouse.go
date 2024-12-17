package day15

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readInputs(textFile string) ([][]string, []string) {
	file, err := os.Open(textFile)
	if err != nil {
		panic("Can't read from the file")
	}
	defer file.Close()

	warehouse := [][]string{}
	directions := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.Contains(line, "<") || strings.Contains(line, ">") || strings.Contains(line, "^") || strings.Contains(line, "v") {
			directions = append(directions, strings.Split(line, "")...)
		} else {
			warehouse = append(warehouse, strings.Split(line, ""))
		}

	}

	return warehouse, directions
}

type Robot struct {
	y, x, dirY, dirX int
}

type Warehouse struct {
	layout [][]string
	robot  Robot
}

func (robot *Robot) changeDir(dir string) {
	switch dir {
	case "<":
		robot.dirX = -1
		robot.dirY = 0
	case "^":
		robot.dirX = 0
		robot.dirY = -1
	case ">":
		robot.dirX = 1
		robot.dirY = 0
	case "v":
		robot.dirX = 0
		robot.dirY = 1
	}
}

func (wh *Warehouse) findRobot() (int, int) {
	for y, line := range wh.layout {
		for x, el := range line {
			if el == "@" {
				return y, x
			}
		}
	}
	return -1, -1
}

func (wh *Warehouse) moveRobot() {
	if wh.layout[wh.robot.y+wh.robot.dirY][wh.robot.x+wh.robot.dirX] == "#" {
		return
	}

	if wh.layout[wh.robot.y+wh.robot.dirY][wh.robot.x+wh.robot.dirX] == "." {
		wh.layout[wh.robot.y][wh.robot.x] = "."
		wh.robot.y += wh.robot.dirY
		wh.robot.x += wh.robot.dirX
		wh.layout[wh.robot.y][wh.robot.x] = "@"

	} else {
		boxes := 0
		curY, curX := wh.robot.y+wh.robot.dirY, wh.robot.x+wh.robot.dirX
		for wh.layout[curY][curX] != "#" {
			if wh.layout[curY][curX] == "O" {
				boxes++
				curX += wh.robot.dirX
				curY += wh.robot.dirY
			} else if wh.layout[curY][curX] == "." {
				for boxes > 0 {
					wh.layout[curY][curX] = "O"
					boxes--
					curX -= wh.robot.dirX
					curY -= wh.robot.dirY
				}
				wh.layout[wh.robot.y][wh.robot.x] = "."
				wh.robot.y += wh.robot.dirY
				wh.robot.x += wh.robot.dirX
				wh.layout[wh.robot.y][wh.robot.x] = "@"
				break
			}
		}
	}

}

func (wh *Warehouse) Show() {
	for _, line := range wh.layout {
		fmt.Println(line)
	}
}

func (wh *Warehouse) SumGPS(box string) int {
	res := 0
	for y, line := range wh.layout {
		for x, el := range line {
			if el == box {
				res += y*100 + x
			}
		}
	}
	return res
}

func Simulate(input, box string) {
	layout, dirs := readInputs(input)
	wh := Warehouse{layout: layout, robot: Robot{}}
	y, x := wh.findRobot()
	wh.robot.y = y
	wh.robot.x = x
	wh.Show()
	for _, dir := range dirs {
		wh.robot.changeDir(dir)
		wh.moveRobot()
	}
	wh.Show()
	fmt.Println(wh.SumGPS(box))
}
