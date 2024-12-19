package day16

// This approach doesn't work for me :(

import (
	"math"
)


func FindStartAndEnd(maze [][]string) map[string][2]int {
	res := make(map[string][2]int)
	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[0]); j++ {
			if maze[i][j] == "S" {
				res["S"] = [2]int{i, j}
			}
			if maze[i][j] == "E" {
				res["E"] = [2]int{i, j}
			}
		}
	}
	return res
}

func getHeuristic(start, goal [2]int) int {
	return int(math.Pow(float64(start[0] - goal[0]), 2) + math.Pow(float64(start[1] - goal[1]), 2))
}

type Node struct {
	g, h, f int
	dir int
	parent *Node
	position [2]int
}


func reversed[T any](s []T) []T {
    reversed := make([]T, len(s))
    for i, j := 0, len(s)-1; j >= 0; i, j = i+1, j-1 {
        reversed[i] = s[j]
    }
    return reversed
}


func GetPathAStar(maze [][]string, startEnd map[string][2]int) [][2]int {
	start := &Node{g: 0, h: 0, f: 0, position: startEnd["S"]}
	goal := &Node{position: startEnd["E"]}

	var path [][2]int

	openList := []*Node{start}
	closedList := []*Node{}

	for len(openList) > 0 {
		curNode := openList[0]
		curInd := 0
		for ind, item := range openList {
			if item.f < curNode.f {
				curNode = item
				curInd = ind
			}
		}

		openList = append(openList[:curInd], openList[curInd+1:]...)
		closedList = append(closedList, curNode)

		if curNode.position == goal.position {
			path = [][2]int{}
			cur := curNode
			for cur != nil {
				path = append(path, cur.position)
				cur = cur.parent
			}
			return reversed(path)
		}
		
		children := []Node{}
		for _, newPos := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			nodePos := [2]int{curNode.position[0] + newPos[0], curNode.position[1] + newPos[1]}
			if maze[nodePos[0]][nodePos[1]] == "#" {
				continue
			}
			var newDir int
			if curNode.dir == 0 && newPos[1] == 0 {
				newDir = 1
			} else if curNode.dir == 0 && newPos[0] == 0 {
				newDir = 0
			} else if curNode.dir == 1 && newPos[0] == 0 {
				newDir = 1
			} else {
				newDir = 0
			}

			newNode := Node{parent: curNode, position: nodePos, dir: newDir}
				
			children = append(children, newNode)
		}

		for _, child := range children {
			for _, closedChild := range closedList {
				if closedChild.position == child.position {
					continue
				}
			}

			if child.dir == curNode.dir {
				child.h = getHeuristic(child.position, goal.position)
			} else {
				child.h = getHeuristic(child.position, goal.position) + 1000
			}
			child.g = curNode.g + 1
			// child.h = getHeuristic(child.position, goal.position)
			child.f = child.g + child.h

			for _, openNode := range openList {
				if child.position == openNode.position && child.g > openNode.g {
					continue
				}
			}
			openList = append(openList, &child)
		}
	}
	return path
}

func CalculateScore(path [][2]int) int {
	res := 0
	prev := path[0]
	var dir int
	for i, pos := range path[1:] {
		if i == 0 {
			if pos[1] == prev[1] {
				dir = 1
			} else {
				dir = 0
			}
			prev = pos
			res++
		} else if dir == 1 {
			if pos[1] == prev[1] {
				dir = 1
				res++
			} else {
				dir = 0
				res += 1001
			}
		} else {
			if pos[0] == prev[0] {
				dir = 0
				res++
			} else {
				dir = 1
				res += 1001
			}
		}
	}
	return res
}