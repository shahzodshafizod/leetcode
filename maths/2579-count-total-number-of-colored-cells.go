package maths

// https://leetcode.com/problems/count-total-number-of-colored-cells/

func coloredCells(n int) int64 {
	// return int64(1 + n*(n-1)*2)
	var count int64 = 1
	for i := 1; i < n; i++ {
		count += int64(4 * i)
	}

	return count
}
