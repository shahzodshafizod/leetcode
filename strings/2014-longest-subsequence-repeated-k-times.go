package strings

import (
	"container/list"
	"slices"
)

// https://leetcode.com/problems/longest-subsequence-repeated-k-times/

// Approach: Backtracking + Breadth-First Search
// Time: O(7! * n)
// Space: O(7!)
func longestSubsequenceRepeatedK(s string, k int) string {
	calcCount := func(subseq []rune) int {
		i, n := 0, len(subseq)
		count := 0
		for _, c := range s {
			if c == subseq[i] {
				i++
				if i == n {
					i = 0
					count++
				}
			}
		}
		return count
	}
	count := make(map[rune]int)
	for _, c := range s {
		count[c]++
	}
	var candidates []rune
	for c, cnt := range count {
		if cnt >= k {
			candidates = append(candidates, c)
		}
	}
	slices.Sort(candidates)
	var queue list.List
	subseq := ""
	queue.PushBack(subseq)
	for queue.Len() > 0 {
		subseq = queue.Remove(queue.Front()).(string)
		nexts := []rune(subseq)
		nexts = append(nexts, ' ')
		n := len(nexts)
		for _, c := range candidates {
			nexts[n-1] = c
			if calcCount(nexts) >= k {
				queue.PushBack(string(nexts))
			}
		}
	}
	return subseq
}
