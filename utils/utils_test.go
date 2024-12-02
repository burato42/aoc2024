package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestDigitToIntHappy(t *testing.T) {
	testCases := []struct {
		input string 
		expected int
	}{
		{"123", 123},
		{"-256", -256},
	}

	for _, tc := range testCases {
		actual := DigitToInt(tc.input)
		if actual != tc.expected {
			t.Errorf("digitToInt(%s) = %v, want %d", tc.input, actual, tc.expected)
		}
	}
}

func TestDigitToIntNegative(t *testing.T) {
	testCases := []string{"", "-234.56", " "}

	for _, tc := range testCases {
		assert.Panics(t, func() {DigitToInt(tc)}, "The code didn't panic")
	}
}