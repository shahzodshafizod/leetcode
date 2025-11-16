package twopointers

// https://leetcode.com/problems/happy-number/

// Approach: Floyd's Tortoise and Hare (Sangpusht va Xargushi Floyd) for cycle detection
// Space: O(1)
func isHappy(n int) bool {
	getSum := func(n int) int {
		var sum, digit int
		for ; n > 0; n /= 10 {
			digit = n % 10
			sum += digit * digit
		}

		return sum
	}

	slow, fast := n, n
	for {
		slow = getSum(slow)
		fast = getSum(fast)
		fast = getSum(fast)

		if fast == 1 {
			return true
		}

		if slow == fast {
			break
		}
	}

	return false
}
