package maths

// https://leetcode.com/problems/palindrome-number/

func isPalindrome(x int) bool {
	if x < 0 || x != 0 && x%10 == 0 {
		return false
	}

	half := 0
	for x > half {
		half = half*10 + x%10
		x /= 10
	}

	return x == half || x == half/10
}

// func isPalindrome(x int) bool {
// 	var copy, rev = x, 0
// 	for copy > 0 {
// 		rev = rev*10 + copy%10
// 		copy /= 10
// 	}
// 	return x == rev
// }
