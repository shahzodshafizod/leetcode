package slidingwindows

// https://leetcode.com/problems/longest-harmonious-subsequence/

// Approach #2: Hash Table
// Time: O(n)
// Space: O(n)
func findLHS(nums []int) int {
	freq := make(map[int]int, len(nums))
	for _, num := range nums {
		freq[num]++
	}

	length := 0

	for num, cnt1 := range freq {
		if cnt2, ok := freq[num+1]; ok {
			length = max(length, cnt1+cnt2)
		}
	}

	return length
}

// // Approach #1: Sliding Window
// // Time: O(n log n)
// // Space: O(1)
// func findLHS(nums []int) int {
// 	sort.Ints(nums)
// 	var length = 0
// 	for left, right, n := 0, 1, len(nums); right < n; right++ {
// 		for nums[right]-nums[left] > 1 {
// 			left++
// 		}
// 		if nums[right]-nums[left] == 1 {
// 			length = max(length, right-left+1)
// 		}
// 	}
// 	return length
// }
