package day4

import (
	"aoc2024/utils"
	"testing"
)

func TestFindWords(t *testing.T) {
	input := "./../day4/sample.txt"
	actual := FindWords(utils.ReadTextToMatrix(input))
	expected := 18
	if actual != expected {
		t.Errorf("FindWords(%s) = %d, want %d", input, actual, expected)
	}
}

func TestFindXWords(t *testing.T) {
	input := "./../day4/sample.txt"
	actual := FindXWords(utils.ReadTextToMatrix(input))
	expected := 9
	if actual != expected {
		t.Errorf("FindXWords(%s) = %d, want %d", input, actual, expected)
	}
}