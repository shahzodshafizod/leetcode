package strings

/*
Problem:
Given a string, find the length of the longest substring without repeating characters.

Step 1: Verify the constraints
	- Is the substring contiguous?
		: Yes, look for a substring not a subsequence.
	- Does case sensitivity matter?
		: No, assume all characters in the string are lowercase.
Step 2: Write out some test cases
	- "abccabb" -> 3 ("abc", "cab")
	- "cccccc" -> 1 ("c")
	- "" -> 0 ("")
	- "abcbda" -> 4 ("abc", "cbda")
Sliding Window Technique:
	Form a window over some portion of sequential data, then move that window
	throughout the data to capture different parts of it.
Step 8: Can we optimize out solution?
	- 1st Hint: Use a sliding window to represent the current substring.
	- 2nd Hint: The size of the window will change based on new characters, and characters
		we've already seen before.
	- 3rd Hint: Our seen characters hashmap keeps track of what characters we've seen,
		and the index we saw them at.
*/

// https://leetcode.com/problems/longest-substring-without-repeating-characters/

// time: O(N)
// space: O(N)
func lengthOfLongestSubstring(s string) int {
	slength := len(s)
	if slength <= 1 {
		return slength
	}
	var maxLength = 0
	var hash = make(map[byte]int)
	left, right := 0, 0
	for ; right < slength; right++ {
		if index, exists := hash[s[right]]; exists && index >= left {
			maxLength = max(maxLength, right-left)
			left = index + 1
		}
		hash[s[right]] = right
		// maxLength = max(maxLength, right-left+1) // remove other maxLength = max()'s
	}
	maxLength = max(maxLength, right-left)
	return maxLength
}

// // time: O(N^2)
// // space: O(N)
// func lengthOfLongestSubstring(s string) int {
// 	length := len(s)
// 	if length <= 1 {
// 		return length
// 	}
// 	var maxLength = 0
// 	var hash map[byte]bool
// 	var right int
// 	for left := 0; left < length; left++ {
// 		hash = make(map[byte]bool)
// 		right = left
// 		for ; right < length; right++ {
// 			if hash[s[right]] {
// 				break
// 			}
// 			hash[s[right]] = true
// 		}
// 		maxLength = max(maxLength, len(hash))
// 		if right >= length {
// 			break
// 		}
// 	}
// 	return maxLength
// }

// // time: O(N^3)
// // space: O(1)
// func lengthOfLongestSubstring(s string) int {
// 	var maxLength = 0
// 	var length int
// 	var exists bool
// 	for k := len(s) - 1; k >= 0; k-- {
// 		length = 0
// 		for i := k; i >= 0; i-- {
// 			exists = false
// 			for j := k; j > i; j-- {
// 				if s[j] == s[i] {
// 					exists = true
// 					break
// 				}
// 			}
// 			if exists {
// 				maxLength = max(maxLength, length)
// 				length = 0
// 			}
// 			length++
// 		}
// 		maxLength = max(maxLength, length)
// 	}
// 	return maxLength
// }
