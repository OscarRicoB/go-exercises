package codility

import (
	"testing"
)

func TestBinaryGap(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{1041, 5},
		{15, 0},
		{32, 0},
		{6, 0},
		{561892, 3},
	}

	for _, test := range tests {
		result := BinaryGap(test.input)
		if result != test.expected {
			t.Errorf("BinaryGap(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}
