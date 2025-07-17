package dp

// https://leetcode.com/problems/find-the-maximum-length-of-valid-subsequence-i/

func maximumLength(nums []int) int {
    var cnt, end [2]int
	for _, num := range nums {
		cnt[num%2]++
		end[num%2] = end[1-num%2]+1
	}
	return max(cnt[0],cnt[1],end[0],end[1])
}