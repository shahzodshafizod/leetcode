package graphs

// https://leetcode.com/problems/maximum-candies-you-can-get-from-boxes/

// Approach: Breadth-First Search
// Time: O(nn)
// Space: O(nn)
func maxCandies(
	status []int,
	candies []int,
	keys [][]int,
	containedBoxes [][]int,
	initialBoxes []int,
) int {
	n := len(status)
	seen := make([]bool, n)
	queue := make([]int, n)
	qid, qlen := 0, 0
	for _, node := range initialBoxes {
		seen[node] = true
		queue[qlen] = node
		qlen++
	}
	count := 0
	var node int
	for qid < qlen {
		for k := qid; k < qlen; k++ {
			for _, box := range keys[queue[k]] {
				if box != queue[k] {
					status[box] = 1
				}
			}
		}
		for k := qlen; qid < k; qid++ {
			node = queue[qid]
			if status[node] == 0 {
				continue
			}
			count += candies[node]
			for _, next := range containedBoxes[node] {
				if !seen[next] {
					seen[next] = true
					queue[qlen] = next
					qlen++
				}
			}
		}
	}
	return count
}
