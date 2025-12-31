package design

import (
	"math/rand"
)

// https://leetcode.com/problems/insert-delete-getrandom-o1-duplicates-allowed/

// Approach #2: ArrayList + HashMap with Set of Indices
// Time: insert O(1), remove O(1), getRandom O(1)
// Space: O(n)
type RandomizedCollection struct {
	nums    []int                // ArrayList to store values
	indices map[int]map[int]bool // value -> set of indices (using map as set)
}

func NewRandomizedCollection() RandomizedCollection {
	return RandomizedCollection{
		nums:    []int{},
		indices: make(map[int]map[int]bool),
	}
}

func (r *RandomizedCollection) Insert(val int) bool {
	// Check if val exists
	_, exists := r.indices[val]

	// Add index to set
	if r.indices[val] == nil {
		r.indices[val] = make(map[int]bool)
	}

	r.indices[val][len(r.nums)] = true

	// Append to array
	r.nums = append(r.nums, val)

	return !exists
}

func (r *RandomizedCollection) Remove(val int) bool {
	// Check if val exists
	if _, exists := r.indices[val]; !exists || len(r.indices[val]) == 0 {
		return false
	}

	// Get any index of val (iterate to get first one from map)
	var removeIdx int
	for idx := range r.indices[val] {
		removeIdx = idx

		break
	}

	delete(r.indices[val], removeIdx)

	lastIdx := len(r.nums) - 1
	lastVal := r.nums[lastIdx]

	// Move last element to removed position
	r.nums[removeIdx] = lastVal

	// Update indices for last element (only if removeIdx != lastIdx)
	if r.indices[lastVal] == nil {
		r.indices[lastVal] = make(map[int]bool)
	}

	r.indices[lastVal][removeIdx] = true
	delete(r.indices[lastVal], lastIdx)

	// Remove last element
	r.nums = r.nums[:len(r.nums)-1]

	// Clean up empty maps
	if len(r.indices[val]) == 0 {
		delete(r.indices, val)
	}

	return true
}

func (r *RandomizedCollection) GetRandom() int {
	return r.nums[rand.Intn(len(r.nums))]
}

// // Approach #1: Brute Force (Slice only)
// // Time: insert O(1), remove O(n), getRandom O(1)
// // Space: O(n)
// type RandomizedCollection struct {
// 	nums []int
// }

// func NewRandomizedCollection() RandomizedCollection {
// 	return RandomizedCollection{
// 		nums: []int{},
// 	}
// }

// func (r *RandomizedCollection) Insert(val int) bool {
// 	// Check if val exists - O(n)
// 	exists := slices.Contains(r.nums, val)
// 	r.nums = append(r.nums, val)
// 	return !exists
// }

// func (r *RandomizedCollection) Remove(val int) bool {
// 	// Find and remove first occurrence - O(n)
// 	for i, num := range r.nums {
// 		if num == val {
// 			r.nums = append(r.nums[:i], r.nums[i+1:]...)
// 			return true
// 		}
// 	}
// 	return false
// }

// func (r *RandomizedCollection) GetRandom() int {
// 	return r.nums[rand.Intn(len(r.nums))]
// }

/*
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := NewRandomizedCollection();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
