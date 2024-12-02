package day1

import (
	"testing"
)


func TestCalculateDistance(t *testing.T) {
	testCases := []struct {
		left []int
		right []int
		expected int
	}{
		{[]int{1, 2}, []int{3, 4}, 4},
		{[]int{4, 3}, []int{1, 2}, 4},
		{[]int{2, -1}, []int{-3, 4}, 4},
		{[]int{2, -1}, []int{4, -3}, 4},
	}

	for _, tc := range testCases {
		actual := CalculateDist(tc.left, tc.right)
		if actual != tc.expected {
			t.Errorf("CalculateDist(%v, %v) = %d, want %d", tc.left, tc.right, actual, tc.expected)
		}
	}
}

func TestCalculateSimilarity(t *testing.T) {
	testCases := []struct {
		left []int
		right []int
		expected int
	}{
		{[]int{1, 2}, []int{3, 4}, 0},
		{[]int{1, 2}, []int{1, 4}, 1},
		{[]int{1, 2}, []int{1, 1}, 2},
		{[]int{1, 2, 2}, []int{1, 1, 2}, 6},
		{[]int{3, 4, 2, 1, 3, 3}, []int{4, 3, 5, 3, 9, 3}, 31},
	}

	for _, tc := range testCases {
		actual := CalculateSimilarity(tc.left, tc.right)
		if actual != tc.expected {
			t.Errorf("CalculateSimilarity(%v, %v) = %d, want %d", tc.left, tc.right, actual, tc.expected)
		}
	}
}