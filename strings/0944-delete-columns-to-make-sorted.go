package strings

// https://leetcode.com/problems/delete-columns-to-make-sorted/

// Approach: Optimized - Single pass column check
// Iterate through each column and check if it's sorted
// Count columns where any adjacent pair is out of order
// N = number of strings, M = length of each string
// Time: O(N * M) - check each character once
// Space: O(1) - only counter variable
func minDeletionSize(strs []string) int {
	if len(strs) == 0 {
		return 0
	}

	rows := len(strs)
	cols := len(strs[0])
	deleteCount := 0

	for col := range cols {
		// Check if this column is sorted
		for row := 1; row < rows; row++ {
			if strs[row][col] < strs[row-1][col] {
				// Found unsorted pair, delete this column
				deleteCount++

				break
			}
		}
	}

	return deleteCount
}
