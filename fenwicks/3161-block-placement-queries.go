package fenwicks

import (
	"sort"
)

// https://leetcode.com/problems/block-placement-queries/

func getResults(queries [][]int) []bool {
	blocks := []int{0, 50001}

	for _, query := range queries {
		if query[0] == 1 {
			blocks = append(blocks, query[1])
		}
	}

	sort.Ints(blocks)
	n := len(blocks)
	ptrs := make(map[int][2]int, n)

	ptrs[blocks[0]] = [2]int{0, blocks[1]}
	for i := 1; i < n-1; i++ {
		ptrs[blocks[i]] = [2]int{blocks[i-1], blocks[i+1]}
	}

	ptrs[blocks[n-1]] = [2]int{blocks[n-2], 0}

	fenw := make([]int, n)
	update := func(idx int, val int) {
		for ; idx < n; idx += idx & -idx {
			fenw[idx] = max(fenw[idx], val)
		}
	}
	query := func(idx int) int { // getMax
		var ans int
		for ; idx > 0; idx -= idx & -idx {
			ans = max(ans, fenw[idx])
		}

		return ans
	}

	var length int
	for i := range n - 1 {
		length = blocks[i+1] - blocks[i]
		update(i+1, length)
	}

	indices := make(map[int]int, n)
	for i, pos := range blocks {
		indices[pos] = i + 1
	}

	var prv, nxt, x, sz, left, right, mid int

	ans := make([]bool, 0, len(queries)-n+2)
	for idx := len(queries) - 1; idx >= 0; idx-- {
		x = queries[idx][1]
		if queries[idx][0] == 1 {
			prv, nxt = ptrs[x][0], ptrs[x][1]
			update(indices[prv], 0)
			update(indices[x], 0)

			ptrs[prv] = [2]int{ptrs[prv][0], nxt}
			ptrs[nxt] = [2]int{prv, ptrs[nxt][1]}
			update(indices[prv], nxt-prv)
		} else {
			sz = queries[idx][2]
			if sz > x || sz > query(n-1) {
				ans = append(ans, false)

				continue
			}

			left, right = 1, n-1
			for left <= right {
				mid = left + (right-left)/2
				if query(mid) >= sz {
					right = mid - 1
				} else {
					left = mid + 1
				}
			}

			ans = append(ans, left <= n-1 && blocks[left-1]+sz <= x)
		}
	}

	left, right = 0, len(ans)-1
	for left < right {
		ans[left], ans[right] = ans[right], ans[left]
		left++
		right--
	}

	return ans
}
