package challenge1

import (
	"sort"
	"strings"
)

// Check if a string has all unique characters
// This function assumes that the input accepts only
// ASCII values
func isUnique(word string) bool {
	// We create a slice of runes because it
	// represents a codepoint for every ASCII char
	// Runes are just int32 values
	seen := make([]rune, len(word))
	for _, char := range word {
		if containsChar(char, seen) {
			return false
		}
		seen = append(seen, char)
	}
	return true
}

// Check if a string has all unique characters
// This implementation does not use any additional structures (maps, slices, etc)
// It sorts the strig alphabetically and compares the current char with the previous one
func isUniqueWithSort(word string) bool {
	chars := strings.Split(word, "")
	sort.Strings(chars)
	previousChar := chars[0]

	for index, element := range chars {
		if index == 0 {
			continue
		}
		if element == previousChar {
			return false
		}
		previousChar = element
	}

	return true
}

// Check if a given array of codepoints has the given char
func containsChar(char rune, array []rune) bool {
	for _, element := range array {
		if element == char {
			return true
		}
	}
	return false
}
