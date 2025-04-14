package arrays

// https://leetcode.com/problems/count-good-triplets/

// Approach: Enumeration
// Time: O(nnn)
// Space: O(1)
func countGoodTriplets(arr []int, a int, b int, c int) int {
	var abs = func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	var count, n = 0, len(arr)
	var i, j, k, x, y, z int
	for i = 0; i < n; i++ {
		x = arr[i]
		for j = i + 1; j < n; j++ {
			y = arr[j]
			if abs(x-y) > a {
				continue
			}
			for k = j + 1; k < n; k++ {
				z = arr[k]
				if abs(y-z) <= b && abs(x-z) <= c {
					count++
				}
			}
		}
	}
	return count
}
