package hashes

import (
	"math"
	"unicode"
)

// https://leetcode.com/problems/shortest-completing-word/

// Approach 2: Optimized using Map (Hash Table)
// Use map for cleaner frequency counting
// Time: O(n * m) where n is number of words, m is average word length
// Space: O(1) - map stores at most 26 letters
func shortestCompletingWord(licensePlate string, words []string) string {
	// Count letters in license plate (ignore case, numbers, spaces)
	plateCounter := make(map[rune]int)

	for _, char := range licensePlate {
		if unicode.IsLetter(char) {
			plateCounter[unicode.ToLower(char)]++
		}
	}

	result := ""
	minLength := math.MaxInt32

	for _, word := range words {
		// Count letters in current word
		wordCounter := make(map[rune]int)
		for _, char := range word {
			wordCounter[char]++
		}

		// Check if word has all required letters with required frequencies
		isValid := true

		for char, count := range plateCounter {
			if wordCounter[char] < count {
				isValid = false

				break
			}
		}

		if isValid && len(word) < minLength {
			minLength = len(word)
			result = word
		}
	}

	return result
}

// // Approach 1: Brute Force
// // For each word, manually count character frequencies and compare
// // Time: O(n * m) where n is number of words, m is average word length
// // Space: O(1) - fixed size array for 26 letters
// func shortestCompletingWord(licensePlate string, words []string) string {
// 	// Count letters in license plate (ignore case, numbers, spaces)
// 	countLetters := func(s string) [26]int {
// 		var counts [26]int

// 		for _, char := range s {
// 			if unicode.IsLetter(char) {
// 				counts[unicode.ToLower(char)-'a']++
// 			}
// 		}

// 		return counts
// 	}

// 	// Check if word contains all letters from license plate
// 	isCompleting := func(word string, targetCounts [26]int) bool {
// 		wordCounts := countLetters(word)
// 		for i := range 26 {
// 			if wordCounts[i] < targetCounts[i] {
// 				return false
// 			}
// 		}

// 		return true
// 	}

// 	targetCounts := countLetters(licensePlate)
// 	result := ""
// 	minLength := math.MaxInt32

// 	for _, word := range words {
// 		if isCompleting(word, targetCounts) {
// 			if len(word) < minLength {
// 				minLength = len(word)
// 				result = word
// 			}
// 		}
// 	}

// 	return result
// }
