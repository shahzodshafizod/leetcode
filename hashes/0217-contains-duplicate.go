package hashes

// https://leetcode.com/problems/contains-duplicate/

func containsDuplicate(nums []int) bool {
	var seen = make(map[int]bool)
	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	return false
}
