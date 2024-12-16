package day14

import (
	"aoc2024/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"time"
)

type Robot struct {
	x, y, velX, velY int
}

type Bathroom struct {
	height, width int
	robots []Robot
}


func ReadInput(textFile string) []Robot {
	file, err := os.Open(textFile)
	if err != nil {
		panic("Can't read from the file")
	}
	defer file.Close()

	robots := []Robot{}
	regex := regexp.MustCompile(`p=(\d+),(\d+) v=(\-?\d+),(\-?\d+)`)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		data := regex.FindAllStringSubmatch(line, -1)
		
		newRobot := &Robot{
            x: utils.DigitToInt(data[0][1]),
            y: utils.DigitToInt(data[0][2]),
            velX: utils.DigitToInt(data[0][3]),
            velY: utils.DigitToInt(data[0][4]),
        }
        robots = append(robots, *newRobot) 

	}
	return robots
}

func modulo(a, b int) int {
	return (a%b + int(math.Abs(float64(b)))) % int(math.Abs(float64(b)))
}

func (robot *Robot) Move(br Bathroom) {
	robot.x = modulo(robot.x + robot.velX, br.width)
	robot.y = modulo(robot.y + robot.velY, br.height)
}

func (br *Bathroom) GetQuadrantCount() [4]int {
	var q1, q2, q3, q4 int
	
	for _, robot := range br.robots {
		// As as the measuremens of the bathroom are odd, we skip the check for this middle line
		if robot.x < br.width/2 {
			if robot.y < br.height/2 {
				q1++
			} else if robot.y > br.height/2 {
				q2++
			}
		} else if robot.x > br.width/2 {
			if robot.y < br.height/2 {
				q3++
			} else if robot.y > br.height/2 {
				q4++
			}
		}
		
	}
	return [4]int{q1, q2, q3, q4}
}

func (br *Bathroom) GetSafetyFactor(qs [4]int) int {
	return qs[0]*qs[1]*qs[2]*qs[3]
}

func clearScreen() {
	switch runtime.GOOS {
	case "linux", "darwin":
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
	case "windows":
			cmd := exec.Command("cmd", "/c", "cls")
			cmd.Stdout = os.Stdout
			cmd.Run()
	default:
			fmt.Println("Unsupported platform. Cannot clear screen.")
	}
}

func SimulateWithGraphics(steps, height, width int, textFile string) int {
	br := Bathroom{height, width, ReadInput(textFile)}
	matrix := make([][]string, height)
	for row := range matrix {
		matrix[row] = make([]string, width)
	}
	
	for j := 0; j < steps; j++ {
		for i := 0; i < len(br.robots); i++ {
			br.robots[i].Move(br)
		}
		
		for k := 0; k < height; k++ {
			for l := 0; l < width; l++ {
				matrix[k][l] = "."
			}
		}
		for _, robot := range br.robots {
			matrix[robot.y][robot.x] = "*"	
		}

		// clearScreen() // Clear the previous frame
		fmt.Printf("Step %v\n", j)
		for _, line := range matrix {
			fmt.Println(line)
		}
		time.Sleep(200 * time.Millisecond)

	}

	return br.GetSafetyFactor(br.GetQuadrantCount())
}

func Simulate(steps, height, width int, textFile string) int {
	br := Bathroom{height, width, ReadInput(textFile)}
	matrix := make([][]string, height)
	for row := range matrix {
		matrix[row] = make([]string, width)
	}
	
	for j := 0; j < steps; j++ {
		for i := 0; i < len(br.robots); i++ {
			br.robots[i].Move(br)
		}
		
		for k := 0; k < height; k++ {
			for l := 0; l < width; l++ {
				matrix[k][l] = "."
			}
		}
		for _, robot := range br.robots {
			matrix[robot.y][robot.x] = "*"	
		}

		for k := 0; k < height; k++ {
			for l := 0; l < width; l++ {
				
				if k < height - 21 && matrix[k][l] == "*" {
					found := true
					r := 0
					for r < 20 {
						if matrix[k+r][l] != "*" {
							found = false
							break
						}
						r++
					}
					if found {
						fmt.Printf("Step %v\n", j)
						for _, line := range matrix {
							fmt.Println(line)
						}
					}
				}
			}
		}

	}
	return 0
}