package greedy

// https://leetcode.com/problems/maximum-value-of-an-ordered-triplet-ii/

func maximumTripletValueII(nums []int) int64 {
	var value int64 = 0

	dmax, lmax := 0, 0
	for _, num := range nums {
		value = max(value, int64(dmax*num))
		dmax = max(dmax, lmax-num)
		lmax = max(lmax, num)
	}

	return value
}
