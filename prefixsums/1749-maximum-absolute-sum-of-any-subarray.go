package prefixsums

// https://leetcode.com/problems/maximum-absolute-sum-of-any-subarray/

func maxAbsoluteSum(nums []int) int {
	var posPrefixSum, maxSum = 0, 0
	var negPrefixSum, minSum = 0, 0
	for _, num := range nums {
		posPrefixSum = max(0, posPrefixSum+num)
		maxSum = max(maxSum, posPrefixSum)
		negPrefixSum = min(0, negPrefixSum+num)
		minSum = min(minSum, negPrefixSum)
	}
	return max(maxSum, -minSum)
}
