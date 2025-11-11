package dp

// https://leetcode.com/problems/pascals-triangle-ii/

func getRow(rowIndex int) []int {
	row := make([]int, 0, rowIndex+1)
	row = append(row, 1)

	var prev int
	for n := 1; n <= rowIndex; n++ {
		prev = 0
		for i := range n {
			prev, row[i] = row[i], prev+row[i]
		}

		row = append(row, 1)
	}

	return row
}
