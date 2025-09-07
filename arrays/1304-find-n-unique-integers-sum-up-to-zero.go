package arrays

// https://leetcode.com/problems/find-n-unique-integers-sum-up-to-zero/

func sumZero(n int) []int {
	nums := make([]int, n)
	diff := (2 - n&1) / 2

	for i := n / 2; i > 0; i-- {
		nums[i-diff] = i
		nums[n-i] = -i
	}

	return nums
}
