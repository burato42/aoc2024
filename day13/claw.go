package day13

import (
	"aoc2024/utils"
	"bufio"
	// "fmt"
	"os"
	"regexp"
)

type Move struct {
	x, y int
}

type Prize struct {
	x, y int
}

type Machine struct {
	a Move
	b Move
	p Prize
}

func ReadInput(textFile string) []Machine {
	file, err := os.Open(textFile)
	if err != nil {
		panic("Can't read from the file")
	}
	defer file.Close()

	machines := []Machine{}
	regexA := regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)`)
	regexB := regexp.MustCompile(`Button B: X\+(\d+), Y\+(\d+)`)
	regexPrize := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)
	scanner := bufio.NewScanner(file)
	var moveA Move
	var moveB Move
	var prize Prize

	for scanner.Scan() {
		line := scanner.Text()
		if regexA.MatchString(line) {
			res := regexA.FindAllStringSubmatch(line, -1)
			moveA = Move{utils.DigitToInt(res[0][1]), utils.DigitToInt(res[0][2])}
		}
		if regexB.MatchString(line) {
			res := regexB.FindAllStringSubmatch(line, -1)
			moveB = Move{utils.DigitToInt(res[0][1]), utils.DigitToInt(res[0][2])}
		}
		if regexPrize.MatchString(line) {
			res := regexPrize.FindAllStringSubmatch(line, -1)
			prize = Prize{utils.DigitToInt(res[0][1]), utils.DigitToInt(res[0][2])}
			machines = append(machines, Machine{moveA, moveB, prize})
		}
	}
	return machines
}

func (machine *Machine) CalcSteps() int {
	alpha := (machine.p.y*machine.b.x - machine.p.x*machine.b.y) / (machine.a.y*machine.b.x - machine.b.y*machine.a.x)
	beta := (machine.p.x - machine.a.x*alpha) / machine.b.x
	if alpha*machine.a.x+beta*machine.b.x == machine.p.x && alpha*machine.a.y+beta*machine.b.y == machine.p.y {
		return 3*alpha + beta
	}

	return 0
}

func Combine(machines []Machine) int {
	sum := 0
	for _, machine := range machines {
		sum += machine.CalcSteps()
	}
	return sum
}

func Combine2(machines []Machine) int {
	sum := 0
	for _, machine := range machines {
		machine.p.x += 10000000000000
		machine.p.y += 10000000000000
		sum += machine.CalcSteps()
	}
	return sum
}
