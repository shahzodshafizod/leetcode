package hashes

// https://leetcode.com/problems/subarray-sums-divisible-by-k/

// time: O(n)
// space: O(k)
func subarraysDivByK(nums []int, k int) int {
	mods := make([]int, k)
	mods[0] = 1

	count, prefmod := 0, 0
	for _, num := range nums {
		// take modulo twice to avoid -ve remainders (indices)
		prefmod = ((prefmod+num)%k + k) % k
		count += mods[prefmod]
		mods[prefmod]++
	}

	return count
}

// // time: O(n ^ 2)
// // space: O(1)
// func subarraysDivByK(nums []int, k int) int {
// 	var count = 0
// 	var sum int
// 	for i, n := 0, len(nums); i < n; i++ {
// 		sum = 0
// 		for j := i; j < n; j++ {
// 			sum += nums[j]
// 			if sum%k == 0 {
// 				count++
// 			}
// 		}
// 	}
// 	return count
// }
