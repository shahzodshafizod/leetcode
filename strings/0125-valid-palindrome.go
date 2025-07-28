package strings

import (
	"unicode"
)

/*
Subproblems - A subproblem is a problem we have to solve along the way to solving the main problem.
Example: Given a string, find the length of the longest substring without repeating characters.
	main problem: find the length of the longest unique substring
	sub problem: pattern matching - unique substring

Palindromes - A palindrome is a string that reads the same forwards and backwards.
Examples:
	- "aba", "a", "race car", "A man, a plan, a canal: Panama" -> "amanaplanacanalpanama"

Problem:
Given a string, determine if it is a palindrome, considering only alphanumeric characters
	and ignoring case sensitivity.

Step 1: Verify the constraints
Step 2: Write out some test cases
	- "aabaa" -> true
	- "aabbaa" -> true
	- "abc" -> false
	- "a" -> true
	- "" -> true
	- "aba" -> true
	- "race car" -> true
	- "A man, a plan, a canal: Panama" -> true
*/

// https://leetcode.com/problems/valid-palindrome/

func isPalindrome(s string) bool {
	length := len(s)
	if length <= 1 {
		return true
	}

	var leftChar, rightChar rune
	for left, right := 0, length-1; left < right; left, right = left+1, right-1 {
		leftChar = rune(s[left])
		for left < right && !unicode.In(leftChar, unicode.Letter, unicode.Digit) {
			left++
			leftChar = rune(s[left])
		}

		rightChar = rune(s[right])
		for right > left && !unicode.In(rightChar, unicode.Letter, unicode.Digit) {
			right--
			rightChar = rune(s[right])
		}

		if left >= right {
			break
		}

		if unicode.ToLower(leftChar) != unicode.ToLower(rightChar) {
			return false
		}
	}

	return true
}

// func isPalindrome(s string) bool {
// 	s = strings.ToLower(regexp.MustCompile(`[^A-Za-z0-9]`).ReplaceAllString(s, ""))
// 	for left, right := 0, len(s)-1; left < right; {
// 		if s[left] != s[right] {
// 			return false
// 		}
// 		left++
// 		right--
// 	}
// 	return true
// }
