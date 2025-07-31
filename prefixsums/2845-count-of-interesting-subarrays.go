package prefixsums

// https://leetcode.com/problems/count-of-interesting-subarrays/

// Approach: Prefix Sum + Hash Table
// Time: O(n)
// Space: O(n)
func countInterestingSubarrays(nums []int, modulo int, k int) int64 {
	precount := 0
	freq := map[int]int{0: 1}

	var total int64

	for _, num := range nums {
		if num%modulo == k {
			precount++
		}

		total += int64(freq[(precount-k+modulo)%modulo])
		freq[precount%modulo]++
	}

	return total
}
