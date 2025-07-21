package dp

// https://leetcode.com/problems/pascals-triangle/

// Approach: Bottom-Up Dynamic Programming
// Time: O(nn)
// Space: O(nn)
func generate(numRows int) [][]int {
	rows := make([][]int, numRows)
	for len := 1; len <= numRows; len++ {
		row := make([]int, len)
		row[0], row[len-1] = 1, 1
		for idx := 1; idx < len-1; idx++ {
			row[idx] = rows[len-2][idx-1] + rows[len-2][idx]
		}
		rows[len-1] = row
	}
	return rows
}

// // Approach: Recursive
// // Time: O(nn)
// // Space: O(nn)
// func generate(numRows int) [][]int {
// 	if numRows == 1 {
// 		return [][]int{{1}}
// 	}
// 	var rows = generate(numRows - 1)
// 	var curr = make([]int, numRows)
// 	curr[0], curr[numRows-1] = 1, 1
// 	for idx := 1; idx < numRows-1; idx++ {
// 		curr[idx] = rows[numRows-2][idx-1] + rows[numRows-2][idx]
// 	}
// 	rows = append(rows, curr)
// 	return rows
// }
