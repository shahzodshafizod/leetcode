package arrays

// https://leetcode.com/problems/check-if-all-1s-are-at-least-length-k-places-away/

// Approach: Single Pass with Distance Tracking
// Time: O(n)
// Space: O(1)
func kLengthApart(nums []int, k int) bool {
	distance := k

	for _, num := range nums {
		if num == 0 {
			distance++
		} else {
			if distance < k {
				return false
			}

			distance = 0
		}
	}

	return true
}
