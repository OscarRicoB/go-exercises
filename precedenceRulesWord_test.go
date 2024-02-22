package main

import (
	"testing"
)

func TestFindWord(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{"a > b", "b > c", "c > d", "d > e"}, "abcde"},
		{[]string{"x > y", "y > z"}, "xyz"},
		{[]string{"P>E", "E>R", "R>U"}, "PERU"},
		{[]string{"I>N", "A>I", "P>A", "S>P"}, "SPAIN"},
		{[]string{"U>N", "G>A", "R>Y", "H>U", "N>G", "A>R"}, "HUNGARY"},
		{[]string{"I>F", "W>I", "S>W", "F>T"}, "SWIFT"},
		{[]string{"R>T", "A>L", "P>O", "O>R", "G>A", "T>U", "U>G"}, "PORTUGAL"},
		{[]string{"W>I", "R>L", "T>Z", "Z>E", "S>W", "E>R", "L>A", "A>N", "N>D", "I>T"}, "SWITZERLAND"},
		//TODO:
		//{[]string{"a > b", "b > c", "c > a"}, "Invalid input, cyclic dependency detected"},
	}

	for _, test := range tests {
		result := FindWord(test.input)
		if result != test.expected {
			t.Errorf("findWord(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}
