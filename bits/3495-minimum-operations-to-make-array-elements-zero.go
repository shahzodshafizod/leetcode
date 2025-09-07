package bits

// https://leetcode.com/problems/minimum-operations-to-make-array-elements-zero/

// Approach #3
// Time: O(16*n)
// Space: O(1)
func minOperations3495(queries [][]int) int64 {
	var (
		count       int64
		left, right int
	)

	for idx := range queries {
		left, right = queries[idx][0], queries[idx][1]

		var (
			ops        int64
			start, end int
		)

		for d := 1; d <= 16; d++ {
			start = 1 << (2 * (d - 1))
			end = start<<2 - 1
			start = max(start, left)
			end = min(end, right)

			if start <= end {
				ops += int64((end - start + 1) * d)
			}
		}

		count += (ops + 1) / 2
	}

	return count
}

// // Approach #2
// // Time: O(16*n)
// // Space: O(1)
// func minOperations3495(queries [][]int) int64 {
// 	get := func(num int) int64 { // O(16)
// 		var count, quats int64 = 0, 1

// 		var end int
// 		for start := 1; start <= num; start <<= 2 {
// 			end = min(4*start-1, num)
// 			count += int64(end-start+1) * quats
// 			quats++
// 		}

// 		return count
// 	}

// 	var (
// 		count       int64
// 		left, right int
// 	)

// 	for idx := range queries {
// 		left, right = queries[idx][0], queries[idx][1]
// 		count += (get(right) - get(left-1) + 1) / 2
// 	}

// 	return count
// }

// // Approach #1
// // Time: O(32*n)
// // Space: O(1)
// func minOperations3495(queries [][]int) int64 {
// 	get := func(num int) int64 { // O(32)
// 		var (
// 			count, bits int64 = 0, 1
// 			end         int
// 		)

// 		for start := 1; start <= num; start <<= 1 {
// 			end = min(start*2-1, num)
// 			count += int64(end-start+1) * ((bits + 1) / 2)
// 			bits++
// 		}

// 		return count
// 	}

// 	var (
// 		count       int64
// 		left, right int
// 	)

// 	for idx := range queries {
// 		left, right = queries[idx][0], queries[idx][1]
// 		count += (get(right) - get(left-1) + 1) / 2
// 	}

// 	return count
// }
