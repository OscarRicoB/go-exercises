package codility

import (
	"slices"
	"testing"
)

func TestMaxCounters(t *testing.T) {
	tests := []struct {
		input1   int
		input2   []int
		expected []int
	}{
		{3, []int{1, 2, 3}, []int{1, 1, 1}},
		{5, []int{3, 4, 4, 6, 1, 4, 4}, []int{3, 2, 2, 4, 2}},
	}

	for _, test := range tests {
		result := MaxCounters(test.input1, test.input2)
		if slices.Compare(result, test.expected) != 0 {
			t.Errorf("MaxCounters(%v, %v) returned %v but expected %v", test.input1, test.input2, result, test.expected)
		} else {
			t.Logf("MaxCounters(%v, %v) returned %v", test.input1, test.input2, result)
		}
	}
}
