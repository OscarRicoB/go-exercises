package codility

import "testing"

func TestFrogJmp(t *testing.T) {
	tests := []struct {
		input1   int
		input2   int
		input3   int
		expected int
	}{
		{10, 85, 30, 3},
		{0, 100, 2, 50},
		{10, 15, 300, 1},
	}
	for _, test := range tests {
		result := FrogJmp(test.input1, test.input2, test.input3)
		if result != test.expected {
			t.Errorf("FrogJmp(%v, %v, %v) = %v, expected %v", test.input1, test.input2, test.input3, result, test.expected)
		}
	}
}
