package strings

/*
Problem:
Given a string, determine if it is almost a palindrome. A string is almost a palindrome
	if it becomes a palindrome by removing 1 letter. Consider only alphanumeric characters
	and ignore case sensitivity.

Step 1: Verify the constraints
	- Do we consider a palindrome as almost a palindrome?
		: Yes, return true is the string is already a palindrome.
Step 2: Write out some test cases
	- "race a car" -> true ("race car" || "rac a car")
	- "abccdba" -> true ("abccba")
	- "abcdefdba" -> false
	- "" -> true
	- "a" -> true
	- "ab" -> true ("a" || "b")
	- "abcdcdba" ->true ("abdcdba" || "abcdcba")
	- "abcdfcdba" -> false
	- "abcdfefcdba" -> false
*/

// https://leetcode.com/problems/valid-palindrome-ii/

func validPalindrome(s string) bool {
	var left2 = -1
	var right2 int
	var skips = 0
	for left1, right1 := 0, len(s)-1; left1 < right1 && skips < 2; left1, right1 = left1+1, right1-1 {
		if s[left1] != s[right1] {
			skips++
			if left2 < 0 {
				left2 = left1
				right2 = right1 - 1
			}
			right1++
		}
	}
	if skips < 2 {
		return true
	}
	skips = 1
	for ; left2 < right2 && skips < 2; left2, right2 = left2+1, right2-1 {
		if s[left2] != s[right2] {
			skips++
			left2--
		}
	}
	return skips < 2
}

// func validPalindrome(s string) bool {
// 	length := len(s)
// 	if length <= 2 {
// 		return true
// 	}

// 	left, right := 0, length-1
// 	for left < right {
// 		if s[left] != s[right] {
// 			return validPalindromeLogic(&s, left+1, right) ||
// 				validPalindromeLogic(&s, left, right-1)
// 		}
// 		left++
// 		right--
// 	}

// 	return true
// }

// func validPalindromeLogic(s *string, left int, right int) bool {
// 	for left < right {
// 		if (*s)[left] != (*s)[right] {
// 			return false
// 		}
// 		left++
// 		right--
// 	}
// 	return true
// }
