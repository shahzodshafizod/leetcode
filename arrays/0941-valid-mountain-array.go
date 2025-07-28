package arrays

// https://leetcode.com/problems/valid-mountain-array/

func validMountainArray(arr []int) bool {
	n := len(arr)
	if n < 3 { // Recall that arr is a mountain array if and only if: arr.length >= 3
		return false
	}
	// walk up
	idx := 0
	for idx+1 < n && arr[idx] < arr[idx+1] {
		idx++
	}
	// peak can't be first or last
	if idx == 0 || idx == n-1 {
		return false
	}
	// walk down
	for idx+1 < n && arr[idx] > arr[idx+1] {
		idx++
	}

	return idx == n-1
}
