package stacks

// https://leetcode.com/problems/check-if-a-parentheses-string-can-be-valid/

func canBeValid(s string, locked string) bool {
	var n = len(s)
	if n&1 == 1 {
		return false
	}
	// check closing brackets
	var lockedCnt, unlockedCnt = 0, 0
	for idx := range s {
		if locked[idx] == '0' {
			unlockedCnt++
		} else if s[idx] == '(' {
			lockedCnt++
		} else if lockedCnt > 0 {
			lockedCnt--
		} else if unlockedCnt > 0 {
			unlockedCnt--
		} else {
			return false
		}
	}
	// check opening brackets
	lockedCnt, unlockedCnt = 0, 0
	for idx := n - 1; idx >= 0; idx-- {
		if locked[idx] == '0' {
			unlockedCnt++
		} else if s[idx] == ')' {
			lockedCnt++
		} else if lockedCnt > 0 {
			lockedCnt--
		} else if unlockedCnt > 0 {
			unlockedCnt--
		} else {
			return false
		}
	}
	return true
}
