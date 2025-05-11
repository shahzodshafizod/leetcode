package greedy

// https://leetcode.com/problems/minimum-equal-sum-of-two-arrays-after-replacing-zeros/

func minSum(nums1 []int, nums2 []int) int64 {
	var sum1, sum2 int64 = 0, 0
	var zeroes1, zeroes2 = 0, 0
	for _, num := range nums1 {
		sum1 += int64(num)
		if num == 0 {
			sum1++
			zeroes1++
		}
	}
	for _, num := range nums2 {
		sum2 += int64(num)
		if num == 0 {
			sum2++
			zeroes2++
		}
	}
	if zeroes1 == 0 && sum1 < sum2 || zeroes2 == 0 && sum2 < sum1 {
		return -1
	}
	return max(sum1, sum2)
}
