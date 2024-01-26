package matrices

// https://leetcode.com/problems/search-a-2d-matrix/

func searchMatrix(matrix [][]int, target int) bool {
	var m, n = len(matrix), len(matrix[0])
	var left, right = 0, m*n - 1
	for left <= right {
		var mid = (left + right) / 2
		var midElem = matrix[mid/n][mid%n]
		if target == midElem {
			return true
		} else if target < midElem {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

// func searchMatrix(matrix [][]int, target int) bool {
// 	var m, n = len(matrix), len(matrix[0])
// 	if target < matrix[0][0] || target > matrix[m-1][n-1] {
// 		return false
// 	}
// 	var row = 0
// 	for ; row < m; row++ {
// 		if target <= matrix[row][n-1] {
// 			break
// 		}
// 	}
// 	// binary search of the array matrix[row]
// 	var left, right = 0, n - 1
// 	for left <= right {
// 		var mid = (left + right) / 2
// 		if target > matrix[row][mid] {
// 			left = mid + 1
// 		} else if target < matrix[row][mid] {
// 			right = mid - 1
// 		} else {
// 			return true
// 		}
// 	}
// 	return false
// }
