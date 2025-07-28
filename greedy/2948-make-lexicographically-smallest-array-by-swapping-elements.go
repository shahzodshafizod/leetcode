package greedy

import (
	"container/list"
	"sort"
)

// https://leetcode.com/problems/make-lexicographically-smallest-array-by-swapping-elements/

func lexicographicallySmallestArray(nums []int, limit int) []int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}

		return x
	}
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)

	groups := make([]*list.List, 0, 10)
	size := 0
	groupIndices := make(map[int]int, len(nums))

	for _, num := range sorted {
		if size == 0 || abs(num-groups[size-1].Back().Value.(int)) > limit {
			groups = append(groups, list.New())
			size++
		}

		groups[size-1].PushBack(num)
		groupIndices[num] = size - 1
	}

	var group *list.List
	for idx, num := range nums {
		group = groups[groupIndices[num]]
		sorted[idx] = group.Remove(group.Front()).(int)
	}

	return sorted
}
