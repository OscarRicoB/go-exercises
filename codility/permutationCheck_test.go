package codility

import "testing"

func TestPermCheck(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{4, 1, 3, 2}, 1},
		{[]int{4, 1, 3}, 0},
	}

	for _, test := range tests {
		result := PermCheck(test.input)
		if result != test.expected {
			t.Errorf("PermCheck(%v) returned %v but expected %v", test.input, result, test.expected)
		} else {
			t.Logf("PermCheck(%v) returned %v", test.input, result)
		}
	}
}
