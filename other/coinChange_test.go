package other

import (
	"slices"
	"testing"
)

func TestCoinChange(t *testing.T) {
	tests := []struct {
		input1   float64
		input2   float64
		expected []int
	}{
		{5, 0.99, []int{1, 0, 0, 0, 0, 4}},
		{3.14, 1.99, []int{0, 1, 1, 0, 0, 1}},
		{3, 0.01, []int{4, 0, 2, 1, 1, 2}},
		{4, 3.14, []int{1, 0, 1, 1, 1, 0}},
		{0.45, 0.34, []int{1, 0, 1, 0, 0, 0}},
	}

	for _, test := range tests {
		result := CoinChange(test.input1, test.input2)
		if slices.Compare(result, test.expected) != 0 {
			t.Errorf("CoinChange(%v, %v) returned %v but expected %v", test.input1, test.input2, result, test.expected)
		} else {
			t.Logf("CoinChange(%v, %v) = returned %v", test.input1, test.input2, result)
		}

	}
}
