package fenwicks

// https://leetcode.com/problems/peaks-in-array/

func countOfPeaks(nums []int, queries [][]int) []int {
	n := len(nums)
	bitree := make([]int, n+1)
	isPeak := func(i int) bool {
		return i > 0 && i < n-1 && nums[i-1] < nums[i] && nums[i] > nums[i+1]
	}
	update := func(index int, value int) {
		index++
		for ; index <= n; index += index & -index {
			bitree[index] += value
		}
	}
	query := func(index int) int {
		var total int

		index++
		for ; index > 0; index -= index & -index {
			total += bitree[index]
		}

		return total
	}

	for i := range n {
		if isPeak(i) {
			update(i, 1)
		}
	}

	var (
		counts                    []int
		index, value, left, right int
		curr                      bool
	)

	for i := range queries {
		if queries[i][0] == 2 {
			index, value = queries[i][1], queries[i][2]
			before := []bool{isPeak(index - 1), isPeak(index), isPeak(index + 1)}

			nums[index] = value
			for i := range before {
				curr = isPeak(index + i - 1)
				if before[i] && !curr {
					update(index+i-1, -1)
				} else if !before[i] && curr {
					update(index+i-1, 1)
				}
			}
		} else {
			left, right = queries[i][1], queries[i][2]
			if left < right {
				// [left+1, right-1]
				counts = append(counts, query(right-1)-query(left))
			} else {
				counts = append(counts, 0)
			}
		}
	}

	return counts
}
