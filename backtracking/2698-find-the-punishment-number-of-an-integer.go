package backtracking

import "strconv"

// https://leetcode.com/problems/find-the-punishment-number-of-an-integer/

// Approach: Recursion of Strings
// Time: O(nn)
// Space: O(n)
func punishmentNumber(n int) int {
	var partition func(target int, idx int, s string) bool
	partition = func(target int, idx int, s string) bool {
		if idx == len(s) || target < 0 {
			return target == 0
		}
		var num = 0
		for _, c := range s[idx:] {
			num = num*10 + int(c-'0')
			if partition(target-num, idx+1, s) {
				return true
			}
			idx++
		}
		return false
	}
	var panishment = 0
	for num := 1; num <= n; num++ {
		if partition(num, 0, strconv.Itoa(num*num)) {
			panishment += num * num
		}
	}
	return panishment
}
