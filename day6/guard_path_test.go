package day6

import (
	"aoc2024/utils"
	"reflect"
	"testing"
)

func TestAddObstacles(t *testing.T) {
	input := [][]string{{".", "."}, {".", "."}}
	actual := AddObstacles(input)
	expected := [][][]string{
		{{"#", "."}, {".", "."}},
		{{".", "#"}, {".", "."}},
		{{".", "."}, {"#", "."}},
		{{".", "."}, {".", "#"}},
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("AddObstacles(%v) = %v, want %v", input, actual, expected)
	}
}

func TestCountLoops(t *testing.T) {
	input := "./../day6/sample.txt"
	actual := CountLoops(utils.ReadTextToMatrix(input))
	expected := 6
	if actual != expected {
		t.Errorf("CountLoops(utils.ReadTextToMatrix(%s)) = %d, want %d", input, actual, expected)
	}
}

func TestCountSteps(t *testing.T) {
	input := "./../day6/sample.txt"
	actual := CountSteps(utils.ReadTextToMatrix(input))
	expected := 41
	if actual != expected {
		t.Errorf("CountSteps(utils.ReadTextToMatrix(%s)) = %d, want %d", input, actual, expected)
	}
}
