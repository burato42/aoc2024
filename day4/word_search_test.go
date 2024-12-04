package day4

import "testing"

func TestFindWords(t *testing.T) {
	input := "./../day4/sample.txt"
	actual := FindWords(ReadTextToMatrix(input))
	expected := 18
	if actual != expected {
		t.Errorf("FindWords(%s) = %d, want %d", input, actual, expected)
	}
}

func TestFindXWords(t *testing.T) {
	input := "./../day4/sample.txt"
	actual := FindXWords(ReadTextToMatrix(input))
	expected := 9
	if actual != expected {
		t.Errorf("FindXWords(%s) = %d, want %d", input, actual, expected)
	}
}