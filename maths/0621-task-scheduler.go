package maths

// https://leetcode.com/problems/task-scheduler/

func leastInterval(tasks []byte, n int) int {
	counts := make(map[byte]int)
	for _, task := range tasks { // O(n)
		counts[task]++
	}

	maximum := 0
	maxCount := 0
	for _, count := range counts { // O(26)
		if count > maximum {
			maximum = count
			maxCount = 1
		} else if count == maximum {
			maxCount++
		}
	}

	// There are (maximum-1) scheduled occurrences,
	// and between each two consecutive occurrences,
	// there are at least n CPU cycles: (n+1) including the task itself.
	// If there are multiple elements with a frequency equal to maxCount, add 1 cycle each
	return max(len(tasks), (maximum-1)*(n+1)+maxCount)
}
