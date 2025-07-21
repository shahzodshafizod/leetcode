package arrays

// https://leetcode.com/problems/three-consecutive-odds/

func threeConsecutiveOdds(arr []int) bool {
	count := 0
	for _, num := range arr {
		if num&1 == 1 {
			count++
		} else {
			count = 0
		}
		if count == 3 {
			return true
		}
	}
	return false
}
