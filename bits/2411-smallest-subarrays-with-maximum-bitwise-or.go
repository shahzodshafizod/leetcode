package bits

// https://leetcode.com/problems/smallest-subarrays-with-maximum-bitwise-or/

func smallestSubarrays(nums []int) []int {
	n := len(nums)

	var pos [32]int

	res := make([]int, n)

	var maxPos, num int
	for i := n - 1; i >= 0; i-- {
		maxPos, num = i, nums[i]
		for j := 0; j < 32; j++ {
			if num&1 == 1 {
				pos[j] = i
			}

			num >>= 1
			maxPos = max(maxPos, pos[j])
		}

		res[i] = maxPos - i + 1
	}

	return res
}
