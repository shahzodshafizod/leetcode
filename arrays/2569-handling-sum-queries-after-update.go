package arrays

// https://leetcode.com/problems/handling-sum-queries-after-update/

// Approach 2: Optimized - Segment Tree with Lazy Propagation
// Use a segment tree to efficiently handle range flip operations.
// Each node stores the count of 1s in its range.
// Lazy propagation defers flip operations until needed.
//
// Key insights:
// 1. For Type 2 queries, we only need count of 1s in nums1, not the array itself
// 2. Flipping a range [l, r] changes count by: (r - l + 1) - 2 * count_in_range
// 3. Segment tree handles range flips in O(log n) instead of O(n)
//
// Time Complexity: O((n + q) * log n) where q = number of queries
// Space Complexity: O(n) for segment tree
func handleQuery(nums1 []int, nums2 []int, queries [][]int) []int64 {
	n := len(nums1)

	// Segment Tree with Lazy Propagation
	type SegmentTree struct {
		n    int
		tree []int  // Stores count of 1s in each segment
		lazy []bool // Stores pending flip operations
	}

	// Build the segment tree
	var build func(st *SegmentTree, node, start, end int)
	build = func(st *SegmentTree, node, start, end int) {
		if start == end {
			st.tree[node] = nums1[start]
		} else {
			mid := (start + end) / 2
			leftChild := 2*node + 1
			rightChild := 2*node + 2
			build(st, leftChild, start, mid)
			build(st, rightChild, mid+1, end)
			st.tree[node] = st.tree[leftChild] + st.tree[rightChild]
		}
	}

	// Push lazy updates down to children
	var push func(st *SegmentTree, node, start, end int)
	push = func(st *SegmentTree, node, start, end int) {
		if st.lazy[node] {
			// Flip current node: count of 1s becomes (total - count of 1s)
			st.tree[node] = (end - start + 1) - st.tree[node]

			// Propagate to children if not a leaf
			if start != end {
				leftChild := 2*node + 1
				rightChild := 2*node + 2
				st.lazy[leftChild] = !st.lazy[leftChild]
				st.lazy[rightChild] = !st.lazy[rightChild]
			}

			st.lazy[node] = false
		}
	}

	// Flip range [l, r]
	var flip func(st *SegmentTree, node, start, end, l, r int)
	flip = func(st *SegmentTree, node, start, end, l, r int) {
		// Push pending updates
		push(st, node, start, end)

		// No overlap
		if start > r || end < l {
			return
		}

		// Complete overlap
		if start >= l && end <= r {
			st.lazy[node] = true
			push(st, node, start, end)
			return
		}

		// Partial overlap
		mid := (start + end) / 2
		leftChild := 2*node + 1
		rightChild := 2*node + 2
		flip(st, leftChild, start, mid, l, r)
		flip(st, rightChild, mid+1, end, l, r)

		// Update current node
		push(st, leftChild, start, mid)
		push(st, rightChild, mid+1, end)
		st.tree[node] = st.tree[leftChild] + st.tree[rightChild]
	}

	// Get total count of 1s
	getTotalOnes := func(st *SegmentTree) int {
		push(st, 0, 0, st.n-1)
		return st.tree[0]
	}

	// Initialize segment tree
	segTree := &SegmentTree{
		n:    n,
		tree: make([]int, 4*n),
		lazy: make([]bool, 4*n),
	}
	build(segTree, 0, 0, n-1)

	// Calculate initial sum of nums2
	var sum2 int64 = 0
	for _, num := range nums2 {
		sum2 += int64(num)
	}

	result := []int64{}

	for _, query := range queries {
		switch query[0] {
		case 1:
			// Flip range [l, r]
			l, r := query[1], query[2]
			flip(segTree, 0, 0, n-1, l, r)
		case 2:
			// Add count of 1s * p to sum2
			count1 := getTotalOnes(segTree)
			sum2 += int64(count1) * int64(query[1])
		default:
			// Return current sum
			result = append(result, sum2)
		}
	}

	return result
}

// // Approach 1: Brute Force
// // For each query:
// // - Type 1: Flip each bit in range [l, r] individually - O(n) per query
// // - Type 2: Update sum by adding count of 1s * p - O(1)
// // - Type 3: Return current sum - O(1)
// //
// // Time Complexity: O(q * n) where q = number of queries, n = length of nums1
// // Space Complexity: O(1) excluding output array
// // This approach will TLE for large inputs (n, q up to 10^5)
// func handleQuery(nums1 []int, nums2 []int, queries [][]int) []int64 {
// 	result := []int64{}

// 	// Calculate initial sum of nums2
// 	var sum2 int64 = 0
// 	for _, num := range nums2 {
// 		sum2 += int64(num)
// 	}

// 	// Calculate count of 1s in nums1
// 	count1 := 0
// 	for _, num := range nums1 {
// 		count1 += num
// 	}

// 	for _, query := range queries {
// 		switch query[0] {
// 		case 1:
// 			// Flip range [l, r] in nums1
// 			l, r := query[1], query[2]
// 			for i := l; i <= r; i++ {
// 				if nums1[i] == 0 {
// 					nums1[i] = 1
// 					count1++
// 				} else {
// 					nums1[i] = 0
// 					count1--
// 				}
// 			}
// 		case 2:
// 			// Add count1 * query[1] to sum2
// 			sum2 += int64(count1) * int64(query[1])
// 		default:
// 			// Return current sum2
// 			result = append(result, sum2)
// 		}
// 	}

// 	return result
// }
