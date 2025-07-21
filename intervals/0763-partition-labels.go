package intervals

// https://leetcode.com/problems/partition-labels/

func partitionLabels(s string) []int {
	last := make(map[rune]int)
	for idx, c := range s {
		last[c] = idx
	}
	start, end := -1, -1
	sizes := make([]int, 0)
	for idx, c := range s {
		end = max(end, last[c])
		if idx == end {
			sizes = append(sizes, end-start)
			start = idx
		}
	}
	return sizes
}
