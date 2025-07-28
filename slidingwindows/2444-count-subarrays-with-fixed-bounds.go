package slidingwindows

// https://leetcode.com/problems/count-subarrays-with-fixed-bounds/

func countSubarrays(nums []int, minK int, maxK int) int64 {
	var count int64 = 0

	badid, minid, maxid := -1, -1, -1

	for idx, num := range nums {
		if num == minK {
			minid = idx
		}

		if num == maxK {
			maxid = idx
		}

		if num < minK || num > maxK {
			badid, minid, maxid = idx, idx, idx
		}

		count += int64(min(minid, maxid) - badid)
	}

	return count
}
