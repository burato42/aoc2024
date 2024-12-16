package day14

import (
	"aoc2024/utils"
	"bufio"
	"go/scanner"
	"math"
	"os"
	"regexp"
	"utils"
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
		robots = append(robots, Robot{
			x: utils.DigitToInt(data[0][1]),
			y: utils.DigitToInt(data[0][2]),
			velX: utils.DigitToInt(data[0][3]),
			velY: utils.DigitToInt(data[0][4]),
		})

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

func (robot *Robot) GetLocation() [2]int {
	return [...]int{robot.y, robot.x}
}

func (br *Bathroom) GetSafetyFactors(robots []Robot)


func Simulate(input string, steps int, br Bathroom) {
	robots := ReadInput(input)
	for _, robot := range robots {
		for i := 0; i < steps; i++ {
			robot.Move(br)
		}
		
	}
}