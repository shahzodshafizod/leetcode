package strings

// https://leetcode.com/problems/shortest-palindrome/

// Approach #3: Rabin Karp Algorithm
// Time: O(n)
// Space: O(1)
func shortestPalindrome(s string) string {
	var forward, backward = 0, 0
	const mod int = 1e9 + 7
	// 29 is the first prime number after base26 to decrease collisions
	const base int = 29
	var digit int
	var power = 1
	var length = 0
	for idx, c := range s {
		digit = int(c-'a') + 1
		forward = (forward*base + digit) % mod
		backward = (digit*power + backward) % mod
		if forward == backward {
			length = idx + 1
		}
		power = power * base % mod
	}
	var prefix = []byte(s[length:])
	for l, r := 0, len(prefix)-1; l < r; l, r = l+1, r-1 {
		prefix[l], prefix[r] = prefix[r], prefix[l]
	}
	return string(prefix) + s
}

// // Approach #2: KMP (Knuth-Morris-Pratt) Algorithm
// // Time: O(n)
// // Space: O(n)
// func shortestPalindrome(s string) string {
// 	var n = len(s)
// 	var rev = []byte(s)
// 	for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
// 		rev[l], rev[r] = rev[r], rev[l]
// 	}
// 	var combined = s + "#" + string(rev)
// 	n = n*2 + 1
// 	var kmp = make([]int, n)
// 	var length = 0
// 	for idx := 1; idx < n; idx++ {
// 		length = kmp[idx-1]
// 		for length > 0 && combined[idx] != combined[length] {
// 			length = kmp[length-1]
// 		}
// 		if combined[idx] == combined[length] {
// 			length++
// 		}
// 		kmp[idx] = length
// 	}
// 	return string(rev[:len(s)-length]) + s
// }

// // Approach #1: Brute-Force
// // Time: O(n^2)
// // Space: O(1)
// func shortestPalindrome(s string) string {
// 	var n = len(s)
// 	var isPalindrome = func(left int, right int) (bool, int) {
// 		for left >= 0 && right < n && s[left] == s[right] {
// 			left--
// 			right++
// 		}
// 		return left < 0, right
// 	}
// 	// 1. find the longest palindrome starting from 0
// 	var length = 0
// 	for left := (n - 1) / 2; left >= 0 && length == 0; left-- {
// 		for _, right := range [2]int{left + 1, left} {
// 			if ok, l := isPalindrome(left, right); ok {
// 				length = l
// 				break
// 			}
// 		}
// 	}
// 	// 2. reverse the remaining and add in the beginning
// 	var prefix = []byte(s[length:])
// 	for l, r := 0, len(prefix)-1; l < r; l, r = l+1, r-1 {
// 		prefix[l], prefix[r] = prefix[r], prefix[l]
// 	}
// 	return string(prefix) + s
// }
