package main

import (
	"testing"
)

func TestOddOccurrencesInArray(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{4, 3, 5, 6, 4, 3, 6}, 5},
		{[]int{9, 3, 9, 3, 9, 7, 9}, 7},
	}

	for _, test := range tests {
		result := OddOccurrencesInArray(test.input)
		if result != test.expected {
			t.Errorf("OddOccurrencesInArray(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}
