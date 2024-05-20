package bits

// https://leetcode.com/problems/sum-of-all-subset-xor-totals/

// time: O(2 ^ n)
// space: O(2 ^ n)
func subsetXORSum(nums []int) int {
	var dfs func(idx int, sum int) int
	dfs = func(idx int, sum int) int {
		if idx < 0 {
			return sum
		}
		return dfs(idx-1, sum^nums[idx]) + // decision to include
			dfs(idx-1, sum) // decision NOT to include
	}
	return dfs(len(nums)-1, 0)
}

// // time: O(n)
// // space: O(1)
// func subsetXORSum(nums []int) int {
// 	var sum = 0
// 	for _, num := range nums {
// 		sum |= num
// 	}
// 	sum <<= len(nums) - 1
// 	return sum
// }
