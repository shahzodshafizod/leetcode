package monotonic

// https://leetcode.com/problems/find-the-number-of-subarrays-where-boundary-elements-are-maximum/

func numberOfSubarrays(nums []int) int64 {
	var count int64 = 0
	var stack = make([][2]int, len(nums))
	var size = 0
	for _, num := range nums {
		for size > 0 && stack[size-1][0] < num {
			size--
		}
		if size == 0 || stack[size-1][0] != num {
			stack[size][0] = num
			stack[size][1] = 0
			size++
		}
		stack[size-1][1]++
		count += int64(stack[size-1][1])
	}
	return count
}
