package arrays

// https://leetcode.com/problems/find-the-difference-of-two-arrays/

func findDifference(nums1 []int, nums2 []int) [][]int {
	var mask [2001]int
	for _, num := range nums1 {
		mask[num+1000] |= 1
	}
	for _, num := range nums2 {
		mask[num+1000] |= 2
	}
	var answer = make([][]int, 2)
	answer[0] = make([]int, 0)
	answer[1] = make([]int, 0)
	for num, mask := range mask {
		switch mask {
		case 1:
			answer[0] = append(answer[0], num-1000)
		case 2:
			answer[1] = append(answer[1], num-1000)
		}
	}
	return answer
}
