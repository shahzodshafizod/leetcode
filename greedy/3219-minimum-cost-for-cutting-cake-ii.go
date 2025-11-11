package greedy

import "sort"

// https://leetcode.com/problems/minimum-cost-for-cutting-cake-ii/

func minimumCost(m int, n int, horizontalCut []int, verticalCut []int) int64 {
	sort.Ints(horizontalCut)
	sort.Ints(verticalCut)

	hid, vid := m-2, n-2
	hpieces, vpieces := 1, 1

	var total int64
	for hid >= 0 && vid >= 0 {
		if horizontalCut[hid] > verticalCut[vid] {
			total += int64(horizontalCut[hid] * vpieces)
			hid--
			hpieces++
		} else {
			total += int64(verticalCut[vid] * hpieces)
			vid--
			vpieces++
		}
	}

	for ; hid >= 0; hid-- {
		total += int64(horizontalCut[hid] * vpieces)
	}
	for ; vid >= 0; vid-- {
		total += int64(verticalCut[vid] * hpieces)
	}

	return total
}
