package arrays

/*
Convert the arr
[9][4][2][10][7][8][8][1][9]

into a sign arr:
[1][1][-1][1][-1][0][1][-1][9]

We don't need the last element, because
there is one wall among two columns.

And then find the longest alternating area
*/

// https://leetcode.com/problems/longest-turbulent-subarray/

// time: O(n)
// space: O(1)
func maxTurbulenceSize(arr []int) int {
	length, maxlen := 0, 0
	prev, curr := 0, 0

	for idx, len := 1, len(arr); idx < len; idx++ {
		if arr[idx-1] > arr[idx] {
			curr = 1
		} else if arr[idx-1] < arr[idx] {
			curr = -1
		} else {
			curr = 0
		}

		if curr == 0 {
			length = 0
		} else if prev+curr == 0 {
			length++
		} else {
			length = 1
		}

		maxlen = max(maxlen, length)
		prev = curr
	}

	return maxlen + 1
}

// // time: O(2n) = O(n)
// // space: O(1)
// func maxTurbulenceSize(arr []int) int {
// 	var n = len(arr)
// 	for idx := 1; idx < n; idx++ {
// 		if arr[idx-1] > arr[idx] {
// 			arr[idx-1] = 1
// 		} else if arr[idx-1] < arr[idx] {
// 			arr[idx-1] = -1
// 		} else {
// 			arr[idx-1] = 0
// 		}
// 	}
// 	var length, maxlen = 0, 0
// 	for idx := 1; idx < n-1; idx++ {
// 		if arr[idx] == 0 {
// 			length = 0
// 		} else if arr[idx-1]+arr[idx] == 0 {
// 			length++
// 		} else {
// 			length = 1
// 		}
// 		maxlen = max(maxlen, length)
// 	}
// 	return maxlen + 1
// }
