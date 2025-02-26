package prefixsums

// https://leetcode.com/problems/number-of-sub-arrays-with-odd-sum/

// Approach: Prefix Sum
// Time: O(n)
// Space: O(1)
func numOfSubarrays(arr []int) int {
	var oddCnt, envCnt = 0, 0
	var totalCount, prefixSum = 0, 0
	const MOD int = 1e9 + 7
	for _, num := range arr {
		prefixSum += num
		if prefixSum&1 == 0 {
			totalCount = (totalCount + oddCnt) % MOD
			envCnt++
		} else {
			totalCount = (totalCount + envCnt + 1) % MOD
			oddCnt++
		}
	}
	return totalCount
}
