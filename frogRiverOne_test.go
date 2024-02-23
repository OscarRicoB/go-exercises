package main

import "testing"

func TestFrogRiverOne(t *testing.T) {
	tests := []struct {
		input1   int
		input2   []int
		expected int
	}{
		{5, []int{1, 3, 1, 4, 2, 3, 5, 4}, 6},
		{3, []int{1, 9, 8, 4, 6, 7, 2, 1, 3, 5}, 8},
		{5, []int{1, 3, 1, 4, 2, 3, 6, 5, 4}, 7},
		{1, []int{1, 2}, 0},
	}

	for _, test := range tests {
		result := FrogRiverOne(test.input1, test.input2)
		if result != test.expected {
			t.Errorf("FrogRiverOne(%v, %v) = %v, expected %v", test.input1, test.input2, result, test.expected)
		}
	}
}
