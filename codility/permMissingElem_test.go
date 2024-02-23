package codility

import "testing"

func TestPermMissingElem(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 6, 7, 8}, 5},
		{[]int{2, 3, 1, 5}, 4},
	}

	for _, test := range tests {
		result := PermMissingElem(test.input)
		if result != test.expected {
			t.Errorf("PermMissingElem(%v) returned %v but expected %v", test.input, result, test.expected)
		} else {
			t.Logf("PermMissingElem(%v) returned %v", test.input, result)
		}
	}

}
