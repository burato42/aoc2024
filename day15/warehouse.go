package day15

import (
	"bufio"
	"os"
	"strings"
)

func ReadInputs(textFile string) ([][]string, []string) {
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
		}
		warehouse = append(warehouse, strings.Split(line, ""))
	}


	return warehouse, directions
}