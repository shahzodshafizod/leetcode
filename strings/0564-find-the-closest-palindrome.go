package strings

import "strconv"

// https://leetcode.com/problems/find-the-closest-palindrome/

// Approach: Binary Search
// Time: O(Len(n) * Log(n))
// Space: O(n)
func nearestPalindromic(n string) string {
	makePalindrome := func(num int64) int64 {
		b := []byte(strconv.FormatInt(num, 10))
		n := len(b)
		left := (n - 1) / 2
		right := n / 2
		for ; left >= 0; left, right = left-1, right+1 {
			b[right] = b[left]
		}
		num, _ = strconv.ParseInt(string(b), 10, 64)
		return num
	}
	prevPalindrome := func(num int64) int64 {
		var left, right int64 = 0, num - 1
		var palindrome, prev, mid int64
		for left <= right {
			mid = (right-left)/2 + left
			palindrome = makePalindrome(mid)
			if palindrome < num {
				prev = palindrome
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		return prev
	}
	nextPalindrome := func(num int64) int64 {
		var left, right int64 = num + 1, 10e17 + 1
		var palindrome, next, mid int64
		for left <= right {
			mid = (right-left)/2 + left
			palindrome = makePalindrome(mid)
			if palindrome > num {
				next = palindrome
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		return next
	}

	num, _ := strconv.ParseInt(n, 10, 64)
	prev := prevPalindrome(num)
	next := nextPalindrome(num)
	if (num - prev) <= (next - num) {
		return strconv.FormatInt(prev, 10)
	}
	return strconv.FormatInt(next, 10)
}

// // Approach: Brute-Force
// func nearestPalindromic(n string) string {
// 	var isPalindrome = func(num []byte) bool {
// 		for l, r := 0, len(num)-1; l < r; l, r = l+1, r-1 {
// 			if num[l] != num[r] {
// 				return false
// 			}
// 		}
// 		return true
// 	}
// 	var prevPalindrome = func(num []byte) ([]byte, int) {
// 		var steps = 0
// 		for {
// 			steps++
// 			var idx = len(num) - 1
// 			for ; idx >= 0; idx-- {
// 				if num[idx] != '0' {
// 					num[idx]--
// 					break
// 				}
// 				num[idx] = '9'
// 			}
// 			if num[0] == '0' {
// 				num = num[1:]
// 			}
// 			if isPalindrome(num) {
// 				break
// 			}
// 		}
// 		return num, steps
// 	}
// 	var nextPalindrome = func(num []byte) ([]byte, int) {
// 		var steps = 0
// 		for {
// 			steps++
// 			var idx = len(num) - 1
// 			var carry = 1
// 			for ; idx >= 0; idx-- {
// 				if num[idx] != '9' {
// 					num[idx]++
// 					carry = 0
// 					break
// 				}
// 				num[idx] = '0'
// 			}
// 			if carry == 1 {
// 				num = append([]byte{'1'}, num...)
// 			}
// 			if isPalindrome(num) {
// 				break
// 			}
// 		}
// 		return num, steps
// 	}
// 	var prev, pSteps = prevPalindrome([]byte(n))
// 	var next, nSteps = nextPalindrome([]byte(n))
// 	var palindrome = string(next)
// 	if pSteps <= nSteps {
// 		palindrome = string(prev)
// 	}
// 	if palindrome == "" {
// 		palindrome = "0"
// 	}
// 	return palindrome
// }
