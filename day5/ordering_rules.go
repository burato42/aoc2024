package day5

import (
	"bufio"
	"os"
	"strings"
)

func ReadTextFile(textFile string) (map[string][]string, [][]string) {
	file, err := os.Open(textFile)
	if err != nil {
		panic("Can't read from the file")
	}
	defer file.Close()

	seqs := [][]string{}
	scanner := bufio.NewScanner(file)
	rules := make(map[string][]string)
	read_rules := true
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			read_rules = false
			continue
		}

		if read_rules {
			fromTo := strings.Split(line, "|")
			from := fromTo[0]
			to := fromTo[1]
			if res, ok := rules[from]; ok {
				rules[from] = append(res, to)
			} else {
				rules[from] = []string{to}
			}
		} else {
			seqs = append(seqs, strings.Split(line, ","))
		}
		
	}

	return rules, seqs
}