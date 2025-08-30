package prefixsums

// https://leetcode.com/problems/minimum-number-of-operations-to-move-all-balls-to-each-box/

func minOperations(boxes string) []int {
	n := len(boxes)
	right, rboxes := 0, 0

	for idx := n - 1; idx >= 0; idx-- {
		if boxes[idx] == '1' {
			rboxes++
		}

		right += rboxes
	}

	left, lboxes := 0, 0
	answer := make([]int, n)

	for idx := range n {
		left += lboxes
		right -= rboxes
		answer[idx] = left + right

		if boxes[idx] == '1' {
			lboxes++
			rboxes--
		}
	}

	return answer
}
