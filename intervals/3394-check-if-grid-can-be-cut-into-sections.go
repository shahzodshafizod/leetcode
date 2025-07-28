package intervals

import "sort"

// https://leetcode.com/problems/check-if-grid-can-be-cut-into-sections/

func checkValidCuts(n int, rectangles [][]int) bool {
	_ = n
	check := func(dim int) bool {
		sort.Slice(rectangles, func(i int, j int) bool {
			return rectangles[i][dim] < rectangles[j][dim]
		})

		last, count := 0, 0
		for idx := range rectangles {
			if last <= rectangles[idx][dim] {
				count++
				if count == 3 {
					break
				}
			}

			last = max(last, rectangles[idx][dim+2])
		}

		return count == 3
	}

	return check(0) || check(1)
}
