package bits

// https://leetcode.com/problems/single-number/

func singleNumber(nums []int) int {
	var single = 0
	for _, num := range nums {
		single ^= num
	}
	return single
}
