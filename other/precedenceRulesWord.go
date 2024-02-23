package other

import "strings"

// A precedence rule is given as "P>E", which means that letter "P" is followed directly by the letter "E".
// Write a function, given an array of precedence rules, that finds the word represented by the given rules.

// Note: Each represented word contains a set of unique characters, i.e. the word does not contain duplicate letters.

// Examples:
// findWord(["P>E","E>R","R>U"]) // PERU
// findWord(["I>N","A>I","P>A","S>P"]) // SPAIN

// findWord(["U>N", "G>A", "R>Y", "H>U", "N>G", "A>R"]) // HUNGARY
// findWord(["I>F", "W>I", "S>W", "F>T"]) // SWIFT
// findWord(["R>T", "A>L", "P>O", "O>R", "G>A", "T>U", "U>G"]) // PORTUGAL
// findWord(["W>I", "R>L", "T>Z", "Z>E", "S>W", "E>R", "L>A", "A>N", "N>D", "I>T"]) // SWITZERLAND

func FindWord(precedences []string) string {
	word, maybeStarts := []rune{}, []rune{}
	rules := map[rune]rune{}
	notFirstLetter := map[rune]bool{}
	var start rune
	for _, p := range precedences {
		p = strings.ReplaceAll(strings.TrimSpace(p), " ", "")
		if _, ok := rules[rune(p[0])]; !ok {
			rules[rune(p[0])] = rune(p[2])
			notFirstLetter[rune(p[2])] = true
		}
		if _, ok := notFirstLetter[rune(p[0])]; !ok {
			maybeStarts = append(maybeStarts, rune(p[0]))
		}
	}

	for _, s := range maybeStarts {
		if _, ok := notFirstLetter[s]; !ok {
			start = s
			break
		}
	}

	for {
		word = append(word, start)
		if next, ok := rules[start]; ok {
			start = next
		} else {
			break
		}
	}
	return string(word)
}
