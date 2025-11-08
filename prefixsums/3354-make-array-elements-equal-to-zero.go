package prefixsums

// https://leetcode.com/problems/make-array-elements-equal-to-zero/

func countValidSelections(nums []int) int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}

		return x
	}

	ans, left, right := 0, 0, 0
	for _, num := range nums {
		right += num
	}

	for _, num := range nums {
		left += num
		right -= num

		if num == 0 {
			if right == left {
				ans += 2
			}

			if abs(left-right) == 1 {
				ans++
			}
		}
	}

	return ans
}
