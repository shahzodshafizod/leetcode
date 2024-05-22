package arrays

import (
	"sort"
)

// https://leetcode.com/problems/find-array-given-subset-sums/

func recoverArray(n int, sums []int) []int {
	sort.Ints(sums)
	var num, exclude int
	var positive bool
	var ans = make([]int, 0, n)
	for length := len(sums); length >= 2; length = len(sums) {
		var count = make(map[int]int)
		for _, sum := range sums {
			count[sum]++
		}
		var including = make([]int, 0, length/2)
		var excluding = make([]int, 0, length/2)
		num = sums[length-1] - sums[length-2] // OR sums[1] - sums[0]
		positive = false
		for _, include := range sums {
			exclude = include - num
			if count[include] > 0 && count[exclude] > 0 {
				count[include]--
				count[exclude]--
				including = append(including, include)
				excluding = append(excluding, exclude)
				if exclude == 0 {
					positive = true
				}
			}
		}
		if positive {
			ans = append(ans, num)
			sums = excluding
		} else {
			ans = append(ans, -num)
			sums = including
		}
	}
	return ans
}
