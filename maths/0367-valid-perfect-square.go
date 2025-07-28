package maths

// https://leetcode.com/problems/valid-perfect-square/

// time: O(log n), n = num
// space: O(1)
func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}

	left, right := 2, num/2

	var mid, square int

	for left <= right {
		mid = (left + right) / 2
		square = mid * mid

		if square > num {
			right = mid - 1
		} else if square < num {
			left = mid + 1
		} else {
			return true
		}
	}

	return false
}
