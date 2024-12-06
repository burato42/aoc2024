package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func DigitToInt(digit string) int {
	number, err := strconv.Atoi(digit)
	if err != nil {
		panic("We have not a number, we can't proceed")
	}
	return number
}

func ReadTextFile(textFile string) string {
	file, err := os.Open(textFile)
	if err != nil {
		panic("Can't read from the file")
	}
	defer file.Close()

	content := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		content += line
	}

	return content
}

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

func IsInside(v, h, height, width int) bool {
	return v >= 0 && h >= 0 && v < height && h < width
}