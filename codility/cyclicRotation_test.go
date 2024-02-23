package codility

import (
	"slices"
	"testing"
)

func TestCyclicRotation(t *testing.T) {
	tests := []struct {
		input    []int
		rotation int
		expected []int
	}{
		{[]int{}, 1, []int{}},
		{[]int{3, 8, 9, 7, 6}, 3, []int{9, 7, 6, 3, 8}},
	}

	for _, test := range tests {
		result := CyclicRotation(test.input, test.rotation)
		if slices.Compare(result, test.expected) != 0 {
			t.Errorf("CyclicRotation(%v, %v) = %v, expected %v", test.input, test.rotation, result, test.expected)
		}
	}
}
