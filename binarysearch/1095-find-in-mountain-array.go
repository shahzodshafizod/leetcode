package binarysearch

// https://leetcode.com/problems/find-in-mountain-array/

/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (m *MountainArray) get(index int) int {}
 * func (m *MountainArray) length() int {}
 */

type MountainArray struct {
	arr []int
}

func (m *MountainArray) get(index int) int {
	if index < 0 || index >= m.length() {
		return -1
	}
	return m.arr[index]
}
func (m *MountainArray) length() int { return len(m.arr) }

func findInMountainArray(target int, ma *MountainArray) int {
	peak := 1
	left, right := 1, ma.length()-2
	for left <= right {
		peak = left + (right-left)/2
		if ma.get(peak-1) < ma.get(peak) {
			left = peak + 1
		} else {
			right = peak - 1
		}
	}
	left, right = 0, peak
	var mid, val int
	for left <= right {
		mid = left + (right-left)/2
		val = ma.get(mid)
		if val < target {
			left = mid + 1
		} else if val > target {
			right = mid - 1
		} else {
			return mid
		}
	}
	left, right = peak, ma.length()-1
	for left <= right {
		mid = left + (right-left)/2
		val = ma.get(mid)
		if val < target {
			right = mid - 1
		} else if val > target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
