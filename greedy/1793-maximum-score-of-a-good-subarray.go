package greedy

// https://leetcode.com/problems/maximum-score-of-a-good-subarray/

/*
Lets say that nums[k]=7 and consider following casses.
You want to increase one end of your interval by one. Which side you should choose?
Lets increse the interval end points one at a time:
[.............., 3, 7, 6, ..........................] 3 and 6 are candidates. Choose 6
[............... 3, 7, 6, 4, .......................] 3 and 4 are candidates. Choose 4
[.............., 3, 7, 6, 4, 7,.....................]
[.............., 3, 7, 6, 4, 7, 2...................]
[.............4, 3, 7, 6, 4, 7, 2...................]
[..........2, 4, 3, 7, 6, 4, 7, 2...................]
[.......3, 2, 4, 3, 7, 6, 4, 7, 2, 1................]
Note: Edge cases are when leftEnd = 0 and rightEnd = len(nums) - 1
*/

func maximumScore(nums []int, k int) int {
	var n = len(nums)
	var score = nums[k]
	var minimum = nums[k]
	var left, right = k, k
	var lnum, rnum int
	for left > 0 || right+1 < n {
		lnum, rnum = 0, 0
		if left > 0 {
			lnum = nums[left-1]
		}
		if right+1 < n {
			rnum = nums[right+1]
		}
		if lnum > rnum {
			left--
			minimum = min(minimum, lnum)
		} else {
			right++
			minimum = min(minimum, rnum)
		}
		score = max(score, minimum*(right-left+1))
	}
	return score
}
