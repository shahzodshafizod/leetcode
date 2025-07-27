package arrays

// https://leetcode.com/problems/count-hills-and-valleys-in-an-array/

func countHillValley(nums []int) int {
	cnt, prv := 0, 0
	for cur, n := 1, len(nums); cur < n-1; cur++ {
		if nums[prv] < nums[cur] && nums[cur] > nums[cur+1] ||
			nums[prv] > nums[cur] && nums[cur] < nums[cur+1] {
			cnt++
			prv = cur
		}
	}
	return cnt
}
