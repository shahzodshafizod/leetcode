package hashes

// https://leetcode.com/problems/contiguous-array/

func findMaxLength(nums []int) int {
	length := 0
	hashset := map[int]int{0: -1}
	sum := 0
	for idx, num := range nums {
		if num == 0 {
			num = -1
		}
		sum += num
		if index, exists := hashset[sum]; exists {
			length = max(length, idx-index)
		} else {
			hashset[sum] = idx
		}
	}
	return length
}
