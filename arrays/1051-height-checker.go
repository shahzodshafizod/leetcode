package arrays

// https://leetcode.com/problems/height-checker/

// time: O(101 x n) = O(n)
// space: O(101) = O(1)
func heightChecker(heights []int) int {
	var expected [101]int
	for _, height := range heights { // O(n)
		expected[height]++
	}

	misplaced := 0
	idx := 0

	for height, count := range expected { // O(101)
		for count > 0 { // O(n)
			count--

			if heights[idx] != height {
				misplaced++
			}

			idx++
		}
	}

	return misplaced
}

// // time: O(n log n)
// // space: O(n)
// func heightChecker(heights []int) int {
// 	var expected = make([]int, len(heights))
// 	copy(expected, heights)
// 	sort.Ints(expected) // O(n log n)
// 	var misplaced = 0
// 	for idx := range expected { // O(n)
// 		if expected[idx] != heights[idx] {
// 			misplaced++
// 		}
// 	}
// 	return misplaced
// }
