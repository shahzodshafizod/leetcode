package monotonic

// https://leetcode.com/problems/maximal-rectangle/

// solution is based on https://leetcode.com/problems/largest-rectangle-in-histogram/
// time: O(rows*cols)
// space: O(cols)
func maximalRectangle(matrix [][]byte) int {
	rows, cols := len(matrix), len(matrix[0])
	maximal := 0
	heights := make([]int, cols)
	stack := make([][2]int, cols) // increasing stack {index, height}

	var top, index, height, width int
	for row := range rows { // O(rows)
		top = -1 // reset the stack

		for col := range cols { // O(rows*cols)
			if matrix[row][col] == '0' {
				heights[col] = 0
			} else {
				heights[col]++
			}

			index = col
			for top >= 0 && stack[top][1] >= heights[col] { // O(rows*2*cols)
				index = stack[top][0]
				width = col - index
				height = stack[top][1]
				maximal = max(maximal, height*width)
				top--
			}

			top++
			stack[top][0] = index
			stack[top][1] = heights[col]
		}

		for top >= 0 { // O(rows*cols)
			width = cols - stack[top][0]
			height = stack[top][1]
			maximal = max(maximal, height*width)
			top--
		}
	}

	return maximal
}
