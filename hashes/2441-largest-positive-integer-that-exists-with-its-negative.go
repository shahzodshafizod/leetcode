package hashes

// https://leetcode.com/problems/largest-positive-integer-that-exists-with-its-negative/

// time: O(n)
// space: O(m)
func findMaxK(nums []int) int {
	var bitset [1001]int16
	k := -1
	var absnum int
	for _, num := range nums {
		if num < 0 {
			absnum = -num
			bitset[absnum] |= 1
		} else {
			absnum = num
			bitset[absnum] |= 2
		}
		if bitset[absnum] == 3 && absnum > k {
			k = absnum
		}
	}
	return k
}

// // time: O(n log n)
// // space: O(1)
// func findMaxK(nums []int) int {
// 	sort.Ints(nums)
// 	for left, right := 0, len(nums)-1; left < right; {
// 		if nums[left] > 0 || nums[right] < 0 {
// 			break
// 		}
// 		if -nums[left] == nums[right] {
// 			return nums[right]
// 		}
// 		if -nums[left] > nums[right] {
// 			left++
// 		} else {
// 			right--
// 		}
// 	}
// 	return -1
// }

// // time: O(n)
// // space: O(n)
// func findMaxK(nums []int) int {
// 	var seen = make(map[int]bool)
// 	var k = -1
// 	var pos int
// 	for _, num := range nums {
// 		pos = int(math.Abs(float64(num)))
// 		if seen[-num] && pos > k {
// 			k = pos
// 		}
// 		seen[num] = true
// 	}
// 	return k
// }
