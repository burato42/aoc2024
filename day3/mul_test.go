package day3

import "testing"

func TestSumMuls(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	actual := SumMuls(input)
	expected := 161
	if actual != expected {
		t.Errorf("SumMuls(%s) = %d, want %d", input, actual, expected)
	}
}

func TestSumMulsWithAccuracy(t *testing.T) {
	input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	actual := SumMulsWithAccuracy(input)
	expected := 48
	if actual != expected {
		t.Errorf("SumMulsWithAccuracy(%s) = %d, want %d", input, actual, expected)
	}
}
