package greedy

// https://leetcode.com/problems/check-if-a-parentheses-string-can-be-valid/

// Approach: Greedy + Stack
// Time: O(n)
// Space: O(1)
func canBeValid(s string, locked string) bool {
	var valid = func(start int, end int, step int, bracket byte) bool {
		var lockedCnt, unlockedCnt = 0, 0
		for idx := start; idx != end; idx += step {
			if locked[idx] == '0' {
				unlockedCnt++
			} else if s[idx] == bracket {
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
	var n = len(s)
	return n&1 == 0 && valid(0,n,1,'(') && valid(n-1,-1,-1,')')
}

