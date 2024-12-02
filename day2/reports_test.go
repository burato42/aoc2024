package day2

import "testing"

func TestIsSafe(t *testing.T) {
	testCases := []struct {
		input []int
		expected bool
	}{
		{[]int{7, 6, 4, 2, 1}, true},
		{[]int{1, 2, 7, 8, 9}, false},
		{[]int{9, 7, 6, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, false},
		{[]int{8, 6, 4, 4, 1}, false},
		{[]int{1, 3, 6, 7, 9}, true},
	}

	for _, tc := range testCases {
		actual := isSafe(tc.input)
		if actual != tc.expected {
			t.Errorf("isSafe(%v) = %v, want %v", tc.input, actual, tc.expected)
		}
	}
}

func TestIsSafeWithTolerance(t *testing.T) {
	testCases := []struct {
		input []int
		expected bool
	}{
		{[]int{7, 6, 4, 2, 1}, true},
		{[]int{1, 2, 7, 8, 9}, false},
		{[]int{9, 7, 6, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, true},
		{[]int{8, 6, 4, 4, 1}, true},
		{[]int{1, 3, 6, 7, 9}, true},
	}

	for _, tc := range testCases {
		actual := isSafeWithTolerance(tc.input)
		if actual != tc.expected {
			t.Errorf("isSafe(%v) = %v, want %v", tc.input, actual, tc.expected)
		}
	}
}