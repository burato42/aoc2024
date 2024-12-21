package day17

import (
	"aoc2024/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"strings"
)

func ReadInput(textFile string) (map[string]int, []int) {
	file, err := os.Open(textFile)
	if err != nil {
		panic("Can't read from the file")
	}
	defer file.Close()

	registers := make(map[string]int)
	program := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Register A:") {
			registers["A"] = utils.DigitToInt(strings.Split(line, " ")[2])
		}
		if strings.HasPrefix(line, "Register B:") {
			registers["B"] = utils.DigitToInt(strings.Split(line, " ")[2])
		}
		if strings.HasPrefix(line, "Register C:") {
			registers["C"] = utils.DigitToInt(strings.Split(line, " ")[2])
		}
		if strings.HasPrefix(line, "Program:") {
			digits := strings.Split(line, " ")[1]
			for _, digit := range strings.Split(digits, ",") {
				program = append(program, utils.DigitToInt(digit))
			}
		}
	}

	return registers, program
}

func Output(registers map[string]int, program []int) []int {
	res := []int{}
	var pointer int
	
	combo := map[int]func() int {
		0: func() int {return 0},
		1: func() int {return 1},
		2: func() int {return 2},
		3: func() int {return 3},
		4: func() int {return registers["A"]},
		5: func() int {return registers["B"]},
		6: func() int {return registers["C"]},
	}

	operand := map[int]func(int) {
		0: func(x int) {
			registers["A"] = int(float64(registers["A"])/math.Pow(2.0, float64(combo[x]())))
			pointer += 2
		},
		1: func(x int) {
			registers["B"] ^= x
			pointer += 2
		},
		2: func(x int) {
			registers["B"] = combo[x]()%8
			pointer += 2
		},
		3: func(x int) {
			if registers["A"] == 0 {
				pointer += 2
				return
			}
			pointer = x
		},
		4: func(_ int) {
			registers["B"] ^= registers["C"]
			pointer += 2
		},
		5: func(x int) {
			res = append(res, combo[x]()%8)
			pointer += 2
		},
		6: func(x int) {
			registers["B"] = int(float64(registers["A"])/math.Pow(2.0, float64(combo[x]())))
			pointer += 2
		},
		7: func(x int) {
			registers["C"] = int(float64(registers["A"])/math.Pow(2.0, float64(combo[x]())))
			pointer += 2
		},
	}
	for pointer < len(program) - 1 {
		operand[program[pointer]](program[pointer+1])
	}

	return res

}


func OutputPart2(registers map[string]int, program []int, stop bool) []int {
	res := []int{}
	var pointer int
	var addPointer int
	
	combo := map[int]func() int {
		0: func() int {return 0},
		1: func() int {return 1},
		2: func() int {return 2},
		3: func() int {return 3},
		4: func() int {return registers["A"]},
		5: func() int {return registers["B"]},
		6: func() int {return registers["C"]},
	}

	operand := map[int]func(int) {
		0: func(x int) {
			registers["A"] = int(float64(registers["A"])/math.Pow(2.0, float64(combo[x]())))
			pointer += 2
		},
		1: func(x int) {
			registers["B"] ^= x
			pointer += 2
		},
		2: func(x int) {
			registers["B"] = combo[x]()%8
			pointer += 2
		},
		3: func(x int) {
			if registers["A"] == 0 {
				pointer += 2
				return
			}
			pointer = x
		},
		4: func(_ int) {
			registers["B"] ^= registers["C"]
			pointer += 2
		},
		5: func(x int) {
			next := combo[x]()%8
			if next != program[addPointer] {
				stop = true
			} else {
				addPointer++
				res = append(res, next)
				pointer += 2
			}
		},
		6: func(x int) {
			registers["B"] = int(float64(registers["A"])/math.Pow(2.0, float64(combo[x]())))
			pointer += 2
		},
		7: func(x int) {
			registers["C"] = int(float64(registers["A"])/math.Pow(2.0, float64(combo[x]())))
			pointer += 2
		},
	}
	for pointer < len(program) - 1 {
		if stop {
			// fmt.Println(res)
			return res
		}
		operand[program[pointer]](program[pointer+1])
	}

	return res

}

func FixCorruption(registers map[string]int, program []int) {
	start := 140737488355328
	end := 281474976710656
	for i := start; i <= end; i++ {
		registers["A"] = i
		if reflect.DeepEqual(OutputPart2(registers, program, false), program ) {
			fmt.Println(i)
			return
		}
	}
}