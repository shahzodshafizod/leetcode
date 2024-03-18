package intervals

// https://leetcode.com/problems/insert-interval/

func insert(intervals [][]int, newInterval []int) [][]int {
	var result = make([][]int, 0)
	for idx, interval := range intervals {
		if newInterval[1] < interval[0] {
			return append(result, append([][]int{newInterval}, intervals[idx:]...)...)
		} else if interval[1] < newInterval[0] {
			result = append(result, interval)
		} else { // merging
			newInterval[0] = min(newInterval[0], interval[0])
			newInterval[1] = max(newInterval[1], interval[1])
		}
	}
	result = append(result, newInterval)
	return result
}
