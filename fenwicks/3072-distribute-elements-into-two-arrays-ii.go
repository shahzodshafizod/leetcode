package fenwicks

import (
	"sort"
)

// https://leetcode.com/problems/distribute-elements-into-two-arrays-ii/

// Approach #2: Binary Indexed Tree
// Time: O(n log n)
// Space: O(n)
func resultArray(nums []int) []int {
	type bit struct {
		tree []int
		size int
	}

	var (
		newBIT = func(size int) *bit {
			return &bit{tree: make([]int, size+1), size: size}
		}
		update = func(b *bit, idx int) {
			idx++
			for idx <= b.size {
				b.tree[idx]++
				idx += idx & -idx
			}
		}
		smallerCount = func(b *bit, idx int) int {
			idx++

			var count int
			for idx > 0 {
				count += b.tree[idx]
				idx -= idx & -idx
			}

			return count
		}
		query = func(b *bit, idx int) int {
			return smallerCount(b, b.size-1) - smallerCount(b, idx)
		}
	)

	greaterCount := func(arr *bit, val int) int {
		return query(arr, val)
	}

	set := make(map[int]struct{})
	for _, num := range nums {
		set[num] = struct{}{}
	}

	snums := make([]int, 0, len(set))
	for num := range set {
		snums = append(snums, num)
	}

	sort.Ints(snums)

	indices := make(map[int]int, len(snums))
	for idx, num := range snums {
		indices[num] = idx
	}

	size := len(indices)
	bit1, bit2 := newBIT(size), newBIT(size)
	update(bit1, indices[nums[0]])
	update(bit2, indices[nums[1]])
	arr1, arr2 := []int{nums[0]}, []int{nums[1]}

	var idx, gc1, gc2 int
	for i, n := 2, len(nums); i < n; i++ {
		idx = indices[nums[i]]
		gc1 = greaterCount(bit1, idx)

		gc2 = greaterCount(bit2, idx)
		if gc1 > gc2 || gc1 == gc2 && len(arr1) <= len(arr2) {
			arr1 = append(arr1, nums[i])
			update(bit1, indices[nums[i]])
		} else {
			arr2 = append(arr2, nums[i])
			update(bit2, indices[nums[i]])
		}
	}

	return append(arr1, arr2...)
}

// // Approach #1: Brute-Force
// // Time: O(nn)
// // Space: O(n)
// func resultArray(nums []int) []int {
// 	greaterCount := func(arr []int, val int) int {
// 		var count int

// 		for _, num := range arr {
// 			if num > val {
// 				count++
// 			}
// 		}

// 		return count
// 	}

// 	arr1, arr2 := []int{nums[0]}, []int{nums[1]}
// 	for idx, n := 2, len(nums); idx < n; idx++ {
// 		g1 := greaterCount(arr1, nums[idx])

// 		g2 := greaterCount(arr2, nums[idx])
// 		if g1 > g2 {
// 			arr1 = append(arr1, nums[idx])
// 		} else if g2 > g1 {
// 			arr2 = append(arr2, nums[idx])
// 		} else if len(arr2) < len(arr1) {
// 			arr2 = append(arr2, nums[idx])
// 		} else {
// 			arr1 = append(arr1, nums[idx])
// 		}
// 	}

// 	return append(arr1, arr2...)
// }
