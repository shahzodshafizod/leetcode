package hashes

// https://leetcode.com/problems/count-equal-and-divisible-pairs-in-an-array/

// Approach: Hash Table
// Time: best: O(N); worst: O(NN)
// Space: O(N)
func countPairs(nums []int, k int) int {
	var pairs = 0
	var indices = make(map[int][]int)
	for curr, num := range nums {
		for _, prev := range indices[num] {
			if prev*curr%k == 0 {
				pairs++
			}
		}
		indices[num] = append(indices[num], curr)
	}
	return pairs
}
