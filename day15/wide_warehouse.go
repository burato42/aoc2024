package day15

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readWideInputs(textFile string) ([][]string, []string) {
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
			level := []string{}
			for _, el := range strings.Split(line, "") {
				if el == "@" {
					level = append(level, el, ".")
				} else if el == "O" {
					level = append(level, "[", "]")
				} else {
					level = append(level, el, el)
				}

			}
			warehouse = append(warehouse, level)
		}

	}

	return warehouse, directions
}

func (wh *Warehouse) moveRobotWide() {
	if wh.layout[wh.robot.y+wh.robot.dirY][wh.robot.x+wh.robot.dirX] == "#" {
		return
	}

	if wh.layout[wh.robot.y+wh.robot.dirY][wh.robot.x+wh.robot.dirX] == "." {
		wh.layout[wh.robot.y][wh.robot.x] = "."
		wh.robot.y += wh.robot.dirY
		wh.robot.x += wh.robot.dirX
		wh.layout[wh.robot.y][wh.robot.x] = "@"

	} else {

		if wh.robot.dirY == 0 {
			boxes := []string{}
			curY, curX := wh.robot.y+wh.robot.dirY, wh.robot.x+wh.robot.dirX
			for wh.layout[curY][curX] != "#" {
				if wh.layout[curY][curX] == "[" || wh.layout[curY][curX] == "]" {
					boxes = append(boxes, wh.layout[curY][curX])
					curX += wh.robot.dirX
					curY += wh.robot.dirY
				} else if wh.layout[curY][curX] == "." {
					for len(boxes) > 0 {
						index := len(boxes) - 1
						item := boxes[index]
						boxes = boxes[:index]
						wh.layout[curY][curX] = item
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
		} else {
			boxes := [][2]int{}
			curY, curX := wh.robot.y+wh.robot.dirY, wh.robot.x+wh.robot.dirX
			var toExplore [][2]int

			if wh.layout[curY][curX] == "[" || wh.layout[curY][curX] == "]" {
				if wh.layout[curY][curX] == "[" {
					toExplore = [][2]int{{curY, curX}}
				} else {
					toExplore = [][2]int{{curY, curX - 1}}
				}

				for len(toExplore) > 0 {
					curY, curX = toExplore[0][0], toExplore[0][1]
					toExplore = toExplore[1:]

					boxes = append(boxes, [2]int{curY, curX})
					if wh.layout[curY+wh.robot.dirY][curX] == "]" {
						toExplore = append(toExplore, [2]int{curY + wh.robot.dirY, curX - 1})
					}
					if wh.layout[curY+wh.robot.dirY][curX+1] == "[" {
						toExplore = append(toExplore, [2]int{curY + wh.robot.dirY, curX + 1})
					}
					if wh.layout[curY+wh.robot.dirY][curX] == "[" {
						toExplore = append(toExplore, [2]int{curY + wh.robot.dirY, curX})
					}
					if wh.layout[curY+wh.robot.dirY][curX] == "#" || wh.layout[curY+wh.robot.dirY][curX+1] == "#" {
						boxes = [][2]int{}
						break
					}
				}
			}

			for len(boxes) > 0 {
				index := len(boxes) - 1
				item := boxes[index]
				boxes = boxes[:index]
				wh.layout[item[0]+wh.robot.dirY][item[1]] = "["
				wh.layout[item[0]+wh.robot.dirY][item[1]+1] = "]"
				wh.layout[item[0]][item[1]] = "."
				wh.layout[item[0]][item[1]+1] = "."
			}
			if wh.layout[wh.robot.y+wh.robot.dirY][wh.robot.x] == "." {
				wh.layout[wh.robot.y][wh.robot.x] = "."
				wh.robot.y += wh.robot.dirY
				wh.robot.x += wh.robot.dirX
				wh.layout[wh.robot.y][wh.robot.x] = "@"
			}

		}

	}
}

func SimulateWide(input, box string) {
	layout, dirs := readWideInputs(input)
	wh := Warehouse{layout: layout, robot: Robot{}}
	y, x := wh.findRobot()
	wh.robot.y = y
	wh.robot.x = x
	wh.Show()
	for _, dir := range dirs {
		wh.robot.changeDir(dir)
		wh.moveRobotWide()
	}
	wh.Show()
	fmt.Println(wh.SumGPS(box))
}
