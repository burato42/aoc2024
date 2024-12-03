package utils

import (
	"bufio"
	"os"
	"strconv"
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
